---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
--
    import "fyne.io/fyne"

## Usage

#### type Settings

```go
type Settings interface {
	Theme() Theme
	SetTheme(Theme)
	Scale() float32

	AddChangeListener(chan Settings)
}
```

Settings describes the application configuration available.
