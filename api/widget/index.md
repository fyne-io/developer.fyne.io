---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
--
    import "fyne.io/fyne/widget"

Package widget defines the UI widgets within the Fyne toolkit

## Usage

#### func  DestroyRenderer

```go
func DestroyRenderer(wid fyne.Widget)
```
DestroyRenderer frees a render implementation for a widget. This is typically
for internal use only. Deprecated: Access to widget renderers is being removed,
render details should be private to a WidgetRenderer.

#### func  Refresh

```go
func Refresh(wid fyne.Widget)
```
Refresh instructs the containing canvas to refresh the specified widget.
Deprecated: Call Widget.Refresh() instead.

#### func  Renderer

```go
func Renderer(wid fyne.Widget) fyne.WidgetRenderer
```
Renderer looks up the render implementation for a widget Deprecated: Access to
widget renderers is being removed, render details should be private to a
WidgetRenderer.

#### func  ShowModalPopUp

```go
func ShowModalPopUp(content fyne.CanvasObject, canvas fyne.Canvas)
```
ShowModalPopUp creates a new popUp for the specified content and displays it on
the passed canvas. A modal PopUp blocks interactions with underlying elements,
covered with a semi-transparent overlay.

#### func  ShowPopUp

```go
func ShowPopUp(content fyne.CanvasObject, canvas fyne.Canvas)
```
ShowPopUp creates a new popUp for the specified content and displays it on the
passed canvas.

#### func  ShowPopUpAtPosition

```go
func ShowPopUpAtPosition(content fyne.CanvasObject, canvas fyne.Canvas, pos fyne.Position)
```
ShowPopUpAtPosition creates a new popUp for the specified content at the
specified absolute position. It will then display the popup on the passed
canvas.

#### func  ShowPopUpMenuAtPosition

```go
func ShowPopUpMenuAtPosition(menu *fyne.Menu, c fyne.Canvas, pos fyne.Position)
```
ShowPopUpMenuAtPosition creates a PopUp menu populated with items from the
passed menu structure. It will automatically be positioned at the provided
location and shown as an overlay on the specified canvas.

#### type AccordionContainer

```go
type AccordionContainer struct {
	BaseWidget
	Items     []*AccordionItem
	MultiOpen bool
}
```

AccordionContainer displays a list of AccordionItems. Each item is represented
by a button that reveals a detailed view when tapped.

#### func  NewAccordionContainer

```go
func NewAccordionContainer(items ...*AccordionItem) *AccordionContainer
```
NewAccordionContainer creates a new accordion widget.

#### func (*AccordionContainer) Append

```go
func (a *AccordionContainer) Append(item *AccordionItem)
```
Append adds the given item to this AccordionContainer.

#### func (*AccordionContainer) Close

```go
func (a *AccordionContainer) Close(index int)
```
Close collapses the item at the given index.

#### func (*AccordionContainer) CloseAll

```go
func (a *AccordionContainer) CloseAll()
```
CloseAll collapses all items.

#### func (*AccordionContainer) CreateRenderer

```go
func (a *AccordionContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*AccordionContainer) MinSize

```go
func (a *AccordionContainer) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*AccordionContainer) Open

```go
func (a *AccordionContainer) Open(index int)
```
Open expands the item at the given index.

#### func (*AccordionContainer) OpenAll

```go
func (a *AccordionContainer) OpenAll()
```
OpenAll expands all items.

#### func (*AccordionContainer) Remove

```go
func (a *AccordionContainer) Remove(item *AccordionItem)
```
Remove deletes the given item from this AccordionContainer.

#### func (*AccordionContainer) RemoveIndex

```go
func (a *AccordionContainer) RemoveIndex(index int)
```
RemoveIndex deletes the item at the given index from this AccordionContainer.

#### type AccordionItem

```go
type AccordionItem struct {
	Title  string
	Detail fyne.CanvasObject
	Open   bool
}
```

AccordionItem represents a single item in an AccordionContainer.

#### func  NewAccordionItem

```go
func NewAccordionItem(title string, detail fyne.CanvasObject) *AccordionItem
```
NewAccordionItem creates a new item for an Accordion.

#### type BaseWidget

```go
type BaseWidget struct {
	Hidden bool
}
```

BaseWidget provides a helper that handles basic widget behaviours.

#### func (*BaseWidget) ExtendBaseWidget

```go
func (w *BaseWidget) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget
functionality.

#### func (*BaseWidget) Hide

```go
func (w *BaseWidget) Hide()
```
Hide this widget so it is no lonver visible

#### func (*BaseWidget) MinSize

```go
func (w *BaseWidget) MinSize() fyne.Size
```
MinSize for the widget - it should never be resized below this value.

#### func (*BaseWidget) Move

```go
func (w *BaseWidget) Move(pos fyne.Position)
```
Move the widget to a new position, relative to its parent. Note this should not
be used if the widget is being managed by a Layout within a Container.

#### func (*BaseWidget) Position

```go
func (w *BaseWidget) Position() fyne.Position
```
Position gets the current position of this widget, relative to its parent.

#### func (*BaseWidget) Refresh

```go
func (w *BaseWidget) Refresh()
```
Refresh causes this widget to be redrawn in it's current state

#### func (*BaseWidget) Resize

```go
func (w *BaseWidget) Resize(size fyne.Size)
```
Resize sets a new size for a widget. Note this should not be used if the widget
is being managed by a Layout within a Container.

#### func (*BaseWidget) Show

```go
func (w *BaseWidget) Show()
```
Show this widget so it becomes visible

#### func (*BaseWidget) Size

```go
func (w *BaseWidget) Size() fyne.Size
```
Size gets the current size of this widget.

#### func (*BaseWidget) Visible

```go
func (w *BaseWidget) Visible() bool
```
Visible returns whether or not this widget should be visible. Note that this may
not mean it is currently visible if a parent has been hidden.

#### type Box

```go
type Box struct {
	BaseWidget

	Horizontal bool
	Children   []fyne.CanvasObject
}
```

Box widget is a simple list where the child elements are arranged in a single
column for vertical or a single row for horizontal arrangement

#### func  NewHBox

```go
func NewHBox(children ...fyne.CanvasObject) *Box
```
NewHBox creates a new horizontally aligned box widget with the specified list of
child objects

#### func  NewVBox

```go
func NewVBox(children ...fyne.CanvasObject) *Box
```
NewVBox creates a new vertically aligned box widget with the specified list of
child objects

#### func (*Box) Append

```go
func (b *Box) Append(object fyne.CanvasObject)
```
Append adds a new CanvasObject to the end/right of the box

#### func (*Box) CreateRenderer

```go
func (b *Box) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Box) MinSize

