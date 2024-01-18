---
redirect_to:
  - https://docs.fyne.io/api/v1.2/tools/playground/

layout: page
tags: [api]
title: Fyne API playground
---


# playground
---
```go
import "fyne.io/fyne/tools/playground"
```


## Usage

#### func  NewSoftwareCanvas

```go
func NewSoftwareCanvas() test.WindowlessCanvas
```
NewSoftwareCanvas creates a new canvas in memory that can render without hardware support

#### func  Render

```go
func Render(obj fyne.CanvasObject)
```
Render takes a canvasobject and converts it into an inline image for showing in the playground

#### func  RenderCanvas

```go
func RenderCanvas(c fyne.Canvas)
```
RenderCanvas takes a canvas and converts it into an inline image for showing in the playground

#### func  RenderWindow

```go
func RenderWindow(w fyne.Window)
```
RenderWindow takes a window and converts it's canvas into an inline image for showing in the playground

#### types
