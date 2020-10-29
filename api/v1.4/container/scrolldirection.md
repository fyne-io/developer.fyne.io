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

#### type ScrollDirection

```go
type ScrollDirection = widget.ScrollDirection
```

ScrollDirection represents the directions in which a Scroll container can scroll its child content.

```go
const (
	ScrollBoth ScrollDirection = iota
	ScrollHorizontalOnly
	ScrollVerticalOnly
)
```
Constants for valid values of ScrollDirection.
