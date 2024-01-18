---
redirect_to:
  - https://docs.fyne.io/api/v1.2/driver/mobile/device.md

layout: page
tags: [api]
title: Fyne API mobile
---


# mobile
---
```go
import "fyne.io/fyne/driver/mobile"
```

## Usage

#### type Device

```go
type Device interface {
	ShowVirtualKeyboard() // Request that the mobile device show the touch screen keyboard (standard layout)
	HideVirtualKeyboard() // Request that the mobile device dismiss the touch screen keyboard
}
```

Device describes functionality only available on mobile
