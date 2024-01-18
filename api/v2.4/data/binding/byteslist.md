---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "binding.BytesList"
package: fyne.io/fyne/v2/data/binding
---
# binding.BytesList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type BytesList

```go
type BytesList interface {
	DataList

	Append(value []byte) error
	Get() ([][]byte, error)
	GetValue(index int) ([]byte, error)
	Prepend(value []byte) error
	Set(list [][]byte) error
	SetValue(index int, value []byte) error
}
```

BytesList supports binding a list of []byte values.


<div class="since">Since: <code>
2.2</code></div>

#### func  NewBytesList

```go
func NewBytesList() BytesList
```
NewBytesList returns a bindable list of []byte values.


<div class="since">Since: <code>
2.2</code></div>
