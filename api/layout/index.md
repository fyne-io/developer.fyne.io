---
layout: page
tags: [api]
title: Fyne API layout
---

# layout
--
    import "fyne.io/fyne/layout"

Package layout defines the various layouts available to Fyne apps

## Usage

#### func  NewAdaptiveGridLayout

```go
func NewAdaptiveGridLayout(rowcols int) fyne.Layout
```
NewAdaptiveGridLayout returns a new grid layout which uses columns when
horizontal but rows when vertical.

#### func  NewBorderLayout

```go
func NewBorderLayout(top, bottom, left, right fyne.CanvasObject) fyne.Layout
```
NewBorderLayout creates a new BorderLayout instance with top, left, bottom and
right objects set. All other items in the container will fill the centre space

#### func  NewCenterLayout

```go
func NewCenterLayout() fyne.Layout
```
NewCenterLayout creates a new CenterLayout instance

#### func  NewFixedGridLayout

```go
func NewFixedGridLayout(size fyne.Size) fyne.Layout
```
NewFixedGridLayout returns a new FixedGridLayout instance Deprecated: use the
replacement NewGridWrapLayout. This method will be removed in 2.0.

#### func  NewFormLayout

```go
func NewFormLayout() fyne.Layout
```
NewFormLayout returns a new FormLayout instance

#### func  NewGridLayout

```go
func NewGridLayout(cols int) fyne.Layout
```
NewGridLayout returns a grid layout arranged in a specified number of columns.
The number of rows will depend on how many children are in the container that
uses this layout.

#### func  NewGridLayoutWithColumns

```go
func NewGridLayoutWithColumns(cols int) fyne.Layout
```
NewGridLayoutWithColumns returns a new grid layout that specifies a column count
and wrap to new rows when needed.

#### func  NewGridLayoutWithRows

```go
func NewGridLayoutWithRows(rows int) fyne.Layout
```
NewGridLayoutWithRows returns a new grid layout that specifies a row count that
creates new columns as required.

#### func  NewGridWrapLayout

```go
func NewGridWrapLayout(size fyne.Size) fyne.Layout
```
NewGridWrapLayout returns a new GridWrapLayout instance

#### func  NewHBoxLayout

```go
func NewHBoxLayout() fyne.Layout
```
NewHBoxLayout returns a horizontal box layout for stacking a number of child
canvas objects or widgets left to right.

#### func  NewMaxLayout

```go
func NewMaxLayout() fyne.Layout
```
NewMaxLayout creates a new MaxLayout instance

#### func  NewSpacer

```go
func NewSpacer() fyne.CanvasObject
```
NewSpacer returns a spacer object which can fill vertical and horizontal space.
This is primarily used with a box layout.

#### func  NewVBoxLayout

```go
func NewVBoxLayout() fyne.Layout
```
NewVBoxLayout returns a vertical box layout for stacking a number of child
canvas objects or widgets top to bottom.

#### type Spacer

```go
type Spacer struct {
	FixHorizontal bool
	FixVertical   bool
}
```

Spacer is any simple object that can be used in a box layout to space out child
objects

#### func (*Spacer) ExpandHorizontal

```go
func (s *Spacer) ExpandHorizontal() bool
```
ExpandHorizontal returns whether or not this spacer expands on the horizontal
axis

#### func (*Spacer) ExpandVertical

```go
func (s *Spacer) ExpandVertical() bool
```
ExpandVertical returns whether or not this spacer expands on the vertical axis

#### func (*Spacer) Hide

```go
func (s *Spacer) Hide()
```
Hide removes this Spacer from layout calculations

#### func (*Spacer) MinSize

```go
func (s *Spacer) MinSize() fyne.Size
```
MinSize returns a 0 size as a Spacer can shrink to no actual size

#### func (*Spacer) Move

```go
func (s *Spacer) Move(pos fyne.Position)
```
Move sets a new position for the Spacer - this will be called by the layout

#### func (*Spacer) Position

```go
func (s *Spacer) Position() fyne.Position
```
Position returns the current position of this Spacer

#### func (*Spacer) Refresh

```go
func (s *Spacer) Refresh()
```
Refresh does nothing for a spacer but is part of the CanvasObject definition

#### func (*Spacer) Resize

```go
func (s *Spacer) Resize(size fyne.Size)
```
Resize sets a new size for the Spacer - this will be called by the layout

#### func (*Spacer) Show

```go
func (s *Spacer) Show()
```
Show sets the Spacer to be part of the layout calculations

#### func (*Spacer) Size

```go
func (s *Spacer) Size() fyne.Size
```
Size returns the current size of this Spacer

#### func (*Spacer) Visible

```go
func (s *Spacer) Visible() bool
```
Visible returns true if this spacer should affect the layout

#### type SpacerObject

```go
type SpacerObject interface {
	ExpandVertical() bool
	ExpandHorizontal() bool
}
```

SpacerObject is any object that can be used to space out child objects
