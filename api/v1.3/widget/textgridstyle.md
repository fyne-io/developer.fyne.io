---
redirect_to:
  - https://docs.fyne.io/api/v1.3/widget/textgridstyle

layout: page
tags: [api]
title: Fyne API widget
---


# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type TextGridStyle

```go
type TextGridStyle interface {
	TextColor() color.Color
	BackgroundColor() color.Color
}
```

TextGridStyle defines a style that can be applied to a TextGrid cell.

```go
var (
	// TextGridStyleDefault is a default style for test grid cells
	TextGridStyleDefault TextGridStyle
	// TextGridStyleWhitespace is the style used for whitespace characters, if enabled
	TextGridStyleWhitespace TextGridStyle
)
```
