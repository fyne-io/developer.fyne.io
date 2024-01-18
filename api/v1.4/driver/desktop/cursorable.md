---
redirect_to:
  - https://docs.fyne.io/api/v1.4/driver/desktop/cursorable.md

layout: page
tags: [api]
title: Fyne API "desktop.Cursorable"
---


# desktop.Cursorable
---
```go
import "fyne.io/fyne/driver/desktop"
```

## Usage

#### type Cursorable

```go
type Cursorable interface {
	Cursor() Cursor
}
```

Cursorable describes any CanvasObject that needs a cursor change
