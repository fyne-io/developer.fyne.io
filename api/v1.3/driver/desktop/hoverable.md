---
redirect_to:
  - https://docs.fyne.io/api/v1.3/driver/desktop/hoverable.md

layout: page
tags: [api]
title: Fyne API desktop
---


# desktop
---
```go
import "fyne.io/fyne/driver/desktop"
```

## Usage

#### type Hoverable

```go
type Hoverable interface {
	MouseIn(*MouseEvent)
	MouseOut()
	MouseMoved(*MouseEvent)
}
```

Hoverable is used when a canvas object wishes to know if a pointer device moves over it.
