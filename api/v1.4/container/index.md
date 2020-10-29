---
layout: page
tags: [api]
title: Fyne API container
---

# container
---
```go
import "fyne.io/fyne/container"
```

Package container provides container widgets that are used to lay out and organise applications

## Usage

#### func  NewAdaptiveGrid

```go
func NewAdaptiveGrid(rowcols int, objects ...fyne.CanvasObject) *fyne.Container
```
NewAdaptiveGrid creates a new container with the specified objects and using the grid layout. When in a horizontal arrangement the rowcols parameter will specify the column count, when in vertical it will specify the rows. On mobile this will dynamically refresh when device is rotated.

#### func  NewBorder

```go
func NewBorder(top, bottom, left, right fyne.CanvasObject, objects ...fyne.CanvasObject) *fyne.Container
```
NewBorder creates a new container with the specified objects and using the border layout. The top, bottom, left and right parameters specify the items that should be placed around edges, the remaining elements will be in the center. Nil can be used to an edge if it should not be filled.

#### func  NewCenter

```go
func NewCenter(objects ...fyne.CanvasObject) *fyne.Container
```
NewCenter creates a new container with the specified objects centered in the available space.

#### func  NewGridWithColumns

```go
func NewGridWithColumns(cols int, objects ...fyne.CanvasObject) *fyne.Container
```
NewGridWithColumns creates a new container with the specified objects and using the grid layout with a specified number of columns. The number of rows will depend on how many children are in the container.

#### func  NewGridWithRows

```go
func NewGridWithRows(rows int, objects ...fyne.CanvasObject) *fyne.Container
```
NewGridWithRows creates a new container with the specified objects and using the grid layout with a specified number of columns. The number of columns will depend on how many children are in the container.

#### func  NewGridWrap

```go
func NewGridWrap(size fyne.Size, objects ...fyne.CanvasObject) *fyne.Container
```
NewGridWrap creates a new container with the specified objects and using the gridwrap layout. Every element will be resized to the size parameter and the content will arrange along a row and flow to a new row if the elements don't fit.

#### func  NewHBox

```go
func NewHBox(objects ...fyne.CanvasObject) *fyne.Container
```
NewHBox creates a new container with the specified objects and using the HBox layout. The objects will be placed in the container from left to right.

#### func  NewMax

```go
func NewMax(objects ...fyne.CanvasObject) *fyne.Container
```
NewMax creates a new container with the specified objects filling the available space.

#### func  NewPadded

```go
func NewPadded(objects ...fyne.CanvasObject) *fyne.Container
```
NewPadded creates a new container with the specified objects inset by standard padding size.

#### func  NewVBox

```go
func NewVBox(objects ...fyne.CanvasObject) *fyne.Container
```
NewVBox creates a new container with the specified objects and using the VBox layout. The objects will be stacked in the container from top to bottom.

#### types

 * [AppTabs](apptabs.html)
 * [Scroll](scroll.html)
 * [ScrollDirection](scrolldirection.html)
 * [Split](split.html)
 * [TabItem](tabitem.html)
 * [TabLocation](tablocation.html)