```go
func (b *Box) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Box) Prepend

```go
func (b *Box) Prepend(object fyne.CanvasObject)
```
Prepend inserts a new CanvasObject at the top/left of the box

#### func (*Box) Refresh

```go
func (b *Box) Refresh()
```
Refresh updates this box to match the current theme

#### type Button

```go
type Button struct {
	DisableableWidget
	Text  string
	Style ButtonStyle
	Icon  fyne.Resource

	Alignment     ButtonAlign
	IconPlacement ButtonIconPlacement

	OnTapped func() `json:"-"`

	HideShadow bool
}
```

Button widget has a text label and triggers an event func when clicked

#### func  NewButton

```go
func NewButton(label string, tapped func()) *Button
```
NewButton creates a new button widget with the set label and tap handler

#### func  NewButtonWithIcon

```go
func NewButtonWithIcon(label string, icon fyne.Resource, tapped func()) *Button
```
NewButtonWithIcon creates a new button widget with the specified label, themed
icon and tap handler

#### func (*Button) CreateRenderer

```go
func (b *Button) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Button) Cursor

```go
func (b *Button) Cursor() desktop.Cursor
```
Cursor returns the cursor type of this widget

#### func (*Button) MinSize

```go
func (b *Button) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Button) MouseIn

```go
func (b *Button) MouseIn(*desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

#### func (*Button) MouseMoved

```go
func (b *Button) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

#### func (*Button) MouseOut

```go
func (b *Button) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

#### func (*Button) SetIcon

```go
func (b *Button) SetIcon(icon fyne.Resource)
```
SetIcon updates the icon on a label - pass nil to hide an icon

#### func (*Button) SetText

```go
func (b *Button) SetText(text string)
```
SetText allows the button label to be changed

#### func (*Button) Tapped

```go
func (b *Button) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any tap
handler

#### type ButtonAlign

```go
type ButtonAlign int
```

ButtonAlign represents the horizontal alignment of a button.

```go
const (
	// ButtonAlignCenter aligns the icon and the text centrally.
	ButtonAlignCenter ButtonAlign = iota
	// ButtonAlignLeading aligns the icon and the text with the leading edge.
	ButtonAlignLeading
	// ButtonAlignTrailing aligns the icon and the text with the trailing edge.
	ButtonAlignTrailing
)
```

#### type ButtonIconPlacement

```go
type ButtonIconPlacement int
```

ButtonIconPlacement represents the ordering of icon & text within a button.

```go
const (
	// ButtonIconLeadingText aligns the icon on the leading edge of the text.
	ButtonIconLeadingText ButtonIconPlacement = iota
	// ButtonIconTrailingText aligns the icon on the trailing edge of the text.
	ButtonIconTrailingText
)
```

#### type ButtonStyle

```go
type ButtonStyle int
```

ButtonStyle determines the behaviour and rendering of a button.

```go
const (
	// DefaultButton is the standard button style
	DefaultButton ButtonStyle = iota
	// PrimaryButton that should be more prominent to the user
	PrimaryButton
)
```

#### type Check

```go
type Check struct {
	DisableableWidget
	Text    string
	Checked bool

	OnChanged func(bool) `json:"-"`
}
```

Check widget has a text label and a checked (or unchecked) icon and triggers an
event func when toggled

#### func  NewCheck

```go
func NewCheck(label string, changed func(bool)) *Check
```
NewCheck creates a new check widget with the set label and change handler

#### func (*Check) CreateRenderer

```go
func (c *Check) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Check) FocusGained

```go
func (c *Check) FocusGained()
```
FocusGained is called when the Check has been given focus.

#### func (*Check) FocusLost

```go
func (c *Check) FocusLost()
```
FocusLost is called when the Check has had focus removed.

#### func (*Check) Focused

```go
func (c *Check) Focused() bool
```
Focused returns whether or not this Check has focus. Deprecated: this method
will be removed as it is no longer required, widgets do not expose their focus
state.

#### func (*Check) Hide

```go
func (c *Check) Hide()
```
Hide this widget, if it was previously visible

#### func (*Check) MinSize

