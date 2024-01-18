---
redirect_to:
  - https://docs.fyne.io/api/v1.4/widget/accordionitem.md

layout: page
tags: [api]
title: Fyne API "widget.AccordionItem"
---


# widget.AccordionItem
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type AccordionItem

```go
type AccordionItem struct {
	Title  string
	Detail fyne.CanvasObject
	Open   bool
}
```

AccordionItem represents a single item in an Accordion.

#### func  NewAccordionItem

```go
func NewAccordionItem(title string, detail fyne.CanvasObject) *AccordionItem
```
NewAccordionItem creates a new item for an Accordion.
