---
redirect_to:
  - https://docs.fyne.io/api/v1.3/widget/buttoniconplacement

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

#### type ButtonIconPlacement

```go
type ButtonIconPlacement int
```

ButtonIconPlacement represents the ordering of icon & text within a button.

```go
const (
	// ButtonIconLeadingText aligns the icon on the leading edge of the text.
	ButtonIconLeadingText ButtonIconPlacement = iota
	// ButtonIconTrailingText aligns the icon on the trailing edge of the text.
	ButtonIconTrailingText
)
```
