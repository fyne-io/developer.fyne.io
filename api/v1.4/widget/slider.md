---
redirect_to:
  - https://docs.fyne.io/api/v1.4/widget/slider

layout: page
tags: [api]
title: Fyne API "widget.Slider"
---


# widget.Slider
---
```go
import "fyne.io/fyne/widget"
```

## Usage

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

#### func (*Slider) SetValue

```go
func (s *Slider) SetValue(value float64)
```
SetValue updates the value of the slider and clamps the value to be within the range.
