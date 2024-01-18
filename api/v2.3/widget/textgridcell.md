---
redirect_to:
  - https://docs.fyne.io/api/v2.3/widget/textgridcell.md

layout: page
tags: [api]
title: Fyne API "widget.TextGridCell"
---


# widget.TextGridCell
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type TextGridCell

```go
type TextGridCell struct {
	Rune  rune
	Style TextGridStyle
}
```

TextGridCell represents a single cell in a text grid. It has a rune for the text content and a style associated with it.
