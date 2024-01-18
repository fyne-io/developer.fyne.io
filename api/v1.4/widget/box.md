---
redirect_to:
  - https://docs.fyne.io/api/v1.4/widget/box

layout: page
tags: [api]
title: Fyne API "widget.Box"
---


# widget.Box
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type Box

```go
type Box struct {
	BaseWidget

	Horizontal bool
	Children   []fyne.CanvasObject
}
```

Box widget is a simple list where the child elements are arranged in a single column for vertical or a single row for horizontal arrangement.

<div class="deprecated"> Deprecated: Use container.NewVBox() or container.NewHBox().</div>

#### func  NewHBox

```go
func NewHBox(children ...fyne.CanvasObject) *Box
```
NewHBox creates a new horizontally aligned box widget with the specified list of child objects.

<div class="deprecated"> Deprecated: Use container.NewHBox() instead.</div>

#### func  NewVBox

```go
func NewVBox(children ...fyne.CanvasObject) *Box
```
NewVBox creates a new vertically aligned box widget with the specified list of child objects.

<div class="deprecated"> Deprecated: Use container.NewVBox() instead.</div>

#### func (*Box) Append

```go
func (b *Box) Append(object fyne.CanvasObject)
```
Append adds a new CanvasObject to the end/right of the box

#### func (*Box) CreateRenderer

```go
func (b *Box) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

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
