---
layout: page
tags: [api]
title: Fyne API mobile
---

# mobile
--
    import "fyne.io/fyne/driver/mobile"


## Usage

#### type Device

```go
type Device interface {
	ShowVirtualKeyboard() // Request that the mobile device show the touch screen keyboard (standard layout)
	HideVirtualKeyboard() // Request that the mobile device dismiss the touch screen keyboard
}
```

Device describes functionality only available on mobile

#### type TouchEvent

```go
type TouchEvent struct {
	fyne.PointEvent
}
```

TouchEvent contains data relating to mobile touch events

#### type Touchable

```go
type Touchable interface {
	TouchDown(*TouchEvent)
	TouchUp(*TouchEvent)
	TouchCancel(*TouchEvent)
}
```

Touchable represents mobile touch events that can be sent to CanvasObjects
