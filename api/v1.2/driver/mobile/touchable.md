---
redirect_to:
  - https://docs.fyne.io/api/v1.2/driver/mobile/touchable.md

layout: page
tags: [api]
title: Fyne API mobile
---


# mobile
---
```go
import "fyne.io/fyne/driver/mobile"
```

## Usage

#### type Touchable

```go
type Touchable interface {
	TouchDown(*TouchEvent)
	TouchUp(*TouchEvent)
	TouchCancel(*TouchEvent)
}
```

Touchable represents mobile touch events that can be sent to CanvasObjects
