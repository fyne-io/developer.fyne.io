---
redirect_to:
  - https://docs.fyne.io/api/v2.2/pointevent

layout: page
tags: [api]
title: Fyne API "fyne.PointEvent"
---


# fyne.PointEvent
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type PointEvent

```go
type PointEvent struct {
	AbsolutePosition Position // The absolute position of the event
	Position         Position // The relative position of the event
}
```

PointEvent describes a pointer input event. The position is relative to the top-left of the CanvasObject this is triggered on.
