---
redirect_to:
  - https://docs.fyne.io/api/v1.4/widget/radio

layout: page
tags: [api]
title: Fyne API "widget.Radio"
---


# widget.Radio
---
```go
import "fyne.io/fyne/widget"
```

## Usage

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

Radio widget has a list of text labels and radio check icons next to each. Changing the selection (only one can be selected) will trigger the changed func.


<div class="deprecated">
Deprecated: Use RadioGroup instead.</div>

#### func  NewRadio

```go
func NewRadio(options []string, changed func(string)) *Radio
```
NewRadio creates a new radio widget with the set options and change handler


<div class="deprecated">
Deprecated: Use NewRadioGroup instead.</div>

#### func (*Radio) Append

```go
func (r *Radio) Append(option string)
```
Append adds a new option to the end of a Radio widget.

#### func (*Radio) CreateRenderer

```go
func (r *Radio) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

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
Tapped is called when a pointer tapped event is captured and triggers any change handler
