---
redirect_to:
  - https://docs.fyne.io/api/v1.3/test/softwarepainter.md

layout: page
tags: [api]
title: Fyne API test
---


# test
---
```go
import "fyne.io/fyne/test"
```

## Usage

#### type SoftwarePainter

```go
type SoftwarePainter interface {
	Paint(fyne.Canvas) image.Image
}
```

SoftwarePainter describes a simple type that can render canvases
