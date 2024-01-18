---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "fyne.Tappable"
package: fyne.io/fyne/v2
---
# fyne.Tappable
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Tappable

```go
type Tappable interface {
	Tapped(*PointEvent)
}
```

Tappable describes any CanvasObject that can also be tapped. This should be implemented by buttons etc that wish to handle pointer interactions.
