---
redirect_to:
  - https://docs.fyne.io/api/v1.3/shortcut.md

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

#### type Shortcut

```go
type Shortcut interface {
	ShortcutName() string
}
```

Shortcut is the interface used to describe a shortcut action
