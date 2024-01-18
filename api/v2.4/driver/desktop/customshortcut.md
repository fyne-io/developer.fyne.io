---
redirect_to:
  - https://docs.fyne.io/api/v2.4/driver/desktop/customshortcut

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
	Modifier fyne.KeyModifier
}
```

CustomShortcut describes a shortcut desktop event.

#### func (*CustomShortcut) Key

```go
func (cs *CustomShortcut) Key() fyne.KeyName
```
Key returns the key name of this shortcut. @implements KeyboardShortcut

#### func (*CustomShortcut) Mod

```go
func (cs *CustomShortcut) Mod() fyne.KeyModifier
```
Mod returns the modifier of this shortcut. @implements KeyboardShortcut

#### func (*CustomShortcut) ShortcutName

```go
func (cs *CustomShortcut) ShortcutName() string
```
ShortcutName returns the shortcut name associated to the event
