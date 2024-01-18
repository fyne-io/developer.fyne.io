---
redirect_to:
  - https://docs.fyne.io/api/v2.3/hardwarekey

layout: page
tags: [api]
title: Fyne API "fyne.HardwareKey"
---


# fyne.HardwareKey
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type HardwareKey

```go
type HardwareKey struct {
	// ScanCode represents a hardware ID for (normally desktop) keyboard events.
	ScanCode int
}
```

HardwareKey contains information associated with physical key events Most applications should use KeyName for cross-platform compatibility.
