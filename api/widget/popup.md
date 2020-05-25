---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
--
    import "fyne.io/fyne/widget"

## Usage

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
Resize changes the size of the PopUp. PopUps always have the size of their
canvas. However, Resize changes the size of the PopUp's content. Implements:
fyne.Widget

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
