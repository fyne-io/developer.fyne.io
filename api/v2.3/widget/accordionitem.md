---
redirect_to:
  - https://docs.fyne.io/api/v2.3/widget/accordionitem

layout: page
tags: [api]
title: Fyne API "widget.AccordionItem"
---


# widget.AccordionItem
---
```go
import "fyne.io/fyne/v2/widget"
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

AccordionItem represents a single item in an Acc rdion.

#### func  NewAccordionItem

```go
func NewAccordionItem(title string, detail fyne.CanvasObject) *AccordionItem
```
NewAccordionItem creates a new item for an Accordion.
