---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
--
    import "fyne.io/fyne"

## Usage

#### type ShortcutPaste

```go
type ShortcutPaste struct {
	Clipboard Clipboard
}
```

ShortcutPaste describes a shortcut paste action.

#### func (*ShortcutPaste) ShortcutName

```go
func (se *ShortcutPaste) ShortcutName() string
```
ShortcutName returns the shortcut name
