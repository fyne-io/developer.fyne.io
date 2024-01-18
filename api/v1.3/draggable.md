---
redirect_to:
  - https://docs.fyne.io/api/v1.3/draggable.md

layout: page
tags: [api]
title: Fyne API fyne
---


# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Draggable

```go
type Draggable interface {
	Dragged(*DragEvent)
	DragEnd()
}
```

Draggable indicates that a CanvasObject can be dragged. This is used for any item that the user has indicated should be moved across the screen.
