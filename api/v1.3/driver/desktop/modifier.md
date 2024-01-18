---
redirect_to:
  - https://docs.fyne.io/api/v1.3/driver/desktop/modifier.md

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