```go
func (c *Check) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Check) MouseIn

```go
func (c *Check) MouseIn(*desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

#### func (*Check) MouseMoved

```go
func (c *Check) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

#### func (*Check) MouseOut

```go
func (c *Check) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

#### func (*Check) SetChecked

```go
func (c *Check) SetChecked(checked bool)
```
SetChecked sets the the checked state and refreshes widget

#### func (*Check) Tapped

```go
func (c *Check) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any change
handler

#### func (*Check) TypedKey

```go
func (c *Check) TypedKey(key *fyne.KeyEvent)
```
TypedKey receives key input events when the Check is focused.

#### func (*Check) TypedRune

```go
func (c *Check) TypedRune(r rune)
```
TypedRune receives text input events when the Check is focused.

#### type CustomTextGridStyle

```go
type CustomTextGridStyle struct {
	FGColor, BGColor color.Color
}
```

CustomTextGridStyle is a utility type for those not wanting to define their own
style types.

#### func (*CustomTextGridStyle) BackgroundColor

```go
func (c *CustomTextGridStyle) BackgroundColor() color.Color
```
BackgroundColor is the color a cell should use for the background.

#### func (*CustomTextGridStyle) TextColor

```go
func (c *CustomTextGridStyle) TextColor() color.Color
```
TextColor is the color a cell should use for the text.

#### type DisableableWidget

```go
type DisableableWidget struct {
	BaseWidget
}
```

DisableableWidget describes an extension to BaseWidget which can be disabled.
Disabled widgets should have a visually distinct style when disabled, normally
using theme.DisabledButtonColor.

#### func (*DisableableWidget) Disable

```go
func (w *DisableableWidget) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style
appropriately.

#### func (*DisableableWidget) Disabled

```go
func (w *DisableableWidget) Disabled() bool
```
Disabled returns true if this widget is currently disabled or false if it can
currently be interacted with.

#### func (*DisableableWidget) Enable

```go
func (w *DisableableWidget) Enable()
```
Enable this widget, updating any style or features appropriately.

#### type Entry

```go
type Entry struct {
	DisableableWidget
	sync.RWMutex

	Text        string
	PlaceHolder string
	OnChanged   func(string) `json:"-"`
	Password    bool
	ReadOnly    bool // Deprecated: Use Disable() instead
	MultiLine   bool
	Wrapping    fyne.TextWrap

	CursorRow, CursorColumn int
	OnCursorChanged         func() `json:"-"`

	// ActionItem is a small item which is displayed at the outer right of the entry (like a password revealer)
	ActionItem fyne.CanvasObject
}
```

Entry widget allows simple text to be input when focused.

#### func  NewEntry

```go
func NewEntry() *Entry
```
NewEntry creates a new single line entry widget.

#### func  NewMultiLineEntry

```go
func NewMultiLineEntry() *Entry
```
NewMultiLineEntry creates a new entry that allows multiple lines

#### func  NewPasswordEntry

```go
func NewPasswordEntry() *Entry
```
NewPasswordEntry creates a new entry password widget

#### func (*Entry) CreateRenderer

```go
func (e *Entry) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Entry) Cursor

```go
func (e *Entry) Cursor() desktop.Cursor
```
Cursor returns the cursor type of this widget

#### func (*Entry) Disable

```go
func (e *Entry) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style
appropriately.

#### func (*Entry) Disabled

```go
func (e *Entry) Disabled() bool
```
Disabled satisfies the fyne.Disableable interface.

#### func (*Entry) DoubleTapped

```go
func (e *Entry) DoubleTapped(_ *fyne.PointEvent)
```
DoubleTapped is called when this entry has been double tapped so we should
select text below the pointer

#### func (*Entry) DragEnd

```go
func (e *Entry) DragEnd()
```
DragEnd is called at end of a drag event - currently ignored

#### func (*Entry) Dragged

```go
func (e *Entry) Dragged(d *fyne.DragEvent)
```
Dragged is called when the pointer moves while a button is held down

#### func (*Entry) Enable

```go
func (e *Entry) Enable()
```
Enable this widget, updating any style or features appropriately.

#### func (*Entry) ExtendBaseWidget

```go
func (e *Entry) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget
functionality.

#### func (*Entry) FocusGained

```go
func (e *Entry) FocusGained()
```
FocusGained is called when the Entry has been given focus.

#### func (*Entry) FocusLost

```go
func (e *Entry) FocusLost()
```
FocusLost is called when the Entry has had focus removed.

#### func (*Entry) Focused

```go
func (e *Entry) Focused() bool
```
Focused returns whether or not this Entry has focus. Deprecated: this method
will be removed as it is no longer required, widgets do not expose their focus
state.

#### func (*Entry) Hide

```go
func (e *Entry) Hide()
```
Hide satisfies the fyne.CanvasObject interface.

#### func (*Entry) KeyDown

```go
func (e *Entry) KeyDown(key *fyne.KeyEvent)
```
KeyDown handler for keypress events - used to store shift modifier state for
text selection

#### func (*Entry) KeyUp

```go
func (e *Entry) KeyUp(key *fyne.KeyEvent)
```
KeyUp handler for key release events - used to reset shift modifier state for
text selection

#### func (*Entry) MinSize

```go
func (e *Entry) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Entry) MouseDown

```go
func (e *Entry) MouseDown(m *desktop.MouseEvent)
```
MouseDown called on mouse click, this triggers a mouse click which can move the
cursor, update the existing selection (if shift is held), or start a selection
dragging operation.

#### func (*Entry) MouseUp

```go
func (e *Entry) MouseUp(_ *desktop.MouseEvent)
```
MouseUp called on mouse release If a mouse drag event has completed then check
to see if it has resulted in an empty selection, if so, and if a text select key
isn't held, then disable selecting

#### func (*Entry) SelectedText

```go
func (e *Entry) SelectedText() string
```
SelectedText returns the text currently selected in this Entry. If there is no
selection it will return the empty string.

#### func (*Entry) SetPlaceHolder

```go
func (e *Entry) SetPlaceHolder(text string)
```
SetPlaceHolder sets the text that will be displayed if the entry is otherwise
empty

#### func (*Entry) SetReadOnly

```go
func (e *Entry) SetReadOnly(ro bool)
```
SetReadOnly sets whether or not the Entry should not be editable Deprecated: Use
Disable() instead.

#### func (*Entry) SetText

```go
func (e *Entry) SetText(text string)
```
SetText manually sets the text of the Entry to the given text value.

#### func (*Entry) Tapped

```go
func (e *Entry) Tapped(ev *fyne.PointEvent)
```
Tapped is called when this entry has been tapped so we should update the cursor
position.

#### func (*Entry) TappedSecondary

```go
func (e *Entry) TappedSecondary(pe *fyne.PointEvent)
```
TappedSecondary is called when right or alternative tap is invoked.

Opens the PopUpMenu with `Paste` item to paste text from the clipboard.

#### func (*Entry) TypedKey

```go
func (e *Entry) TypedKey(key *fyne.KeyEvent)
```
TypedKey receives key input events when the Entry widget is focused.

#### func (*Entry) TypedRune

```go
func (e *Entry) TypedRune(r rune)
```
TypedRune receives text input events when the Entry widget is focused.

#### func (*Entry) TypedShortcut

