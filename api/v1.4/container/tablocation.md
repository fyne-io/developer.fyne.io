---
layout: page
tags: [api]
title: Fyne API container
---

# container
---
```go
import "fyne.io/fyne/container"
```

## Usage

#### type TabLocation

```go
type TabLocation = widget.TabLocation
```

TabLocation is the location where the tabs of a tab container should be rendered

```go
const (
	TabLocationTop TabLocation = iota
	TabLocationLeading
	TabLocationBottom
	TabLocationTrailing
)
```
TabLocation values
