---
layout: page
title: Writing a Custom Widget

redirect_from: /develop/custom-widget.html
---

The standard widgets included with Fyne are designed to support standard user interactions and requirements.
As a GUI often has to provide custom functionality it may be necessary to write a custom widget.
This article outlines how.

A widget is split into two areas - each implementing a standard interface - the `fyne.Widget` and the `fyne.WidgetRenderer`. The widget defines behaviour and state, with the renderer being used to define how it should be drawn to screen.

### fyne.Widget

A widget in Fyne is simply a stateful canvas object that has its rendering
definition separated from the main logic. As you can see from the `fyne.Widget`
interface there is not much that must be implemented.

```go
type Widget interface {
	CanvasObject

	CreateRenderer() WidgetRenderer
}

```

As a widget needs to be used like any other canvas item we inherit from the same interface.
To save writing all the methods required we can make use of the `widget.BaseWidget` type which handles the basics.

Each widget definition will contain much more than the interface requires.
It is standard in a Fyne widget to export the fields which define behaviour
(just like the primitives defined in the `canvas` package).

For example, look at the `widget.Button` type:

```go
type Button struct {
	BaseWidget
	Text  string
	Style ButtonStyle
	Icon  fyne.Resource

	OnTapped func()
}
```

You can see how each of these items store state about the widget behaviour but nothing about how it is rendered.

### fyne.WidgetRenderer

The widget renderer is responsible for managing a list of `fyne.CanvasObject` primitives that come together to create the design of our widget. It is much like
a `fyne.Container` with a custom layout and some additional theme handling.

Every widget must provide a renderer, but it is entirely possible to re-use a renderer from another widget - especially if your widget is a lightweight wrapper around another standard control.

```go
type WidgetRenderer interface {
	Layout(Size)
	MinSize() Size

	Refresh()
	Objects() []CanvasObject
	Destroy()
}
```

As you can see the `Layout(Size)` and `MinSize()` methods are similar to the
`fyne.Layout` interface, without the `[]fyne.CanvasObject` parameter - this is because a widget does need to be laid out but it controls which objects will be included.

The `Refresh()` method is triggered when the widget this renderer draws has changed or if the theme is altered.
In either situation we may need to make adjustments to how it looks.
Lastly the `Destroy()` method is called when this renderer is no longer needed so it should clear any resources that would otherwise leak.

Compare again with the button widget - it's `fyne.WidgetRenderer` implementation is based on the following type:

```go
type buttonRenderer struct {
	icon   *canvas.Image
	label  *canvas.Text
	shadow *fyne.CanvasObject

	objects []fyne.CanvasObject
	button  *Button
}
```

As you can see it has fields to cache the actual image, text and shadow canvas
objects for drawing.
It keeps track of the slice of objects required by `fyne.WidgetRenderer` as a convenience.

Lastly it keeps a reference to the `widget.Button` for all state information.
In the `Refresh()` method it will update the graphical state based on any changes
in the underlying `widget.Button` type.

### Bring it together

A basic widget will extend the `widget.BaseWidget` type and declare any state
that the widget holds.
The `CreateRenderer()` method must exist and return a new `fyne.WidgetRenderer` instance. The widget and driver code in Fyne will ensure that this
is cached accordingly - this method may be called many times (for example if a
widget is hidden and then shown). If `CreateRenderer()` is called again you
should return a new renderer instance as the old one may have been destroyed.

Take care not to keep any important state in your renderer - animation tickers
are well suited to that location but user state would not be. A widget that is
hidden may have it's renderer destroyed and if it is shown again the new renderer
must be able to reflect the same widget state.

### Example Widget

The following covers in detail the creation of a simple Widget and its Renderer.

#### Create a Widget

As described above, we need to inherit from the `widget.BaseWidget` type.

Go provides an simple mechanism for inheriting behaviour from other types. Adding an Anonymous type as a field in the widget structure means that the widget inherits the behaviour of that Anonymous type.

This also means that your widget now implements all of the public methods and properties define by the  `widget.BaseWidget` type.

We add an Anonymous `widget.BaseWidget` as follows:

```go
type MyWidget struct {
	widget.BaseWidget           // Inherit from BaseWidget
	text              string	// The text to display in the widget
}
```

To construct the widget create a New<widgetname> function. The example we are using is called MyWidget so by convention the following function name would be used to construct the widget:

```go
func NewMyWidget(text string) *MyWidget {
	w := &MyWidget{
		text: text,
	}
	w.ExtendBaseWidget(w)
	return w
}
```

Note : It is not necessary to override the Resize method (shown below) but if you need to change the Resize behaviour of your widget and you override the ```Resize(s fyne.Size)``` method then you **must** include the ```w.BaseWidget.Resize(s)```. in the code.

