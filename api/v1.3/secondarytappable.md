---
redirect_to:
  - https://docs.fyne.io/api/v1.3/secondarytappable.md

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

#### type SecondaryTappable

```go
type SecondaryTappable interface {
	TappedSecondary(*PointEvent)
}
```

SecondaryTappable describes a CanvasObject that can be right-clicked or long-tapped.
