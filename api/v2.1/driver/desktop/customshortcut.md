---
redirect_to:
  - https://docs.fyne.io/api/v2.1/driver/desktop/customshortcut.md

layout: page
tags: [api]
title: Fyne API "desktop.CustomShortcut"
---


# desktop.CustomShortcut
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type CustomShortcut

```go
type CustomShortcut struct {
	fyne.KeyName
	Modifier
}
```

CustomShortcut describes a shortcut desktop event.

#### func (*CustomShortcut) ShortcutName

```go
func (cs *CustomShortcut) ShortcutName() string
```
ShortcutName returns the shortcut name associated to the event
