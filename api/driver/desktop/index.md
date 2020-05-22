---
layout: page
tags: [api]
title: Fyne API desktop
---

# desktop
--
    import "fyne.io/fyne/driver/desktop"


## Usage

```go
const (
	// KeyShiftLeft represents the left shift key
	KeyShiftLeft fyne.KeyName = "LeftShift"
	// KeyShiftRight represents the right shift key
	KeyShiftRight fyne.KeyName = "RightShift"
	// KeyControlLeft represents the left control key
	KeyControlLeft fyne.KeyName = "LeftControl"
	// KeyControlRight represents the right control key
	KeyControlRight fyne.KeyName = "RightControl"
	// KeyAltLeft represents the left alt key
	KeyAltLeft fyne.KeyName = "LeftAlt"
	// KeyAltRight represents the right alt key
	KeyAltRight fyne.KeyName = "RightAlt"
	// KeySuperLeft represents the left "Windows" key (or "Command" key on macOS)
	KeySuperLeft fyne.KeyName = "LeftSuper"
	// KeySuperRight represents the right "Windows" key (or "Command" key on macOS)
	KeySuperRight fyne.KeyName = "RightSuper"
	// KeyMenu represents the left or right menu / application key
	KeyMenu fyne.KeyName = "Menu"
)
```

#### type Canvas

```go
type Canvas interface {
	OnKeyDown() func(*fyne.KeyEvent)
	SetOnKeyDown(func(*fyne.KeyEvent))
	OnKeyUp() func(*fyne.KeyEvent)
	SetOnKeyUp(func(*fyne.KeyEvent))
}
```

Canvas defines the desktop specific extensions to a fyne.Canvas.

#### type Cursor

```go
type Cursor int
```

Cursor represents a standard fyne cursor

```go
const (
	// DefaultCursor is the default cursor typically an arrow
	DefaultCursor Cursor = iota
	// TextCursor is the cursor often used to indicate text selection
	TextCursor
	// CrosshairCursor is the cursor often used to indicate bitmaps
	CrosshairCursor
	// PointerCursor is the cursor often used to indicate a link
	PointerCursor
	// HResizeCursor is the cursor often used to indicate horizontal resize
	HResizeCursor
	// VResizeCursor is the cursor often used to indicate vertical resize
	VResizeCursor
)
```

#### type Cursorable

```go
type Cursorable interface {
	Cursor() Cursor
}
```

Cursorable describes any CanvasObject that needs a cursor change

#### type CustomShortcut

```go
type CustomShortcut struct {
	fyne.KeyName
	Modifier
}
```

CustomShortcut describes a shortcut desktop event.

#### func (*CustomShortcut) ShortcutName

```go
func (cs *CustomShortcut) ShortcutName() string
```
ShortcutName returns the shortcut name associated to the event

#### type Driver

```go
type Driver interface {
	// Create a new borderless window that is centered on screen
	CreateSplashWindow() fyne.Window
}
```

Driver represents the extended capabilities of a desktop driver

#### type Hoverable

```go
type Hoverable interface {
	MouseIn(*MouseEvent)
	MouseOut()
	MouseMoved(*MouseEvent)
}
```

Hoverable is used when a canvas object wishes to know if a pointer device moves
over it.

#### type Keyable

```go
type Keyable interface {
	fyne.Focusable

	KeyDown(*fyne.KeyEvent)
	KeyUp(*fyne.KeyEvent)
}
```

Keyable describes any focusable canvas object that can accept desktop key
events. This is the traditional key down and up event that is not applicable to
all devices.

#### type Modifier

```go
type Modifier int
```

Modifier captures any key modifiers (shift etc) pressed during this key event

```go
const (
	// ShiftModifier represents a shift key being held
	ShiftModifier Modifier = 1 << iota
	// ControlModifier represents the ctrl key being held
	ControlModifier
	// AltModifier represents either alt keys being held
	AltModifier
	// SuperModifier represents either super keys being held
	SuperModifier
)
```

#### type MouseButton

```go
type MouseButton int
```

MouseButton represents a single button in a desktop MouseEvent

```go
const (
	// LeftMouseButton is the most common mouse button - on some systems the only one
	LeftMouseButton MouseButton = iota + 1

	// RightMouseButton is the secondary button on most mouse input devices.
	RightMouseButton
)
```

#### type MouseEvent

```go
type MouseEvent struct {
	fyne.PointEvent
	Button   MouseButton
	Modifier Modifier
}
```

MouseEvent contains data relating to desktop mouse events

#### type Mouseable

```go
type Mouseable interface {
	MouseDown(*MouseEvent)
	MouseUp(*MouseEvent)
}
```

Mouseable represents desktop mouse events that can be sent to CanvasObjects
