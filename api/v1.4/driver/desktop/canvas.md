---
redirect_to:
  - https://docs.fyne.io/api/v1.4/driver/desktop/canvas

layout: page
tags: [api]
title: Fyne API "desktop.Canvas"
---


# desktop.Canvas
---
```go
import "fyne.io/fyne/driver/desktop"
```

## Usage

#### type Canvas

```go
type Canvas interface {
	OnKeyDown() func(*fyne.KeyEvent)
	SetOnKeyDown(func(*fyne.KeyEvent))
	OnKeyUp() func(*fyne.KeyEvent)
	SetOnKeyUp(func(*fyne.KeyEvent))
}
```

Canvas defines the desktop specific extensions to a fyne.Canvas.
