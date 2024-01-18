---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "binding.ExternalFloatList"
package: fyne.io/fyne/v2/data/binding
---
# binding.ExternalFloatList
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type ExternalFloatList

```go
type ExternalFloatList interface {
	FloatList

	Reload() error
}
```

ExternalFloatList supports binding a list of float64 values from an external variable.


<div class="since">Since: <code>
2.0</code></div>

#### func  BindFloatList

```go
func BindFloatList(v *[]float64) ExternalFloatList
```
BindFloatList returns a bound list of float64 values, based on the contents of the passed slice. If your code changes the content of the slice this refers to you should call Reload() to inform the bindings.


<div class="since">Since: <code>
2.0</code></div>
