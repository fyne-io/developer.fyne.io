---
redirect_to:
  - https://docs.fyne.io/api/v1.4/layout/spacerobject.md

layout: page
tags: [api]
title: Fyne API "layout.SpacerObject"
---


# layout.SpacerObject
---
```go
import "fyne.io/fyne/layout"
```

## Usage

#### type SpacerObject

```go
type SpacerObject interface {
	ExpandVertical() bool
	ExpandHorizontal() bool
}
```

SpacerObject is any object that can be used to space out child objects
