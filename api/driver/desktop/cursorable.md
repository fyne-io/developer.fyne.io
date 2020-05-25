---
layout: page
tags: [api]
title: Fyne API desktop
---

# desktop
--
    import "fyne.io/fyne/driver/desktop"

## Usage

#### type Cursorable

```go
type Cursorable interface {
	Cursor() Cursor
}
```

Cursorable describes any CanvasObject that needs a cursor change
