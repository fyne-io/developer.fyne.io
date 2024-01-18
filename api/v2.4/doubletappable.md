---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "fyne.DoubleTappable"
package: fyne.io/fyne/v2
---
# fyne.DoubleTappable
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type DoubleTappable

```go
type DoubleTappable interface {
	DoubleTapped(*PointEvent)
}
```

DoubleTappable describes any CanvasObject that can also be double tapped.