```go
func (e *Entry) TypedShortcut(shortcut fyne.Shortcut)
```
TypedShortcut implements the Shortcutable interface

#### type Form

```go
type Form struct {
	BaseWidget

	Items      []*FormItem
	OnSubmit   func()
	OnCancel   func()
	SubmitText string
	CancelText string
}
```

Form widget is two column grid where each row has a label and a widget (usually
an input). The last row of the grid will contain the appropriate form control
buttons if any should be shown. Setting OnSubmit will set the submit button to
be visible and call back the function when tapped. Setting OnCancel will do the
same for a cancel button.

#### func  NewForm

```go
func NewForm(items ...*FormItem) *Form
```
NewForm creates a new form widget with the specified rows of form items and (if
any of them should be shown) a form controls row at the bottom

#### func (*Form) Append

```go
func (f *Form) Append(text string, widget fyne.CanvasObject)
```
Append adds a new row to the form, using the text as a label next to the
specified Widget

#### func (*Form) AppendItem

```go
func (f *Form) AppendItem(item *FormItem)
```
AppendItem adds the specified row to the end of the Form

#### func (*Form) CreateRenderer

```go
func (f *Form) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Form) MinSize

```go
func (f *Form) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Form) Refresh

```go
func (f *Form) Refresh()
```
Refresh updates the widget state when requested.

#### type FormItem

```go
type FormItem struct {
	Text   string
	Widget fyne.CanvasObject
}
```

FormItem provides the details for a row in a form

#### func  NewFormItem

```go
func NewFormItem(text string, widget fyne.CanvasObject) *FormItem
```
NewFormItem creates a new form item with the specified label text and input
widget

#### type Group

```go
type Group struct {
	BaseWidget

	Text string
}
```

Group widget contains a list of widgets that are grouped under a dividing line
and title at the top.

#### func  NewGroup

```go
func NewGroup(title string, children ...fyne.CanvasObject) *Group
```
NewGroup creates a new grouped list widget with a title and the specified list
of child objects.

#### func  NewGroupWithScroller

```go
func NewGroupWithScroller(title string, children ...fyne.CanvasObject) *Group
```
NewGroupWithScroller creates a new grouped list widget with a title and the
specified list of child objects. This group will scroll when the available space
is less than needed to display the items it contains.

#### func (*Group) Append

```go
func (g *Group) Append(object fyne.CanvasObject)
```
Append adds a new CanvasObject to the end of the group

#### func (*Group) CreateRenderer

```go
func (g *Group) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Group) MinSize

```go
func (g *Group) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Group) Prepend

```go
func (g *Group) Prepend(object fyne.CanvasObject)
```
Prepend inserts a new CanvasObject at the top of the group

#### type Hyperlink

```go
type Hyperlink struct {
	BaseWidget
	Text      string
	URL       *url.URL
	Alignment fyne.TextAlign // The alignment of the Text
	Wrapping  fyne.TextWrap  // The wrapping of the Text
	TextStyle fyne.TextStyle // The style of the hyperlink text
}
```

Hyperlink widget is a text component with appropriate padding and layout. When
clicked, the default web browser should open with a URL

#### func  NewHyperlink

```go
func NewHyperlink(text string, url *url.URL) *Hyperlink
```
NewHyperlink creates a new hyperlink widget with the set text content

#### func  NewHyperlinkWithStyle

```go
func NewHyperlinkWithStyle(text string, url *url.URL, alignment fyne.TextAlign, style fyne.TextStyle) *Hyperlink
```
NewHyperlinkWithStyle creates a new hyperlink widget with the set text content

#### func (*Hyperlink) CreateRenderer

```go
func (hl *Hyperlink) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Hyperlink) Cursor

```go
func (hl *Hyperlink) Cursor() desktop.Cursor
```
Cursor returns the cursor type of this widget

#### func (*Hyperlink) MinSize

```go
func (hl *Hyperlink) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*Hyperlink) Resize

```go
func (hl *Hyperlink) Resize(size fyne.Size)
```
Resize sets a new size for the hyperlink. Note this should not be used if the
widget is being managed by a Layout within a Container.

#### func (*Hyperlink) SetText

```go
func (hl *Hyperlink) SetText(text string)
```
SetText sets the text of the hyperlink

#### func (*Hyperlink) SetURL

```go
func (hl *Hyperlink) SetURL(url *url.URL)
```
SetURL sets the URL of the hyperlink, taking in a URL type

#### func (*Hyperlink) SetURLFromString

```go
func (hl *Hyperlink) SetURLFromString(str string) error
```
SetURLFromString sets the URL of the hyperlink, taking in a string type

#### func (*Hyperlink) Tapped

```go
func (hl *Hyperlink) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any change
handler

#### type Icon

```go
type Icon struct {
	BaseWidget

	Resource fyne.Resource // The resource for this icon
}
```

Icon widget is a basic image component that load's its resource to match the
theme.

#### func  NewIcon

```go
func NewIcon(res fyne.Resource) *Icon
```
NewIcon returns a new icon widget that displays a themed icon resource

#### func (*Icon) CreateRenderer

```go
func (i *Icon) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Icon) MinSize

```go
func (i *Icon) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Icon) SetResource

```go
func (i *Icon) SetResource(res fyne.Resource)
```
SetResource updates the resource rendered in this icon widget

#### type Label

```go
type Label struct {
	BaseWidget
	Text      string
	Alignment fyne.TextAlign // The alignment of the Text
	Wrapping  fyne.TextWrap  // The wrapping of the Text
	TextStyle fyne.TextStyle // The style of the label text
}
```

Label widget is a label component with appropriate padding and layout.

#### func  NewLabel

```go
func NewLabel(text string) *Label
```
NewLabel creates a new label widget with the set text content

#### func  NewLabelWithStyle

```go
func NewLabelWithStyle(text string, alignment fyne.TextAlign, style fyne.TextStyle) *Label
```
NewLabelWithStyle creates a new label widget with the set text content

#### func (*Label) CreateRenderer

```go
func (l *Label) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Label) MinSize

