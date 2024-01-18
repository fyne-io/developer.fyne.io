---
redirect_to:
  - https://docs.fyne.io/api/v2.3/widget/richtextstyle.md

layout: page
tags: [api]
title: Fyne API "widget.RichTextStyle"
---


# widget.RichTextStyle
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type RichTextStyle

```go
type RichTextStyle struct {
	Alignment fyne.TextAlign
	ColorName fyne.ThemeColorName
	Inline    bool
	SizeName  fyne.ThemeSizeName
	TextStyle fyne.TextStyle
}
```

RichTextStyle describes the details of a text object inside a RichText widget.


<div class="since">Since: <code>
2.1</code></div>
