---
redirect_to:
  - https://docs.fyne.io/api/v2.1/driver/desktop/modifier

layout: page
tags: [api]
title: Fyne API "desktop.Modifier"
---


# desktop.Modifier
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Modifier

```go
type Modifier int
```

Modifier captures any key modifiers (shift etc) pressed during this key event

```go
const (
	// ShiftModifier represents a shift key being held
	ShiftModifier Modifier = 1 << iota
	// ControlModifier represents the ctrl key being held
	ControlModifier
	// AltModifier represents either alt keys being held
	AltModifier
	// SuperModifier represents either super keys being held
	SuperModifier
)
```
