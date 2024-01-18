---
redirect_to:
  - https://docs.fyne.io/api/v1.3/driver/desktop/driver

layout: page
tags: [api]
title: Fyne API desktop
---


# desktop
---
```go
import "fyne.io/fyne/driver/desktop"
```

## Usage

#### type Driver

```go
type Driver interface {
	// Create a new borderless window that is centered on screen
	CreateSplashWindow() fyne.Window
}
```

Driver represents the extended capabilities of a desktop driver
