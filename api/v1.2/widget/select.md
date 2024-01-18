---
redirect_to:
  - https://docs.fyne.io/api/v1.2/widget/select.md

layout: page
tags: [api]
title: Fyne API widget
---


# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

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

Select widget has a list of options, with the current one shown, and triggers an event func when clicked

#### func  NewSelect

```go
func NewSelect(options []string, changed func(string)) *Select
```
NewSelect creates a new select widget with the set list of options and changes handler

#### func (*Select) CreateRenderer

```go
func (s *Select) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

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

#### func (*Select) Resize

```go
func (s *Select) Resize(size fyne.Size)
```
Resize sets a new size for a widget. Note this should not be used if the widget is being managed by a Layout within a Container.

#### func (*Select) SetSelected

```go
func (s *Select) SetSelected(text string)
```
SetSelected sets the current option of the select widget

#### func (*Select) Tapped

```go
func (s *Select) Tapped(*fyne.PointEvent)
```
Tapped is called when a pointer tapped event is captured and triggers any tap handler

#### func (*Select) TappedSecondary

```go
func (s *Select) TappedSecondary(*fyne.PointEvent)
```
TappedSecondary is called when a secondary pointer tapped event is captured
