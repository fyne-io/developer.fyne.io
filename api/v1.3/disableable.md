---
redirect_to:
  - https://docs.fyne.io/api/v1.3/disableable.md

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

#### type Disableable

```go
type Disableable interface {
	Enable()
	Disable()
	Disabled() bool
}
```

Disableable describes any CanvasObject that can be disabled. This is primarily used with objects that also implement the Tappable interface.
