---
redirect_to:
  - https://docs.fyne.io/api/v1.4/shortcutcut

layout: page
tags: [api]
title: Fyne API "fyne.ShortcutCut"
---


# fyne.ShortcutCut
---
```go
import "fyne.io/fyne"
```

## Usage

#### type ShortcutCut

```go
type ShortcutCut struct {
	Clipboard Clipboard
}
```

ShortcutCut describes a shortcut cut action.

#### func (*ShortcutCut) ShortcutName

```go
func (se *ShortcutCut) ShortcutName() string
```
ShortcutName returns the shortcut name
