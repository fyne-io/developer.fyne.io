---
redirect_to:
  - https://docs.fyne.io/api/v1.4/secondarytappable

layout: page
tags: [api]
title: Fyne API "fyne.SecondaryTappable"
---


# fyne.SecondaryTappable
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