```go
func (l *Label) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Label) Refresh

```go
func (l *Label) Refresh()
```
Refresh checks if the text content should be updated then refreshes the
graphical context

#### func (*Label) Resize

```go
func (l *Label) Resize(size fyne.Size)
```
Resize sets a new size for the label. Note this should not be used if the widget
is being managed by a Layout within a Container.

#### func (*Label) SetText

```go
func (l *Label) SetText(text string)
```
SetText sets the text of the label

#### type Orientation

```go
type Orientation int
```

Orientation controls the horizontal/vertical layout of a widget

```go
const (
	Horizontal Orientation = 0
	Vertical   Orientation = 1
)
```
Orientation constants to control widget layout

#### type PopUp

```go
type PopUp struct {
	BaseWidget

	Content fyne.CanvasObject
	Canvas  fyne.Canvas
}
```

PopUp is a widget that can float above the user interface. It wraps any standard
elements with padding and a shadow. If it is modal then the shadow will cover
the entire canvas it hovers over and block interactions.

#### func  NewModalPopUp

```go
func NewModalPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp
```
NewModalPopUp creates a new popUp for the specified content and displays it on
the passed canvas. A modal PopUp blocks interactions with underlying elements,
covered with a semi-transparent overlay. Deprecated: This will no longer show
the pop-up in 2.0. Use ShowModalPopUp instead.

#### func  NewPopUp

```go
func NewPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *PopUp
```
NewPopUp creates a new popUp for the specified content and displays it on the
passed canvas. Deprecated: This will no longer show the pop-up in 2.0. Use
ShowPopUp() instead.

#### func  NewPopUpAtPosition

```go
func NewPopUpAtPosition(content fyne.CanvasObject, canvas fyne.Canvas, pos fyne.Position) *PopUp
```
NewPopUpAtPosition creates a new popUp for the specified content at the
specified absolute position. It will then display the popup on the passed
canvas. Deprecated: Use ShowPopUpAtPosition() instead.

#### func  NewPopUpMenu

```go
func NewPopUpMenu(menu *fyne.Menu, c fyne.Canvas) *PopUp
```
NewPopUpMenu creates a PopUp widget populated with menu items from the passed
menu structure. It will automatically be shown as an overlay on the specified
canvas. Deprecated: Use ShowPopUpMenuAtPosition instead.

#### func  NewPopUpMenuAtPosition

```go
func NewPopUpMenuAtPosition(menu *fyne.Menu, c fyne.Canvas, pos fyne.Position) *PopUp
```
NewPopUpMenuAtPosition creates a PopUp widget populated with menu items from the
passed menu structure. It will automatically be positioned at the provided
location and shown as an overlay on the specified canvas. Deprecated: Use
ShowPopUpMenuAtPosition instead.

#### func (*PopUp) CreateRenderer

```go
func (p *PopUp) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*PopUp) Hide

```go
func (p *PopUp) Hide()
```
Hide this widget, if it was previously visible

#### func (*PopUp) MinSize

```go
func (p *PopUp) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*PopUp) Move

```go
func (p *PopUp) Move(pos fyne.Position)
```
Move the widget to a new position. A PopUp position is absolute to the top, left
of its canvas. For PopUp this actually moves the content so checking Position()
will not return the same value as is set here.

#### func (*PopUp) Resize

```go
func (p *PopUp) Resize(size fyne.Size)
```
Resize satisfies the fyne.CanvasObject interface. PopUps always have the size of
their canvas. However, Resize changes the size of the PopUp's content.

#### func (*PopUp) Show

```go
func (p *PopUp) Show()
```
Show this pop-up as overlay if not already shown.

#### func (*PopUp) ShowAtPosition

```go
func (p *PopUp) ShowAtPosition(pos fyne.Position)
```
ShowAtPosition shows this pop-up at the given position.

#### func (*PopUp) Tapped

```go
func (p *PopUp) Tapped(_ *fyne.PointEvent)
```
Tapped is called when the user taps the popUp background - if not modal then
dismiss this widget

#### func (*PopUp) TappedSecondary

```go
func (p *PopUp) TappedSecondary(_ *fyne.PointEvent)
```
TappedSecondary is called when the user right/alt taps the background - if not
modal then dismiss this widget

#### type ProgressBar

```go
type ProgressBar struct {
	BaseWidget

	Min, Max, Value float64
}
```

ProgressBar widget creates a horizontal panel that indicates progress

#### func  NewProgressBar

```go
func NewProgressBar() *ProgressBar
```
NewProgressBar creates a new progress bar widget. The default Min is 0 and Max
is 1, Values set should be between those numbers. The display will convert this
to a percentage.

#### func (*ProgressBar) CreateRenderer

```go
func (p *ProgressBar) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*ProgressBar) MinSize

```go
func (p *ProgressBar) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*ProgressBar) SetValue

```go
func (p *ProgressBar) SetValue(v float64)
```
SetValue changes the current value of this progress bar (from p.Min to p.Max).
The widget will be refreshed to indicate the change.

#### type ProgressBarInfinite

```go
type ProgressBarInfinite struct {
	BaseWidget
}
```

ProgressBarInfinite widget creates a horizontal panel that indicates waiting
indefinitely An infinite progress bar loops 0% -> 100% repeatedly until Stop()
is called

#### func  NewProgressBarInfinite

```go
func NewProgressBarInfinite() *ProgressBarInfinite
```
NewProgressBarInfinite creates a new progress bar widget that loops indefinitely
from 0% -> 100% SetValue() is not defined for infinite progress bar To stop the
looping progress and set the progress bar to 100%, call
ProgressBarInfinite.Stop()

#### func (*ProgressBarInfinite) CreateRenderer

```go
func (p *ProgressBarInfinite) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*ProgressBarInfinite) Hide

```go
func (p *ProgressBarInfinite) Hide()
```
Hide this widget, if it was previously visible

#### func (*ProgressBarInfinite) MinSize

