---
redirect_to:
  - https://docs.fyne.io/api/v2.2/widget/disableablewidget.md

layout: page
tags: [api]
title: Fyne API "widget.DisableableWidget"
---


# widget.DisableableWidget
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type DisableableWidget

```go
type DisableableWidget struct {
	BaseWidget
}
```

DisableableWidget describes an extension to BaseWidget which can be disabled. Disabled widgets should have a visually distinct style when disabled, normally using theme.DisabledButtonColor.

#### func (*DisableableWidget) Disable

```go
func (w *DisableableWidget) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style appropriately.

#### func (*DisableableWidget) Disabled

```go
func (w *DisableableWidget) Disabled() bool
```
Disabled returns true if this widget is currently disabled or false if it can currently be interacted with.

#### func (*DisableableWidget) Enable

```go
func (w *DisableableWidget) Enable()
```
Enable this widget, updating any style or features appropriately.
