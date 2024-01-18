---
redirect_to:
  - https://docs.fyne.io/api/v1.4/driver/mobile/device.md

layout: page
tags: [api]
title: Fyne API "mobile.Device"
---


# mobile.Device
---
```go
import "fyne.io/fyne/driver/mobile"
```

## Usage

#### type Device

```go
type Device interface {
	// Request that the mobile device show the touch screen keyboard (standard layout)
	ShowVirtualKeyboard()
	// Request that the mobile device show the touch screen keyboard (custom layout)
	ShowVirtualKeyboardType(KeyboardType)
	// Request that the mobile device dismiss the touch screen keyboard
	HideVirtualKeyboard()
}
```

Device describes functionality only available on mobile
