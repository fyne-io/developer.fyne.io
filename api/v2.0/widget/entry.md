---
redirect_to:
  - https://docs.fyne.io/api/v2.0/widget/entry.md

layout: page
tags: [api]
title: Fyne API "widget.Entry"
---


# widget.Entry
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Entry

```go
type Entry struct {
	DisableableWidget

	Text string
	// Since: 2.0
	TextStyle   fyne.TextStyle
	PlaceHolder string
	OnChanged   func(string) `json:"-"`
	// Since: 2.0
	OnSubmitted func(string) `json:"-"`
	Password    bool
	MultiLine   bool
	Wrapping    fyne.TextWrap

	// Set a validator that this entry will check against
	// Since: 1.4
	Validator fyne.StringValidator

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

#### func  NewEntryWithData

```go
func NewEntryWithData(data binding.String) *Entry
```
NewEntryWithData returns an Entry widget connected to the specified data source.


<div class="since">Since: <code>
2.0</code></div>

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

#### func (*Entry) Bind

```go
func (e *Entry) Bind(data binding.String)
```
Bind connects the specified data source to this Entry. The current value will be displayed and any changes in the data will cause the widget to update. User interactions with this Entry will set the value into the data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Entry) CreateRenderer

```go
func (e *Entry) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Entry) Cursor

```go
func (e *Entry) Cursor() desktop.Cursor
```
Cursor returns the cursor type of this widget


<div class="implements">Implements: <code>
desktop.Cursorable</code></div>

#### func (*Entry) Disable

```go
func (e *Entry) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style appropriately.


<div class="implements">Implements: <code>
fyne.Disableable</code></div>

#### func (*Entry) Disabled

```go
func (e *Entry) Disabled() bool
```
Disabled returns whether the entry is disabled or read-only.


<div class="implements">Implements: <code>
fyne.Disableable</code></div>

#### func (*Entry) DoubleTapped

```go
func (e *Entry) DoubleTapped(p *fyne.PointEvent)
```
DoubleTapped is called when this entry has been double tapped so we should select text below the pointer


<div class="implements">Implements: <code>
fyne.DoubleTappable</code></div>

#### func (*Entry) DragEnd

```go
func (e *Entry) DragEnd()
```
DragEnd is called at end of a drag event.


<div class="implements">Implements: <code>
fyne.Draggable</code></div>

#### func (*Entry) Dragged

```go
func (e *Entry) Dragged(d *fyne.DragEvent)
```
Dragged is called when the pointer moves while a button is held down. It updates the selection accordingly.


<div class="implements">Implements: <code>
fyne.Draggable</code></div>

#### func (*Entry) Enable

```go
func (e *Entry) Enable()
```
Enable this widget, updating any style or features appropriately.


<div class="implements">Implements: <code>
fyne.Disableable</code></div>

#### func (*Entry) ExtendBaseWidget

```go
func (e *Entry) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.

#### func (*Entry) FocusGained

```go
func (e *Entry) FocusGained()
```
FocusGained is called when the Entry has been given focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Entry) FocusLost

```go
func (e *Entry) FocusLost()
```
FocusLost is called when the Entry has had focus removed.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Entry) Hide

```go
func (e *Entry) Hide()
```
Hide hides the entry.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Entry) KeyDown

```go
func (e *Entry) KeyDown(key *fyne.KeyEvent)
```
KeyDown handler for keypress events - used to store shift modifier state for text selection


<div class="implements">Implements: <code>
desktop.Keyable</code></div>

#### func (*Entry) KeyUp

```go
func (e *Entry) KeyUp(key *fyne.KeyEvent)
```
KeyUp handler for key release events - used to reset shift modifier state for text selection


<div class="implements">Implements: <code>
desktop.Keyable</code></div>

#### func (*Entry) Keyboard

```go
func (e *Entry) Keyboard() mobile.KeyboardType
```
Keyboard implements the Keyboardable interface


<div class="implements">Implements: <code>
mobile.Keyboardable</code></div>

#### func (*Entry) MinSize

```go
func (e *Entry) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Entry) MouseDown

```go
func (e *Entry) MouseDown(m *desktop.MouseEvent)
```
MouseDown called on mouse click, this triggers a mouse click which can move the cursor, update the existing selection (if shift is held), or start a selection dragging operation.


<div class="implements">Implements: <code>
desktop.Mouseable</code></div>

#### func (*Entry) MouseUp

```go
func (e *Entry) MouseUp(m *desktop.MouseEvent)
```
MouseUp called on mouse release If a mouse drag event has completed then check to see if it has resulted in an empty selection, if so, and if a text select key isn't held, then disable selecting


<div class="implements">Implements: <code>
desktop.Mouseable</code></div>

#### func (*Entry) SelectedText

```go
func (e *Entry) SelectedText() string
```
SelectedText returns the text currently selected in this Entry. If there is no selection it will return the empty string.

#### func (*Entry) SetOnValidationChanged

```go
func (e *Entry) SetOnValidationChanged(callback func(error))
```
SetOnValidationChanged is intended for parent widgets or containers to hook into the validation. The function might be overwritten by a parent that cares about child validation (e.g. widget.Form).

#### func (*Entry) SetPlaceHolder

```go
func (e *Entry) SetPlaceHolder(text string)
```
SetPlaceHolder sets the text that will be displayed if the entry is otherwise empty

#### func (*Entry) SetText

```go
func (e *Entry) SetText(text string)
```
SetText manually sets the text of the Entry to the given text value.

#### func (*Entry) SetValidationError

```go
func (e *Entry) SetValidationError(err error)
```
SetValidationError manually updates the validation status until the next input change

#### func (*Entry) Tapped

```go
func (e *Entry) Tapped(ev *fyne.PointEvent)
```
Tapped is called when this entry has been tapped so we should update the cursor position.


<div class="implements">Implements: <code>
fyne.Tappable</code></div>

#### func (*Entry) TappedSecondary

```go
func (e *Entry) TappedSecondary(pe *fyne.PointEvent)
```
TappedSecondary is called when right or alternative tap is invoked.

Opens the PopUpMenu with `Paste` item to paste text from the clipboard.


<div class="implements">Implements: <code>
fyne.SecondaryTappable</code></div>

#### func (*Entry) TypedKey

```go
func (e *Entry) TypedKey(key *fyne.KeyEvent)
```
TypedKey receives key input events when the Entry widget is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Entry) TypedRune

```go
func (e *Entry) TypedRune(r rune)
```
TypedRune receives text input events when the Entry widget is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Entry) TypedShortcut

```go
func (e *Entry) TypedShortcut(shortcut fyne.Shortcut)
```
TypedShortcut implements the Shortcutable interface


<div class="implements">Implements: <code>
fyne.Shortcutable</code></div>

#### func (*Entry) Unbind

```go
func (e *Entry) Unbind()
```
Unbind disconnects any configured data source from this Entry. The current value will remain at the last value of the data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Entry) Validate

```go
func (e *Entry) Validate() error
```
Validate validates the current text in the widget
