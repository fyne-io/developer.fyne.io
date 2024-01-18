---
redirect_to:
  - https://docs.fyne.io/api/v2.4/settings

layout: page
tags: [api]
title: Fyne API "fyne.Settings"
package: fyne.io/fyne/v2
---
# fyne.Settings
---

```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Settings

```go
type Settings interface {
	Theme() Theme
	SetTheme(Theme)
	// ThemeVariant defines which preferred version of a theme should be used (i.e. light or dark)
	//
	// Since: 2.0
	ThemeVariant() ThemeVariant
	Scale() float32
	// PrimaryColor indicates a user preference for a named primary color
	//
	// Since: 1.4
	PrimaryColor() string

	AddChangeListener(chan Settings)
	BuildType() BuildType

	ShowAnimations() bool
}
```

Settings describes the application configuration available.
