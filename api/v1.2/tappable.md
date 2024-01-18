---
redirect_to:
  - https://docs.fyne.io/api/v1.2/tappable.md

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

#### type Tappable

```go
type Tappable interface {
	Tapped(*PointEvent)
	TappedSecondary(*PointEvent)
}
```

Tappable describes any CanvasObject that can also be tapped. This should be implemented by buttons etc that wish to handle pointer interactions.
