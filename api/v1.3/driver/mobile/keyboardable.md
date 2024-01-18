---
redirect_to:
  - https://docs.fyne.io/api/v1.3/driver/mobile/keyboardable

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

#### type Keyboardable

```go
type Keyboardable interface {
	fyne.Focusable

	Keyboard() KeyboardType
}
```

Keyboardable describes any CanvasObject that needs a keyboard