```go
func (w *MyWidget) Resize(s fyne.Size) { // See note below
	w.BaseWidget.Resize(s)
}
```


The final code to add is the ```CreateRenderer() fyne.WidgetRenderer``` method as follows:

```go
func (w *MyWidget) CreateRenderer() fyne.WidgetRenderer {
	return newMyWidgetRenderer(w.minSize, w.text)
}
```

The widget is now almost complete except for the ```newMyWidgetRenderer()``` function which creates your new widget renderer. 

**Warning**: Do not keep a reference to the widget renderer in the widget (or any other part of your application). The fyne application caches the renderer and may destroy it at any time. The fyne application will call ```CreateRenderer()``` again if it requires a new renderer instance. It is ok for the widget renderer to refer to the widget but not the other way round.

#### Create a Widget Renderer

The Widget Renderer is how we tell the fyne application what to draw and how to lay it out.

The Widget Renderer has it's own state and methods (see ```fyne.WidgetRenderer``` interface) but it should only be concerned with drawing (rendering) the widget.

The Widget Renderer is responsible for:

1) Defining the objects that are to be drawn. These are returned by the ```Objects() []fyne.CanvasObject``` method. Each object returned will need to be created by the Widget Renderer and initialised before being returned.
3) The position and size of all canvas items within the space passed in to the Layout method.
3) Ensuring that the display is refreshed with changes to the state of the widget. 

Our example renderer could be stored in the following struct:

```
type myWidgetRenderer struct {
	widget     *MyWidget
	background *canvas.Rectangle
	text       *canvas.Text
}
```

As you can see it is all about canvas objects that the fyne application will render.

For this example we need a canvas.Rectangle as a background and a canvas.Text to display the widget text.

We hold a reference to the widget so we can refer to the data it holds. If the data in the widget changes we want the renderer to be able to refresh the display. See the Refresh() method below.

When we construct the renderer we pass in a reference to the widget and store it.

```go
func newMyWidgetRenderer(myWidget *MyWidget) *myWidgetRenderer {
	return &myWidgetRenderer{
		widget:     myWidget,
		background: canvas.NewRectangle(theme.BackgroundColor()),
		text:       canvas.NewText(myWidget.text, theme.ForegroundColor()),
	}
}
```

Notice that the size and position of the canvas objects are not defined here, only initial colours and text properties.

The Objects() method should return the canvas objects in a list as follows:

```go
func (r *myWidgetRenderer) Objects() []fyne.CanvasObject {
    // The order is critical, rect is drawn first then text
	return []fyne.CanvasObject{r.background, r.text}
}
```

Each object returned will need to be sized and positioned in the Layout() method as follows:

```go
func (r *myWidgetRenderer) Layout(s fyne.Size) {
	ts := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	r.text.Move(fyne.Position{X: (s.Width - ts.Width) / 2, Y: (s.Height - ts.Height) / 2})
	r.background.Resize(s)
}
```

1. The size of the canvas.Text (r.text) object is measured so we can center it. 
2. The position of the canvas.Text (r.text) object is centered
3. The size of the canvas.Rectangle (r.background) is set to the same size as the widget.

We also need to define the minimum size of our widget. 

We do this in the renderer because only the renderer knows the size of the canvas objects it contains.

```go
func (r *myWidgetRenderer) MinSize() fyne.Size {
	return fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
}
```

For this example the size is detirmined by the size of the canvas.Text object.

The Refresh() method will be called (by the application logic) if the application changes the state of the widget. 

For example if we change the text property of the widget we would then call Refresh() as follows:

```go
myWidget.text = "Update Widget"
myWidget.Refresh()
```

 For this to work we need to define the Refresh() method of the renderer as follows:

```go
func (r *myWidgetRenderer) Refresh() {
	r.text.Text = r.widget.text
	r.text.Refresh()
}
```

This updates the canvas objects and calls their respective Refresh() methods to force them to re-display the data.

We finish with the remaining methods of the ```fyne.WidgetRenderer``` interface

```go
func (r *myWidgetRenderer) Destroy() {} // Called when the renderer is destroyed
```

#### Finally

If you are unsure that the interface requirements of the renderer have been met you can include the following code at the top of the widget. This is useful if you accidentally change a method or another developer comes along later and makes changes. 

```
var _ fyne.WidgetRenderer = (*myWidgetRenderer)(nil)
```

If ```myWidgetRenderer``` does NOT implement the ```fyne.WidgetRenderer``` interface this line will be in error!

#### Use your Widget

```go
func main() {
	app := app.New()
	w := app.NewWindow("My Widget")
	mw := NewMyWidget(100, 100, "Widget")
	w.SetContent(mw)
	w.ShowAndRun()
}
```

If you stretch the widget the text will remain centered.
