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
To save writing all the functions required we can make use of the `widget.BaseWidget` type which handles the basics.

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

As you can see the `Layout(Size)` and `MinSize()` functions are similar to the
`fyne.Layout` interface, without the `[]fyne.CanvasObject` parameter - this is because a widget does need to be laid out but it controls which objects will be included.

The `Refresh()` function is triggered when the widget this renderer draws has changed or if the theme is altered.
In either situation we may need to make adjustments to how it looks.
Lastly the `Destroy()` function is called when this renderer is no longer needed so it should clear any resources that would otherwise leak.

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
In the `Refresh()` function it will update the graphical state based on any changes
in the underlying `widget.Button` type.

### Bring it together

A basic widget will extend the `widget.BaseWidget` type and declare any state
that the widget holds.
The `CreateRenderer()` function must exist and return a new `fyne.WidgetRenderer` instance. The widget and driver code in Fyne will ensure that this
is cached accordingly - this method may be called many times (for example if a
widget is hidden and then shown). If `CreateRenderer()` is called again you
should return a new renderer instance as the old one may have been destroyed.

Take care not to keep any important state in your renderer - animation tickers
are well suited to that location but user state would not be. A widget that is
hidden may have it's renderer destroyed and if it is shown again the new renderer
must be able to reflect the same widget state.

### Example Widget

The following covers in detail the creation of a simple Widget and it's Renderer.

#### Create a Widget

So the widget requires the`widget.BaseWidget`. For example:

```go
type MyWidget struct {
	widget.BaseWidget           // Inherit from BaseWidget
	minSize           fyne.Size // The minimum size of the widget
	text              string	// The text to display in the widget
}
```

To construct the widget create a New<widgetname> function. The example we are using is called MyWidget so by convention the following constructor function name would be used:

```go
func NewMyWidget(W, H float32, text string) *MyWidget {
	w := &MyWidget{ 							 // Initialise the widget storage
		text:    text,                           // Save the text
		minSize: fyne.Size{Width: W, Height: H}, // Set the size
	}
	w.ExtendBaseWidget(w) // Required to initialise the base widget
	return w
}
```

We need to override the MinSize() function so we can return the minimum size of our widget. 

```go
func (w *MyWidget) MinSize() fyne.Size {
	return w.minSize
}
func (w *MyWidget) Resize(s fyne.Size) { // See note below
	w.BaseWidget.Resize(s)
}
```

Note : It is not necessary to override the Resize function (shown above) but if you need to change the Resize behaviour of your widget and you add a ```Resize(s fyne.Size)``` function then you **must** include the ```w.BaseWidget.Resize(s)```. in the code.

The final function to add is the ```CreateRenderer() fyne.WidgetRenderer``` function:

```go
func (w *MyWidget) CreateRenderer() fyne.WidgetRenderer {
	return newMyWidgetRenderer(w.minSize, w.text)
}
```

**Warning**: Do not keep a reference to the widget renderer in the widget (or any other part of your application). The fyne application caches the renderer and may destroy it at any time.

The fyne application will call ```CreateRenderer()``` again if it requires a new renderer instance.

It is ok for the widget renderer to refer to the widget but not the other way round.

The widget is now almost complete except for the ```newMyWidgetRenderer()``` function which creates your new widget renderer. 

#### Create a Widget Renderer

The widget renderer is how we tell the fyne application what to draw and how to lay it out.

The widget renderer has it's own state and functions (see ```fyne.WidgetRenderer``` interface) but this should only be concerned with drawing (rendering) the widget.

The renderer is responsible for:

1) The objects that are to be drawn. These are returned by the ```Objects() []fyne.CanvasObject``` function. Notice that you can return any object as long as it implements the ```fyne.CanvasObject``` interface.
2) Each object returned by the Objects() function will need to be created by the widget renderer and initialised before returning.
3) The fyne application will then call the renderers Layout() function. This is where the size and position of each object is detirmined.

Our example renderer could be stored in the following struct:

```
type myWidgetRenderer struct {
	minSize fyne.Size // The size of the widget
	rect    *canvas.Rectangle
	text    *canvas.Text
}
```

As you can see it is about canvas objects that the fyne application will render. A lower case name is used as the renderer code is private and only created by the Widget. 

The constructor is the place to pass in properties like minimum size, text values and other data.

```go
func newMyWidgetRenderer(size fyne.Size, text string) *myWidgetRenderer {
	return &myWidgetRenderer{
		minSize: size,  // For the MinSize() function
		rect:    canvas.NewRectangle(theme.BackgroundColor()),
		text:    canvas.NewText(text, theme.ForegroundColor()),
	}
}
```

Notice that the size and position of the canvas objects are not defined here, only colours, text and other properties. It might be appropriate to pass in a reference to the widget instead of lots of different parameters. This is left to your design requirements.

The Objects() function can now return the canvas objects as a list:

```go
func (r *myWidgetRenderer) Objects() []fyne.CanvasObject {
    // The order is critical, rect is drawn first then text
	return []fyne.CanvasObject{r.rect, r.text}
}
```

Each object returned here will need to be sized and positioned in the Layout() function:

```go
func (r *myWidgetRenderer) Layout(s fyne.Size) {
	// Measure the size of the text so we can Center it
	si := fyne.MeasureText(r.text.Text, r.text.TextSize, r.text.TextStyle)
	r.text.Move(fyne.Position{X: (s.Width - si.Width) / 2, Y: (s.Height - si.Height) / 2})
	// Resize the background
	r.rect.Resize(s)
}
```

We finish with the remaining functions of the ```fyne.WidgetRenderer``` interface

```go
func (r *myWidgetRenderer) Destroy() {} // Called when the renderer is destroyed
func (r *myWidgetRenderer) Refresh() {} // This does not seem to be used!
func (r *myWidgetRenderer) MinSize() fyne.Size {
	return r.minSize // Return the minimum size of the widget
}
```

#### Finally

If you are unsure that the interface requirements of the renderer have been met you can include the following code at the top of the widget. This is useful if you accidentally change a function or another developer comes along later and makes changes. 

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

#### Next Steps

Simply add the Tapped function from the fyne.Tappable interface

```go
var _ fyne.Tappable = (*MyWidget)(nil) // This just checks the Tappable interface

func (w *MyWidget) Tapped(*fyne.PointEvent) {
	fmt.Println("MyWidget has been tapped")
}
```

You can now  LEFT click on your widget :-)

See also:

```
SecondaryTappable
DoubleTappable
Disableable
```

Note: Implementing the DoubleTappable interface introduces a slight delay of about 0.5 seconds before the Tapped() function is called. This is to allow for the next tap to be recognised.
