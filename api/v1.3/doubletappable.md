---
redirect_to:
  - https://docs.fyne.io/api/v1.3/doubletappable

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

#### type DoubleTappable

```go
type DoubleTappable interface {
	DoubleTapped(*PointEvent)
}
```

DoubleTappable describes any CanvasObject that can also be double tapped.
