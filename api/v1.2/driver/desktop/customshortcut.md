---
redirect_to:
  - https://docs.fyne.io/api/v1.2/driver/desktop/customshortcut.md

layout: page
tags: [api]
title: Fyne API desktop
---


# desktop
---
```go
import "fyne.io/fyne/driver/desktop"
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
