---
redirect_to:
  - https://docs.fyne.io/api/v2.0/driver/mobile/touchable

layout: page
tags: [api]
title: Fyne API "mobile.Touchable"
---


# mobile.Touchable
---
```go
import "fyne.io/fyne/v2/driver/mobile"
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
