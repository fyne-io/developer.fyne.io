---
redirect_to:
  - https://docs.fyne.io/api/v1.3/shortcutable.md

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

#### type Shortcutable

```go
type Shortcutable interface {
	TypedShortcut(Shortcut)
}
```

Shortcutable describes any CanvasObject that can respond to shortcut commands (quit, cut, copy, and paste).
