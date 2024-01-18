---
redirect_to:
  - https://docs.fyne.io/api/v1.2/shortcutselectall

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

#### type ShortcutSelectAll

```go
type ShortcutSelectAll struct{}
```

ShortcutSelectAll describes a shortcut selectAll action.

#### func (*ShortcutSelectAll) ShortcutName

```go
func (se *ShortcutSelectAll) ShortcutName() string
```
ShortcutName returns the shortcut name
