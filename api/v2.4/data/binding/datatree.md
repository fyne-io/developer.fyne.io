---
redirect_to:
  - https://docs.fyne.io/api/v2.4/data/binding/datatree

layout: page
tags: [api]
title: Fyne API "binding.DataTree"
package: fyne.io/fyne/v2/data/binding
---
# binding.DataTree
---

```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type DataTree

```go
type DataTree interface {
	DataItem
	GetItem(id string) (DataItem, error)
	ChildIDs(string) []string
}
```

DataTree is the base interface for all bindable data trees.


<div class="since">Since: <code>
2.4</code></div>
