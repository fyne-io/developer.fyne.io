---
redirect_to:
  - https://docs.fyne.io/api/v2.0/widget/buttonalign

layout: page
tags: [api]
title: Fyne API "widget.ButtonAlign"
---


# widget.ButtonAlign
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ButtonAlign

```go
type ButtonAlign int
```

ButtonAlign represents the horizontal alignment of a button.

```go
const (
	// ButtonAlignCenter aligns the icon and the text centrally.
	ButtonAlignCenter ButtonAlign = iota
	// ButtonAlignLeading aligns the icon and the text with the leading edge.
	ButtonAlignLeading
	// ButtonAlignTrailing aligns the icon and the text with the trailing edge.
	ButtonAlignTrailing
)
```
