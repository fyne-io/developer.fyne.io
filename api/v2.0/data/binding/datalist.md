---
redirect_to:
  - https://docs.fyne.io/api/v2.0/data/binding/datalist.md

layout: page
tags: [api]
title: Fyne API "binding.DataList"
---


# binding.DataList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type DataList

```go
type DataList interface {
	DataItem
	GetItem(int) (DataItem, error)
	Length() int
}
```

DataList is the base interface for all bindable data lists.


<div class="since">Since: <code>
2.0</code></div>
