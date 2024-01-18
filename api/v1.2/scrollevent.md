---
redirect_to:
  - https://docs.fyne.io/api/v1.2/scrollevent.md

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

#### type ScrollEvent

```go
type ScrollEvent struct {
	PointEvent
	DeltaX, DeltaY int
}
```

ScrollEvent defines the parameters of a pointer or other scroll event. The DeltaX and DeltaY represent how large the scroll was in two dimensions.
