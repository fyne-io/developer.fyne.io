---
redirect_to:
  - https://docs.fyne.io/api/v2.4/keyevent

layout: page
tags: [api]
title: Fyne API "fyne.KeyEvent"
package: fyne.io/fyne/v2
---
# fyne.KeyEvent
---

```go
import "fyne.io/fyne/v2"
```

## Usage

#### type KeyEvent

```go
type KeyEvent struct {
	// Name describes the keyboard event that is consistent across platforms.
	Name KeyName
	// Physical is a platform specific field that reports the hardware information of physical keyboard events.
	Physical HardwareKey
}
```

KeyEvent describes a keyboard input event.