```go
func (p *ProgressBarInfinite) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*ProgressBarInfinite) Running

```go
func (p *ProgressBarInfinite) Running() bool
```
Running returns the current state of the infinite progress animation

#### func (*ProgressBarInfinite) Show

```go
func (p *ProgressBarInfinite) Show()
```
Show this widget, if it was previously hidden

#### func (*ProgressBarInfinite) Start

```go
func (p *ProgressBarInfinite) Start()
```
Start the infinite progress bar background thread to update it continuously

#### func (*ProgressBarInfinite) Stop

```go
func (p *ProgressBarInfinite) Stop()
```
Stop the infinite progress goroutine and sets value to the Max

#### type Radio

```go
type Radio struct {
	DisableableWidget
	Horizontal bool
	Required   bool
	OnChanged  func(string) `json:"-"`
	Options    []string
	Selected   string
}
```

Radio widget has a list of text labels and radio check icons next to each.
Changing the selection (only one can be selected) will trigger the changed func.

#### func  NewRadio

```go
func NewRadio(options []string, changed func(string)) *Radio
```
NewRadio creates a new radio widget with the set options and change handler

#### func (*Radio) Append

```go
func (r *Radio) Append(option string)
```
Append adds a new option to the end of a Radio widget.

#### func (*Radio) CreateRenderer

```go
func (r *Radio) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Radio) MinSize

```go
func (r *Radio) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Radio) MouseIn

```go
func (r *Radio) MouseIn(event *desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

#### func (*Radio) MouseMoved

```go
func (r *Radio) MouseMoved(event *desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

#### func (*Radio) MouseOut

```go
func (r *Radio) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

#### func (*Radio) SetSelected

```go
func (r *Radio) SetSelected(option string)
```
SetSelected sets the radio option, it can be used to set a default option.

#### func (*Radio) Tapped

```go
func (r *Radio) Tapped(event *fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any change
handler

#### type ScrollContainer

```go
type ScrollContainer struct {
	BaseWidget

	Direction ScrollDirection
	Content   fyne.CanvasObject
	Offset    fyne.Position
}
```

ScrollContainer defines a container that is smaller than the Content. The Offset
is used to determine the position of the child widgets within the container.

#### func  NewHScrollContainer

```go
func NewHScrollContainer(content fyne.CanvasObject) *ScrollContainer
```
NewHScrollContainer create a scrollable parent wrapping the specified content.
Note that this may cause the MinSize.Width to be smaller than that of the passed
object.

#### func  NewScrollContainer

```go
func NewScrollContainer(content fyne.CanvasObject) *ScrollContainer
```
NewScrollContainer creates a scrollable parent wrapping the specified content.
Note that this may cause the MinSize to be smaller than that of the passed
object.

#### func  NewVScrollContainer

```go
func NewVScrollContainer(content fyne.CanvasObject) *ScrollContainer
```
NewVScrollContainer create a scrollable parent wrapping the specified content.
Note that this may cause the MinSize.Height to be smaller than that of the
passed object.

#### func (*ScrollContainer) CreateRenderer

```go
func (s *ScrollContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*ScrollContainer) DragEnd

```go
func (s *ScrollContainer) DragEnd()
```
DragEnd will stop scrolling on mobile has stopped

#### func (*ScrollContainer) Dragged

```go
func (s *ScrollContainer) Dragged(e *fyne.DragEvent)
```
Dragged will scroll on any drag - bar or otherwise - for mobile

#### func (*ScrollContainer) MinSize

```go
func (s *ScrollContainer) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*ScrollContainer) Refresh

```go
func (s *ScrollContainer) Refresh()
```
Refresh causes this widget to be redrawn in it's current state

#### func (*ScrollContainer) Resize

```go
func (s *ScrollContainer) Resize(size fyne.Size)
```
Resize sets a new size for the scroll container.

#### func (*ScrollContainer) Scrolled

```go
func (s *ScrollContainer) Scrolled(ev *fyne.ScrollEvent)
```
Scrolled is called when an input device triggers a scroll event

#### func (*ScrollContainer) SetMinSize

```go
func (s *ScrollContainer) SetMinSize(size fyne.Size)
```
SetMinSize specifies a minimum size for this scroll container. If the specified
size is larger than the content size then scrolling will not be enabled This can
be helpful to set scrolling in only 1 direction.

#### type ScrollDirection

```go
type ScrollDirection int
```

ScrollDirection represents the directions in which a ScrollContainer can scroll
its child content.

```go
const (
	ScrollBoth ScrollDirection = iota
	ScrollHorizontalOnly
	ScrollVerticalOnly
)
```
Constants for valid values of ScrollDirection.

#### type Select

```go
type Select struct {
	BaseWidget

	Selected    string
	Options     []string
	PlaceHolder string
	OnChanged   func(string) `json:"-"`
}
```

Select widget has a list of options, with the current one shown, and triggers an
event func when clicked

#### func  NewSelect

```go
func NewSelect(options []string, changed func(string)) *Select
```
NewSelect creates a new select widget with the set list of options and changes
handler

#### func (*Select) ClearSelected

```go
func (s *Select) ClearSelected()
```
ClearSelected clears the current option of the select widget. After clearing the
current option, the Select widget's PlaceHolder will be displayed.

#### func (*Select) CreateRenderer

```go
func (s *Select) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Select) Hide

```go
func (s *Select) Hide()
```
Hide satisfies the fyne.CanvasObject interface.

#### func (*Select) MinSize

```go
func (s *Select) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Select) MouseIn

```go
func (s *Select) MouseIn(*desktop.MouseEvent)
```
MouseIn is called when a desktop pointer enters the widget

#### func (*Select) MouseMoved

```go
func (s *Select) MouseMoved(*desktop.MouseEvent)
```
MouseMoved is called when a desktop pointer hovers over the widget

#### func (*Select) MouseOut

```go
func (s *Select) MouseOut()
```
MouseOut is called when a desktop pointer exits the widget

#### func (*Select) Move

```go
func (s *Select) Move(pos fyne.Position)
```
Move satisfies the fyne.CanvasObject interface.

#### func (*Select) Resize

```go
func (s *Select) Resize(size fyne.Size)
```
Resize sets a new size for a widget. Note this should not be used if the widget
is being managed by a Layout within a Container.

#### func (*Select) SetSelected

```go
func (s *Select) SetSelected(text string)
```
SetSelected sets the current option of the select widget

#### func (*Select) Tapped

```go
func (s *Select) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any tap
handler

