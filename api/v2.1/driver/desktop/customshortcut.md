---
layout: page
tags: [api]
title: Fyne API "desktop.CustomShortcut"
package: fyne.io/fyne/v2/driver/desktop
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
