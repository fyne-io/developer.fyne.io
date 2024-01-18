---
redirect_to:
  - https://docs.fyne.io/api/v1.2/menuitem

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

#### type MenuItem

```go
type MenuItem struct {
	Label  string
	Action func()
}
```

MenuItem is a sligne item within any menu, it contains a dispay Label and Action function that is called when tapped.

#### func  NewMenuItem

```go
func NewMenuItem(label string, action func()) *MenuItem
```
NewMenuItem creates a new menu item from the passed label and action parameters.
