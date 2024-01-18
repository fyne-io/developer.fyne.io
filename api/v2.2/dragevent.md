---
redirect_to:
  - https://docs.fyne.io/api/v2.2/dragevent

layout: page
tags: [api]
title: Fyne API "fyne.DragEvent"
---


# fyne.DragEvent
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type DragEvent

```go
type DragEvent struct {
	PointEvent
	Dragged Delta
}
```

DragEvent defines the parameters of a pointer or other drag event. The DraggedX and DraggedY fields show how far the item was dragged since the last event.