#### type SelectEntry

```go
type SelectEntry struct {
	Entry
}
```

SelectEntry is an input field which supports selecting from a fixed set of
options.

#### func  NewSelectEntry

```go
func NewSelectEntry(options []string) *SelectEntry
```
NewSelectEntry creates a SelectEntry.

#### func (*SelectEntry) MinSize

```go
func (e *SelectEntry) MinSize() fyne.Size
```
MinSize satisfies the fyne.CanvasObject interface.

#### func (*SelectEntry) Resize

```go
func (e *SelectEntry) Resize(size fyne.Size)
```
Resize satisfies the fyne.CanvasObject interface.

#### func (*SelectEntry) SetOptions

```go
func (e *SelectEntry) SetOptions(options []string)
```
SetOptions sets the options the user might select from.

#### type Slider

```go
type Slider struct {
	BaseWidget

	Value float64
	Min   float64
	Max   float64
	Step  float64

	Orientation Orientation
	OnChanged   func(float64)
}
```

Slider is a widget that can slide between two fixed values.

#### func  NewSlider

```go
func NewSlider(min, max float64) *Slider
```
NewSlider returns a basic slider.

#### func (*Slider) CreateRenderer

```go
func (s *Slider) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer links this widget to its renderer.

#### func (*Slider) DragEnd

```go
func (s *Slider) DragEnd()
```
DragEnd function.

#### func (*Slider) Dragged

```go
func (s *Slider) Dragged(e *fyne.DragEvent)
```
Dragged function.

#### func (*Slider) MinSize

```go
func (s *Slider) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### type SplitContainer

```go
type SplitContainer struct {
	BaseWidget
	Offset     float64
	Horizontal bool
	Leading    fyne.CanvasObject
	Trailing   fyne.CanvasObject
}
```

SplitContainer defines a container whose size is split between two children.

#### func  NewHSplitContainer

```go
func NewHSplitContainer(left, right fyne.CanvasObject) *SplitContainer
```
NewHSplitContainer create a splitable parent wrapping the specified children.

#### func  NewVSplitContainer

```go
func NewVSplitContainer(top, bottom fyne.CanvasObject) *SplitContainer
```
NewVSplitContainer create a splitable parent wrapping the specified children.

#### func (*SplitContainer) CreateRenderer

```go
func (s *SplitContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*SplitContainer) SetOffset

```go
func (s *SplitContainer) SetOffset(offset float64)
```
SetOffset sets the offset (0.0 to 1.0) of the SplitContainer divider. 0.0 -
Leading is min size, Trailing uses all remaining space. 0.5 - Leading & Trailing
equally share the available space. 1.0 - Trailing is min size, Leading uses all
remaining space.

#### type TabContainer

```go
type TabContainer struct {
	BaseWidget

	Items     []*TabItem
	OnChanged func(tab *TabItem)
}
```

TabContainer widget allows switching visible content from a list of TabItems.
Each item is represented by a button at the top of the widget.

#### func  NewTabContainer

```go
func NewTabContainer(items ...*TabItem) *TabContainer
```
NewTabContainer creates a new tab bar widget that allows the user to choose
between different visible containers

#### func (*TabContainer) Append

```go
func (t *TabContainer) Append(item *TabItem)
```
Append adds a new TabItem to the rightmost side of the tab panel

#### func (*TabContainer) CreateRenderer

```go
func (t *TabContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*TabContainer) CurrentTab

```go
func (t *TabContainer) CurrentTab() *TabItem
```
CurrentTab returns the currently selected TabItem.

#### func (*TabContainer) CurrentTabIndex

```go
func (t *TabContainer) CurrentTabIndex() int
```
CurrentTabIndex returns the index of the currently selected TabItem.

#### func (*TabContainer) MinSize

```go
func (t *TabContainer) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*TabContainer) Remove

```go
func (t *TabContainer) Remove(item *TabItem)
```
Remove tab by value

#### func (*TabContainer) RemoveIndex

```go
func (t *TabContainer) RemoveIndex(index int)
```
RemoveIndex removes tab by index

#### func (*TabContainer) SelectTab

```go
func (t *TabContainer) SelectTab(item *TabItem)
```
SelectTab sets the specified TabItem to be selected and its content visible.

#### func (*TabContainer) SelectTabIndex

```go
func (t *TabContainer) SelectTabIndex(index int)
```
SelectTabIndex sets the TabItem at the specific index to be selected and its
content visible.

#### func (*TabContainer) SetTabLocation

```go
func (t *TabContainer) SetTabLocation(l TabLocation)
```
SetTabLocation sets the location of the tab bar

#### func (*TabContainer) Show

```go
func (t *TabContainer) Show()
```
Show this widget, if it was previously hidden

#### type TabItem

```go
type TabItem struct {
	Text    string
	Icon    fyne.Resource
	Content fyne.CanvasObject
}
```

TabItem represents a single view in a TabContainer. The Text and Icon are used
for the tab button and the Content is shown when the corresponding tab is
active.

#### func  NewTabItem

```go
func NewTabItem(text string, content fyne.CanvasObject) *TabItem
```
NewTabItem creates a new item for a tabbed widget - each item specifies the
content and a label for its tab.

#### func  NewTabItemWithIcon

```go
func NewTabItemWithIcon(text string, icon fyne.Resource, content fyne.CanvasObject) *TabItem
```
NewTabItemWithIcon creates a new item for a tabbed widget - each item specifies
the content and a label with an icon for its tab.

#### type TabLocation

```go
type TabLocation int
```

TabLocation is the location where the tabs of a tab container should be rendered

```go
const (
	TabLocationTop TabLocation = iota
	TabLocationLeading
	TabLocationBottom
	TabLocationTrailing
)
```
TabLocation values

#### type TextGrid

```go
type TextGrid struct {
	BaseWidget
	Rows []TextGridRow

	ShowLineNumbers bool
	ShowWhitespace  bool
}
```

TextGrid is a monospaced grid of characters. This is designed to be used by a
text editor, code preview or terminal emulator.

#### func  NewTextGrid

```go
func NewTextGrid() *TextGrid
```
NewTextGrid creates a new empty TextGrid widget.

#### func  NewTextGridFromString

```go
func NewTextGridFromString(content string) *TextGrid
```
NewTextGridFromString creates a new TextGrid widget with the specified string
content.

#### func (*TextGrid) CreateRenderer

```go
func (t *TextGrid) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to it's
renderer

