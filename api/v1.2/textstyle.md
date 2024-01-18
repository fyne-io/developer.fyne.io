---
redirect_to:
  - https://docs.fyne.io/api/v1.2/textstyle.md

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

#### type TextStyle

```go
type TextStyle struct {
	Bold      bool // Should text be bold
	Italic    bool // Should text be italic
	Monospace bool // Use the system monospace font instead of regular
}
```

TextStyle represents the styles that can be applied to a text canvas object or text based widget.
