---
redirect_to:
  - https://docs.fyne.io/api/v2.4/driver/androidcontext

layout: page
tags: [api]
title: Fyne API "driver.AndroidContext"
package: fyne.io/fyne/v2/driver
---
# driver.AndroidContext
---

```go
import "fyne.io/fyne/v2/driver"
```

## Usage

#### type AndroidContext

```go
type AndroidContext struct {
	VM, Env, Ctx uintptr
}
```

AndroidContext is passed to the `RunNative` callback when it is executed on an Android device. The VM, Env and Ctx pointers are reqiured to make various calls into JVM methods.


<div class="since">Since: <code>
2.3</code></div>
