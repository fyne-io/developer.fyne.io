---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutCopy"
package: fyne.io/fyne/v2
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
