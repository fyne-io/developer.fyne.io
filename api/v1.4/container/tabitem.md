---
layout: page
tags: [api]
title: Fyne API container
---

# container
---
```go
import "fyne.io/fyne/container"
```

## Usage

#### type TabItem

```go
type TabItem = widget.TabItem
```

TabItem represents a single view in a TabContainer. The Text and Icon are used for the tab button and the Content is shown when the corresponding tab is active.

#### func  NewTabItem

```go
func NewTabItem(text string, content fyne.CanvasObject) *TabItem
```
NewTabItem creates a new item for a tabbed widget - each item specifies the content and a label for its tab.

#### func  NewTabItemWithIcon

```go
func NewTabItemWithIcon(text string, icon fyne.Resource, content fyne.CanvasObject) *TabItem
```
NewTabItemWithIcon creates a new item for a tabbed widget - each item specifies the content and a label with an icon for its tab.
