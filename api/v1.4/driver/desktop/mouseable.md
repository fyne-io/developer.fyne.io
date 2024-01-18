---
redirect_to:
  - https://docs.fyne.io/api/v1.4/driver/desktop/mouseable.md

layout: page
tags: [api]
title: Fyne API "desktop.Mouseable"
---


# desktop.Mouseable
---
```go
import "fyne.io/fyne/driver/desktop"
```

## Usage

#### type Mouseable

```go
type Mouseable interface {
	MouseDown(*MouseEvent)
	MouseUp(*MouseEvent)
}
```

Mouseable represents desktop mouse events that can be sent to CanvasObjects