#### func (*TextGrid) MinSize

```go
func (t *TextGrid) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*TextGrid) Resize

```go
func (t *TextGrid) Resize(size fyne.Size)
```
Resize is called when this widget changes size. We should make sure that we
refresh cells.

#### func (*TextGrid) Row

```go
func (t *TextGrid) Row(row int) TextGridRow
```
Row returns the content of a specified row as a TextGridRow. If the index is out
of bounds it returns an empty row object.

#### func (*TextGrid) SetRow

```go
func (t *TextGrid) SetRow(row int, content TextGridRow)
```
SetRow updates the specified row of the grid's contents using the specified
content and style and then refreshes. If the row is beyond the end of the
current buffer it will be expanded.

#### func (*TextGrid) SetStyle

```go
func (t *TextGrid) SetStyle(row, col int, style TextGridStyle)
```
SetStyle sets a grid style to the cell at named row and column

#### func (*TextGrid) SetStyleRange

```go
func (t *TextGrid) SetStyleRange(startRow, startCol, endRow, endCol int, style TextGridStyle)
```
SetStyleRange sets a grid style to all the cells between the start row and
column through to the end row and column.

#### func (*TextGrid) SetText

```go
func (t *TextGrid) SetText(text string)
```
SetText updates the buffer of this textgrid to contain the specified text. New
lines and columns will be added as required. Lines are separated by '\n'. The
grid will use default text style and any previous content and style will be
removed.

#### func (*TextGrid) Text

```go
func (t *TextGrid) Text() string
```
Text returns the contents of the buffer as a single string (with no style
information). It reconstructs the lines by joining with a `\n` character.

#### type TextGridCell

```go
type TextGridCell struct {
	Rune  rune
	Style TextGridStyle
}
```

TextGridCell represents a single cell in a text grid. It has a rune for the text
content and a style associated with it.

#### type TextGridRow

```go
type TextGridRow struct {
	Cells []TextGridCell
	Style TextGridStyle
}
```

TextGridRow represents a row of cells cell in a text grid. It contains the cells
for the row and an optional style.

#### type TextGridStyle

```go
type TextGridStyle interface {
	TextColor() color.Color
	BackgroundColor() color.Color
}
```

TextGridStyle defines a style that can be applied to a TextGrid cell.

```go
var (
	// TextGridStyleDefault is a default style for test grid cells
	TextGridStyleDefault TextGridStyle
	// TextGridStyleWhitespace is the style used for whitespace characters, if enabled
	TextGridStyleWhitespace TextGridStyle
)
```

#### type Toolbar

```go
type Toolbar struct {
	BaseWidget
	Items []ToolbarItem
}
```

Toolbar widget creates a horizontal list of tool buttons

#### func  NewToolbar

```go
func NewToolbar(items ...ToolbarItem) *Toolbar
```
NewToolbar creates a new toolbar widget.

#### func (*Toolbar) Append

```go
func (t *Toolbar) Append(item ToolbarItem)
```
Append a new ToolbarItem to the end of this Toolbar

#### func (*Toolbar) CreateRenderer

```go
func (t *Toolbar) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its
renderer

#### func (*Toolbar) MinSize

```go
func (t *Toolbar) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Toolbar) Prepend

```go
func (t *Toolbar) Prepend(item ToolbarItem)
```
Prepend a new ToolbarItem to the start of this Toolbar

#### type ToolbarAction

```go
type ToolbarAction struct {
	Icon        fyne.Resource
	OnActivated func()
}
```

ToolbarAction is push button style of ToolbarItem

#### func (*ToolbarAction) ToolbarObject

```go
func (t *ToolbarAction) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets a button to render this ToolbarAction

#### type ToolbarItem

```go
type ToolbarItem interface {
	ToolbarObject() fyne.CanvasObject
}
```

ToolbarItem represents any interface element that can be added to a toolbar

#### func  NewToolbarAction

```go
func NewToolbarAction(icon fyne.Resource, onActivated func()) ToolbarItem
```
NewToolbarAction returns a new push button style ToolbarItem

#### func  NewToolbarSeparator

```go
func NewToolbarSeparator() ToolbarItem
```
NewToolbarSeparator returns a new separator item for a Toolbar to assist with
ToolbarItem grouping

#### func  NewToolbarSpacer

```go
func NewToolbarSpacer() ToolbarItem
```
NewToolbarSpacer returns a new spacer item for a Toolbar to assist with
ToolbarItem alignment

#### type ToolbarSeparator

```go
type ToolbarSeparator struct {
}
```

ToolbarSeparator is a thin, visible divide that can be added to a Toolbar. This
is typically used to assist visual grouping of ToolbarItems.

#### func (*ToolbarSeparator) ToolbarObject

```go
func (t *ToolbarSeparator) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the visible line object for this ToolbarSeparator

#### type ToolbarSpacer

```go
type ToolbarSpacer struct {
}
```

ToolbarSpacer is a blank, stretchable space for a toolbar. This is typically
used to assist layout if you wish some left and some right aligned items. Space
will be split evebly amongst all the spacers on a toolbar.

#### func (*ToolbarSpacer) ToolbarObject

```go
func (t *ToolbarSpacer) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the actual spacer object for this ToolbarSpacer
