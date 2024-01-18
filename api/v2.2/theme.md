---
redirect_to:
  - https://docs.fyne.io/api/v2.2/theme

layout: page
tags: [api]
title: Fyne API "fyne.Theme"
---


# fyne.Theme
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Theme

```go
type Theme interface {
	Color(ThemeColorName, ThemeVariant) color.Color
	Font(TextStyle) Resource
	Icon(ThemeIconName) Resource
	Size(ThemeSizeName) float32
}
```

Theme defines the method to look up colors, sizes and fonts that make up a Fyne theme.


<div class="since">Since: <code>
2.0</code></div>
