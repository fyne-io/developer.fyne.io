---
redirect_to:
  - https://docs.fyne.io/api/v1.4/driver/desktop/mouseevent.md

layout: page
tags: [api]
title: Fyne API "desktop.MouseEvent"
---


# desktop.MouseEvent
---
```go
import "fyne.io/fyne/driver/desktop"
```

## Usage

#### type MouseEvent

```go
type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier Modifier
}
```

MouseEvent contains data relating to desktop mouse events
