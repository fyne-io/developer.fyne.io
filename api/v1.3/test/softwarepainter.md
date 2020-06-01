---
layout: page
tags: [api]
title: Fyne API test
---

# test
--
    import "fyne.io/fyne/test"

## Usage

#### type SoftwarePainter

```go
type SoftwarePainter interface {
	Paint(fyne.Canvas) image.Image
}
```

SoftwarePainter describes a simple type that can render canvases
