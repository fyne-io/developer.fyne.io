---
redirect_to:
  - https://docs.fyne.io/api/v1.2/layout/index.md

layout: page
tags: [api]
title: Fyne API layout
---


# layout
---
```go
import "fyne.io/fyne/layout"
```

Package layout defines the various layouts available to Fyne apps

## Usage

#### func  NewAdaptiveGridLayout

```go
func NewAdaptiveGridLayout(rowcols int) fyne.Layout
```
NewAdaptiveGridLayout returns a new grid layout which uses columns when horizontal but rows when vertical.

#### func  NewBorderLayout

```go
func NewBorderLayout(top, bottom, left, right fyne.CanvasObject) fyne.Layout
```
NewBorderLayout creates a new BorderLayout instance with top, left, bottom and right objects set. All other items in the container will fill the centre space

#### func  NewCenterLayout

```go
func NewCenterLayout() fyne.Layout
```
NewCenterLayout creates a new CenterLayout instance

#### func  NewFixedGridLayout

```go
func NewFixedGridLayout(size fyne.Size) fyne.Layout
```
NewFixedGridLayout returns a new FixedGridLayout instance

#### func  NewFormLayout

```go
func NewFormLayout() fyne.Layout
```
NewFormLayout returns a new FormLayout instance

#### func  NewGridLayout

```go
func NewGridLayout(cols int) fyne.Layout
```
NewGridLayout returns a grid layout arranged in a specified number of columns. The number of rows will depend on how many children are in the container that uses this layout.

#### func  NewGridLayoutWithColumns

```go
func NewGridLayoutWithColumns(cols int) fyne.Layout
```
NewGridLayoutWithColumns returns a new grid layout that specifies a column count and wrap to new rows when needed.

#### func  NewGridLayoutWithRows

```go
func NewGridLayoutWithRows(rows int) fyne.Layout
```
NewGridLayoutWithRows returns a new grid layout that specifies a row count that creates new columns as required.

#### func  NewHBoxLayout

```go
func NewHBoxLayout() fyne.Layout
```
NewHBoxLayout returns a horizontal box layout for stacking a number of child canvas objects or widgets left to right.

#### func  NewMaxLayout

```go
func NewMaxLayout() fyne.Layout
```
NewMaxLayout creates a new MaxLayout instance

#### func  NewSpacer

```go
func NewSpacer() fyne.CanvasObject
```
NewSpacer returns a spacer object which can fill vertical and horizontal space. This is primarily used with a box layout.

#### func  NewVBoxLayout

```go
func NewVBoxLayout() fyne.Layout
```
NewVBoxLayout returns a vertical box layout for stacking a number of child canvas objects or widgets top to bottom.

#### types

 * [Spacer](spacer.html)
 * [SpacerObject](spacerobject.html)
