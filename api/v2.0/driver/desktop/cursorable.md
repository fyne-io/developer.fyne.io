---
redirect_to:
  - https://docs.fyne.io/api/v2.0/driver/desktop/cursorable

layout: page
tags: [api]
title: Fyne API "desktop.Cursorable"
---


# desktop.Cursorable
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type Cursorable

```go
type Cursorable interface {
	Cursor() Cursor
}
```

Cursorable describes any CanvasObject that needs a cursor change
