---
redirect_to:
  - https://docs.fyne.io/api/v2.4/driver/desktop/mouseable

layout: page
tags: [api]
title: Fyne API "desktop.Mouseable"
package: fyne.io/fyne/v2/driver/desktop
---
# desktop.Mouseable
---

```go
import "fyne.io/fyne/v2/driver/desktop"
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
