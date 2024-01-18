---
redirect_to:
  - https://docs.fyne.io/api/v2.0/doubletappable

layout: page
tags: [api]
title: Fyne API "fyne.DoubleTappable"
---


# fyne.DoubleTappable
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type DoubleTappable

```go
type DoubleTappable interface {
	DoubleTapped(*PointEvent)
}
```

DoubleTappable describes any CanvasObject that can also be double tapped.
