---
redirect_to:
  - https://docs.fyne.io/api/v1.4/theme

layout: page
tags: [api]
title: Fyne API "fyne.Theme"
---


# fyne.Theme
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Theme

```go
type Theme interface {
	BackgroundColor() color.Color
	ButtonColor() color.Color
	DisabledButtonColor() color.Color
	// Deprecated: Hyperlinks now use the primary color for consistency.
	HyperlinkColor() color.Color
	TextColor() color.Color
	DisabledTextColor() color.Color
	// Deprecated: Icons now use the text colour for consistency.
	IconColor() color.Color
	// Deprecated: Disabled icons match disabled text color for consistency.
	DisabledIconColor() color.Color
	PlaceHolderColor() color.Color
	PrimaryColor() color.Color
	HoverColor() color.Color
	FocusColor() color.Color
	ScrollBarColor() color.Color
	ShadowColor() color.Color

	TextSize() int
	TextFont() Resource
	TextBoldFont() Resource
	TextItalicFont() Resource
	TextBoldItalicFont() Resource
	TextMonospaceFont() Resource

	Padding() int
	IconInlineSize() int
	ScrollBarSize() int
	ScrollBarSmallSize() int
}
```

Theme defines the requirements of any Fyne theme.
