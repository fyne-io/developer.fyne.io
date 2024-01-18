---
redirect_to:
  - https://docs.fyne.io/api/v2.1/menu.md

layout: page
tags: [api]
title: Fyne API "fyne.Menu"
---


# fyne.Menu
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Menu

```go
type Menu struct {
	Label string
	Items []*MenuItem
}
```

Menu stores the information required for a standard menu. A menu can pop down from a MainMenu or could be a pop out menu.

#### func  NewMenu

```go
func NewMenu(label string, items ...*MenuItem) *Menu
```
NewMenu creates a new menu given the specified label (to show in a MainMenu) and list of items to display.
