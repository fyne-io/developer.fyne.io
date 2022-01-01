---
layout: page
tags: [api]
title: Fyne API "desktop.MouseEvent"
package: fyne.io/fyne/v2/driver/desktop
---

# desktop.MouseEvent
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type MouseEvent

```go
type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier Modifier
}
```

MouseEvent contains data relating to desktop mouse events
