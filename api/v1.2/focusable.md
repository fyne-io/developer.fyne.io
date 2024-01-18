---
redirect_to:
  - https://docs.fyne.io/api/v1.2/focusable.md

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

#### type Focusable

```go
type Focusable interface {
	FocusGained()
	FocusLost()
	Focused() bool

	TypedRune(rune)
	TypedKey(*KeyEvent)
}
```

Focusable describes any CanvasObject that can respond to being focused. It will receive the FocusGained and FocusLost events appropriately. When focused it will also have TypedRune called as text is input and TypedKey called when other keys are pressed.
