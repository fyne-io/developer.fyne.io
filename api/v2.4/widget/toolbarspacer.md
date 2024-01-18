---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "widget.ToolbarSpacer"
package: fyne.io/fyne/v2/widget
---
# widget.ToolbarSpacer
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ToolbarSpacer

```go
type ToolbarSpacer struct {
}
```

ToolbarSpacer is a blank, stretchable space for a toolbar. This is typically used to assist layout if you wish some left and some right aligned items. Space will be split evebly amongst all the spacers on a toolbar.

#### func  NewToolbarSpacer

```go
func NewToolbarSpacer() *ToolbarSpacer
```
NewToolbarSpacer returns a new spacer item for a Toolbar to assist with ToolbarItem alignment

#### func (*ToolbarSpacer) ToolbarObject

```go
func (t *ToolbarSpacer) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the actual spacer object for this ToolbarSpacer
