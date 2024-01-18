---
redirect_to:
  - https://docs.fyne.io/api/v2.0/data/binding/boollist

layout: page
tags: [api]
title: Fyne API "binding.BoolList"
---


# binding.BoolList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type BoolList

```go
type BoolList interface {
	DataList

	Append(bool) error
	Get() ([]bool, error)
	GetValue(int) (bool, error)
	Prepend(bool) error
	Set([]bool) error
	SetValue(int, bool) error
}
```

BoolList supports binding a list of bool values.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewBoolList

```go
func NewBoolList() BoolList
```
NewBoolList returns a bindable list of bool values.


<div class="since">Since: <code>
2.0</code></div>
