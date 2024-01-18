---
redirect_to:
  - https://docs.fyne.io/api/v2.0/shortcutcopy

layout: page
tags: [api]
title: Fyne API "fyne.ShortcutCopy"
---


# fyne.ShortcutCopy
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutCopy

```go
type ShortcutCopy struct {
	Clipboard Clipboard
}
```

ShortcutCopy describes a shortcut copy action.

#### func (*ShortcutCopy) ShortcutName

```go
func (se *ShortcutCopy) ShortcutName() string
```
ShortcutName returns the shortcut name
