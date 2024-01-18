---
redirect_to:
  - https://docs.fyne.io/api/v2.1/animationcurve

layout: page
tags: [api]
title: Fyne API "fyne.AnimationCurve"
---


# fyne.AnimationCurve
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type AnimationCurve

```go
type AnimationCurve func(float32) float32
```

AnimationCurve represents an animation algorithm for calculating the progress through a timeline. Custom animations can be provided by implementing the "func(float32) float32" definition. The input parameter will start at 0.0 when an animation starts and travel up to 1.0 at which point it will end. A linear animation would return the same output value as is passed in.
