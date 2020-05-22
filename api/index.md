---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
--
    import "fyne.io/fyne"

Package fyne describes the objects and components available to any Fyne app.
These can all be created, manipulated and tested without rendering (for speed).
Your main package should use the app package to create an application with a
default driver that will render your UI.

A simple application may look like this:

    package main

    import "fyne.io/fyne/app"
    import "fyne.io/fyne/widget"

    func main() {
    	a := app.New()

    	w := a.NewWindow("Hello")
    	w.SetContent(widget.NewVBox(
    		widget.NewLabel("Hello Fyne!"),
    		widget.NewButton("Quit", func() {
    			a.Quit()
    		})))

    	w.ShowAndRun()
    }

## Usage

```go
const SettingsScaleAuto = float32(-1.0)
```
SettingsScaleAuto is a specific scale value that indicates a canvas should scale
according to the DPI of the window that contains it.

#### func  IsHorizontal

```go
func IsHorizontal(orient DeviceOrientation) bool
```
IsHorizontal is a helper utility that determines if a passed orientation is
horizontal

#### func  IsVertical

```go
func IsVertical(orient DeviceOrientation) bool
```
IsVertical is a helper utility that determines if a passed orientation is
vertical

#### func  LogError

```go
func LogError(reason string, err error)
```
LogError reports an error to the command line with the specified err cause, if
not nil. The function also reports basic information about the code location.

#### func  Max

```go
func Max(x, y int) int
```
Max returns the larger of the passed values.

#### func  Min

```go
func Min(x, y int) int
```
Min returns the smaller of the passed values.

#### func  SetCurrentApp

```go
func SetCurrentApp(current App)
```
SetCurrentApp is an internal function to set the app instance currently running.

#### type App

```go
type App interface {
	// Create a new window for the application.
	// The first window to open is considered the "master" and when closed
	// the application will exit.
	NewWindow(title string) Window

	// Open a URL in the default browser application.
	OpenURL(url *url.URL) error

	// Icon returns the application icon, this is used in various ways
	// depending on operating system.
	// This is also the default icon for new windows.
	Icon() Resource

	// SetIcon sets the icon resource used for this application instance.
	SetIcon(Resource)

	// Run the application - this starts the event loop and waits until Quit()
	// is called or the last window closes.
	// This should be called near the end of a main() function as it will block.
	Run()

	// Calling Quit on the application will cause the application to exit
	// cleanly, closing all open windows.
	// This function does no thing on a mobile device as the application lifecycle is
	// managed by the operating system.
	Quit()

	// Driver returns the driver that is rendering this application.
	// Typically not needed for day to day work, mostly internal functionality.
	Driver() Driver

	// UniqueID returns the application unique identifier, if set.
	// This must be set for use of the Preferences() functions... see NewWithId(string)
	UniqueID() string

	// SendNotification sends a system notification that will be displayed in the operating system's notification area.
	SendNotification(*Notification)

	// Settings return the globally set settings, determining theme and so on.
	Settings() Settings

	// Preferences returns the application preferences, used for storing configuration and state
	Preferences() Preferences
}
```

An App is the definition of a graphical application. Apps can have multiple
windows, it will exit when the first window to be shown is closed. You can also
cause the app to exit by calling Quit(). To start an application you need to
call Run() somewhere in your main() function. Alternatively use the
window.ShowAndRun() function for your main window.

#### func  CurrentApp

```go
func CurrentApp() App
```
CurrentApp returns the current application, for which there is only 1 per
process.

#### type Canvas

```go
type Canvas interface {
	Content() CanvasObject
	SetContent(CanvasObject)
	Refresh(CanvasObject)
	Focus(Focusable)
	Unfocus()
	Focused() Focusable

	// Size returns the current size of this canvas
	Size() Size
	// Scale returns the current scale (multiplication factor) this canvas uses to render
	// The pixel size of a CanvasObject can be found by multiplying by this value.
	Scale() float32
	// SetScale sets ths scale for this canvas only, overriding system and user settings.
	//
	// Deprecated: Settings are now calculated solely on the user configuration and system setup.
	SetScale(float32)

	// Overlay returns the current overlay.
	//
	// Deprecated: Overlays are stacked now.
	// This method returns the top of the overlay stack.
	// Use Overlays() instead.
	Overlay() CanvasObject
	// Overlays returns the overlay stack.
	Overlays() OverlayStack
	// SetOverlay sets the overlay for the canvas.
	//
	// Deprecated: Overlays are stacked now.
	// This method replaces the whole stack by the given overlay.
	// Use Overlays() instead.
	SetOverlay(CanvasObject)

	OnTypedRune() func(rune)
	SetOnTypedRune(func(rune))
	OnTypedKey() func(*KeyEvent)
	SetOnTypedKey(func(*KeyEvent))
	AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut))

	Capture() image.Image

	// PixelCoordinateForPosition returns the x and y pixel coordinate for a given position on this canvas.
	// This can be used to find absolute pixel positions or pixel offsets relative to an object top left.
	PixelCoordinateForPosition(Position) (int, int)
}
```

Canvas defines a graphical canvas to which a CanvasObject or Container can be
added. Each canvas has a scale which is automatically applied during the render
process.

#### type CanvasObject

```go
type CanvasObject interface {
	// geometry
	Size() Size
	Resize(Size)
	Position() Position
	Move(Position)
	MinSize() Size

	// visibility
	Visible() bool
	// Show shows this object.
	Show()
	// Hide hides this object.
	Hide()

	Refresh()
}
```

CanvasObject describes any graphical object that can be added to a canvas.
Objects have a size and position that can be controlled through this API.
MinSize is used to determine the minimum size which this object should be
displayed. An object will be visible by default but can be hidden with Hide()
and re-shown with Show().

Note: If this object is controlled as part of a Layout you should not call
Resize(Size) or Move(Position).

#### type Clipboard

```go
type Clipboard interface {
	// Content returns the clipboard content
	Content() string
	// SetContent sets the clipboard content
	SetContent(content string)
}
```

Clipboard represents the system clipboard interface

#### type Container

```go
type Container struct {
	Hidden bool // Is this Container hidden

	Layout  Layout         // The Layout algorithm for arranging child CanvasObjects
	Objects []CanvasObject // The set of CanvasObjects this container holds
}
```

Container is a CanvasObject that contains a collection of child objects. The
layout of the children is set by the specified Layout.

#### func  NewContainer

```go
func NewContainer(objects ...CanvasObject) *Container
```
NewContainer returns a new Container instance holding the specified
CanvasObjects.

#### func  NewContainerWithLayout

```go
func NewContainerWithLayout(layout Layout, objects ...CanvasObject) *Container
```
NewContainerWithLayout returns a new Container instance holding the specified
CanvasObjects which will be laid out according to the specified Layout.

#### func (*Container) AddObject

```go
func (c *Container) AddObject(o CanvasObject)
```
AddObject adds another CanvasObject to the set this Container holds.

#### func (*Container) Hide

```go
func (c *Container) Hide()
```
Hide sets this container, and all its children, to be not visible.

#### func (*Container) MinSize

```go
func (c *Container) MinSize() Size
```
MinSize calculates the minimum size of a Container. This is delegated to the
Layout, if specified, otherwise it will mimic MaxLayout.

#### func (*Container) Move

```go
func (c *Container) Move(pos Position)
```
Move the container (and all its children) to a new position, relative to its
parent.

#### func (*Container) Position

```go
func (c *Container) Position() Position
```
Position gets the current position of this Container, relative to its parent.

#### func (*Container) Refresh

```go
func (c *Container) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Container) Resize

```go
func (c *Container) Resize(size Size)
```
Resize sets a new size for the Container.

#### func (*Container) Show

```go
func (c *Container) Show()
```
Show sets this container, and all its children, to be visible.

#### func (*Container) Size

```go
func (c *Container) Size() Size
```
Size returns the current size of this container.

#### func (*Container) Visible

```go
func (c *Container) Visible() bool
```
Visible returns true if the container is currently visible, false otherwise.

#### type Device

```go
type Device interface {
	Orientation() DeviceOrientation
	IsMobile() bool
	HasKeyboard() bool
	SystemScaleForWindow(Window) float32

	// Deprecated: Use SystemScaleForWindow instead - system scale can vary depending on window placement
	SystemScale() float32
}
```

Device provides information about the devices the code is running on

#### func  CurrentDevice

```go
func CurrentDevice() Device
```
CurrentDevice returns the device information for the current hardware (via the
driver)

#### type DeviceOrientation

```go
type DeviceOrientation int
```

DeviceOrientation represents the different ways that a mobile device can be held

```go
const (
	// OrientationVertical is the default vertical orientation
	OrientationVertical DeviceOrientation = iota
	// OrientationVerticalUpsideDown is the portrait orientation held upside down
	OrientationVerticalUpsideDown
	// OrientationHorizontalLeft is used to indicate a landscape orientation with the top to the left
	OrientationHorizontalLeft
	// OrientationHorizontalRight is used to indicate a landscape orientation with the top to the right
	OrientationHorizontalRight
)
```

#### type Disableable

```go
type Disableable interface {
	Enable()
	Disable()
	Disabled() bool
}
```

Disableable describes any CanvasObject that can be disabled. This is primarily
used with objects that also implement the Tappable interface.

#### type DoubleTappable

```go
type DoubleTappable interface {
	DoubleTapped(*PointEvent)
}
```

DoubleTappable describes any CanvasObject that can also be double tapped.

#### type DragEvent

```go
type DragEvent struct {
	PointEvent
	DraggedX, DraggedY int
}
```

DragEvent defines the parameters of a pointer or other drag event. The DraggedX
and DraggedY fields show how far the item was dragged since the last event.

#### type Draggable

```go
type Draggable interface {
	Dragged(*DragEvent)
	DragEnd()
}
```

Draggable indicates that a CanvasObject can be dragged. This is used for any
item that the user has indicated should be moved across the screen.

#### type Driver

```go
type Driver interface {
	// Create a new UI Window.
	CreateWindow(string) Window
	// Get a slice containing all app windows.
	AllWindows() []Window

	// Return the size required to render the given string of specified
	// font size and style.
	RenderedTextSize(string, int, TextStyle) Size

	// FileReaderForURI opens a file reader for the given resource indicator.
	// This may refer to a filesystem (typical on desktop) or data from another application.
	FileReaderForURI(string) (FileReadCloser, error)

	// FileWriterForURI opens a file writer for the given resource indicator.
	// This should refer to a filesystem resource as external data will not be writable.
	FileWriterForURI(string) (FileWriteCloser, error)

	// Get the canvas that is associated with a given CanvasObject.
	CanvasForObject(CanvasObject) Canvas
	// Get the position of a given CanvasObject relative to the top/left of a canvas.
	AbsolutePositionForObject(CanvasObject) Position

	// Get the device that the application is currently running on
	Device() Device
	// Start the main event loop of the driver.
	Run()
	// Close the driver and open windows then exit the application.
	Quit()
}
```

Driver defines an abstract concept of a Fyne render driver. Any implementation
must provide at least these methods.

#### type FileReadCloser

```go
type FileReadCloser interface {
	io.ReadCloser
	Name() string
	URI() string
}
```

FileReadCloser represents a cross platform data stream from a file or provider
of data. It may refer to an item on a filesystem or data in another application
that we have access to.

#### func  OpenFileFromURI

```go
func OpenFileFromURI(uri string) (FileReadCloser, error)
```
OpenFileFromURI loads a file read stream from a resource identifier. This is
mostly provided so that file references can be saved using their URI and loaded
again later.

#### type FileWriteCloser

```go
type FileWriteCloser interface {
	io.WriteCloser
	Name() string
	URI() string
}
```

FileWriteCloser represents a cross platform data writer for a file resource.
This will normally refer to a local file resource.

#### func  SaveFileToURI

```go
func SaveFileToURI(uri string) (FileWriteCloser, error)
```
SaveFileToURI loads a file write stream to a resource identifier. This is mostly
provided so that file references can be saved using their URI and written to
again later.

#### type Focusable

```go
type Focusable interface {
	FocusGained()
	FocusLost()
	Focused() bool // Deprecated: this is an internal detail, canvas tracks current focused object

	TypedRune(rune)
	TypedKey(*KeyEvent)
}
```

Focusable describes any CanvasObject that can respond to being focused. It will
receive the FocusGained and FocusLost events appropriately. When focused it will
also have TypedRune called as text is input and TypedKey called when other keys
are pressed.

#### type KeyEvent

```go
type KeyEvent struct {
	Name KeyName
}
```

KeyEvent describes a keyboard input event.

#### type KeyName

```go
type KeyName string
```

KeyName represents the name of a key that has been pressed

```go
const (
	// KeyEscape is the "esc" key
	KeyEscape KeyName = "Escape"
	// KeyReturn is the carriage return (main keyboard)
	KeyReturn KeyName = "Return"
	// KeyTab is the tab advance key
	KeyTab KeyName = "Tab"
	// KeyBackspace is the delete-before-cursor key
	KeyBackspace KeyName = "BackSpace"
	// KeyInsert is the insert mode key
	KeyInsert KeyName = "Insert"
	// KeyDelete is the delete-after-cursor key
	KeyDelete KeyName = "Delete"
	// KeyRight is the right arrow key
	KeyRight KeyName = "Right"
	// KeyLeft is the left arrow key
	KeyLeft KeyName = "Left"
	// KeyDown is the down arrow key
	KeyDown KeyName = "Down"
	// KeyUp is the up arrow key
	KeyUp KeyName = "Up"
	// KeyPageUp is the page up num-pad key
	KeyPageUp KeyName = "Prior"
	// KeyPageDown is the page down num-pad key
	KeyPageDown KeyName = "Next"
	// KeyHome is the line-home key
	KeyHome KeyName = "Home"
	// KeyEnd is the line-end key
	KeyEnd KeyName = "End"

	// KeyF1 is the first function key
	KeyF1 KeyName = "F1"
	// KeyF2 is the second function key
	KeyF2 KeyName = "F2"
	// KeyF3 is the third function key
	KeyF3 KeyName = "F3"
	// KeyF4 is the fourth function key
	KeyF4 KeyName = "F4"
	// KeyF5 is the fifth function key
	KeyF5 KeyName = "F5"
	// KeyF6 is the sixth function key
	KeyF6 KeyName = "F6"
	// KeyF7 is the seventh function key
	KeyF7 KeyName = "F7"
	// KeyF8 is the eighth function key
	KeyF8 KeyName = "F8"
	// KeyF9 is the ninth function key
	KeyF9 KeyName = "F9"
	// KeyF10 is the tenth function key
	KeyF10 KeyName = "F10"
	// KeyF11 is the eleventh function key
	KeyF11 KeyName = "F11"
	// KeyF12 is the twelfth function key
	KeyF12 KeyName = "F12"

	// KeyEnter is the enter/ return key (keypad)
	KeyEnter KeyName = "KP_Enter"

	// Key0 represents the key 0
	Key0 KeyName = "0"
	// Key1 represents the key 1
	Key1 KeyName = "1"
	// Key2 represents the key 2
	Key2 KeyName = "2"
	// Key3 represents the key 3
	Key3 KeyName = "3"
	// Key4 represents the key 4
	Key4 KeyName = "4"
	// Key5 represents the key 5
	Key5 KeyName = "5"
	// Key6 represents the key 6
	Key6 KeyName = "6"
	// Key7 represents the key 7
	Key7 KeyName = "7"
	// Key8 represents the key 8
	Key8 KeyName = "8"
	// Key9 represents the key 9
	Key9 KeyName = "9"
	// KeyA represents the key A
	KeyA KeyName = "A"
	// KeyB represents the key B
	KeyB KeyName = "B"
	// KeyC represents the key C
	KeyC KeyName = "C"
	// KeyD represents the key D
	KeyD KeyName = "D"
	// KeyE represents the key E
	KeyE KeyName = "E"
	// KeyF represents the key F
	KeyF KeyName = "F"
	// KeyG represents the key G
	KeyG KeyName = "G"
	// KeyH represents the key H
	KeyH KeyName = "H"
	// KeyI represents the key I
	KeyI KeyName = "I"
	// KeyJ represents the key J
	KeyJ KeyName = "J"
	// KeyK represents the key K
	KeyK KeyName = "K"
	// KeyL represents the key L
	KeyL KeyName = "L"
	// KeyM represents the key M
	KeyM KeyName = "M"
	// KeyN represents the key N
	KeyN KeyName = "N"
	// KeyO represents the key O
	KeyO KeyName = "O"
	// KeyP represents the key P
	KeyP KeyName = "P"
	// KeyQ represents the key Q
	KeyQ KeyName = "Q"
	// KeyR represents the key R
	KeyR KeyName = "R"
	// KeyS represents the key S
	KeyS KeyName = "S"
	// KeyT represents the key T
	KeyT KeyName = "T"
	// KeyU represents the key U
	KeyU KeyName = "U"
	// KeyV represents the key V
	KeyV KeyName = "V"
	// KeyW represents the key W
	KeyW KeyName = "W"
	// KeyX represents the key X
	KeyX KeyName = "X"
	// KeyY represents the key Y
	KeyY KeyName = "Y"
	// KeyZ represents the key Z
	KeyZ KeyName = "Z"

	// KeySpace is the space key
	KeySpace KeyName = "Space"
	// KeyApostrophe is the key "'"
	KeyApostrophe KeyName = "'"
	// KeyComma is the key ","
	KeyComma KeyName = ","
	// KeyMinus is the key "-"
	KeyMinus KeyName = "-"
	// KeyPeriod is the key "." (full stop)
	KeyPeriod KeyName = "."
	// KeySlash is the key "/"
	KeySlash KeyName = "/"
	// KeyBackslash is the key "\"
	KeyBackslash KeyName = "\\"
	// KeyLeftBracket is the key "["
	KeyLeftBracket KeyName = "["
	// KeyRightBracket is the key "]"
	KeyRightBracket KeyName = "]"
	// KeySemicolon is the key ";"
	KeySemicolon KeyName = ";"
	// KeyEqual is the key "="
	KeyEqual KeyName = "="
	// KeyBackTick is the key "`" on a US keyboard
	KeyBackTick KeyName = "`"
)
```

#### type Layout

```go
type Layout interface {
	// Layout will manipulate the listed CanvasObjects Size and Position
	// to fit within the specified size.
	Layout([]CanvasObject, Size)
	// MinSize calculates the smallest size that will fit the listed
	// CanvasObjects using this Layout algorithm.
	MinSize(objects []CanvasObject) Size
}
```

Layout defines how CanvasObjects may be laid out in a specified Size.

#### type MainMenu

```go
type MainMenu struct {
	Items []*Menu
}
```

MainMenu defines the data required to show a menu bar (desktop) or other
appropriate top level menu.

#### func  NewMainMenu

```go
func NewMainMenu(items ...*Menu) *MainMenu
```
NewMainMenu creates a top level menu structure used by fyne.Window for
displaying a menubar (or appropriate equivalent).

#### type Menu

```go
type Menu struct {
	Label string
	Items []*MenuItem
}
```

Menu stores the information required for a standard menu. A menu can pop down
from a MainMenu or could be a pop out menu.

#### func  NewMenu

```go
func NewMenu(label string, items ...*MenuItem) *Menu
```
NewMenu creates a new menu given the specified label (to show in a MainMenu) and
list of items to display.

#### type MenuItem

```go
type MenuItem struct {
	ChildMenu   *Menu
	IsSeparator bool
	Label       string
	Action      func()
}
```

MenuItem is a single item within any menu, it contains a display Label and
Action function that is called when tapped.

#### func  NewMenuItem

```go
func NewMenuItem(label string, action func()) *MenuItem
```
NewMenuItem creates a new menu item from the passed label and action parameters.

#### func  NewMenuItemSeparator

```go
func NewMenuItemSeparator() *MenuItem
```
NewMenuItemSeparator creates a menu item that is to be used as a separator.

#### type Notification

```go
type Notification struct {
	Title, Content string
}
```

Notification represents a user notification that can be sent to the operating
system.

#### func  NewNotification

```go
func NewNotification(title, content string) *Notification
```
NewNotification creates a notification that can be passed to
App.SendNotification.

#### type OverlayStack

```go
type OverlayStack interface {
	// Add adds an overlay on the top of the overlay stack.
	Add(overlay CanvasObject)
	// List returns the overlays currently on the overlay stack.
	List() []CanvasObject
	// Remove removes the given object and all objects above it from the overlay stack.
	Remove(overlay CanvasObject)
	// Top returns the top-most object of the overlay stack.
	Top() CanvasObject
}
```

OverlayStack is a stack of CanvasObjects intended to be used as overlays of a
Canvas.

#### type PointEvent

```go
type PointEvent struct {
	AbsolutePosition Position // The absolute position of the event
	Position         Position // The relative position of the event
}
```

PointEvent describes a pointer input event. The position is relative to the
top-left of the CanvasObject this is triggered on.

#### type Position

```go
type Position struct {
	X int // The position from the parent's left edge
	Y int // The position from the parent's top edge
}
```

Position describes a generic X, Y coordinate relative to a parent Canvas or
CanvasObject.

#### func  NewPos

```go
func NewPos(x int, y int) Position
```
NewPos returns a newly allocated Position representing the specified
coordinates.

#### func (Position) Add

```go
func (p Position) Add(p2 Position) Position
```
Add returns a new Position that is the result of offsetting the current position
by p2 X and Y.

#### func (Position) IsZero

```go
func (p Position) IsZero() bool
```
IsZero returns whether the Position is at the zero-point.

#### func (Position) Subtract

```go
func (p Position) Subtract(p2 Position) Position
```
Subtract returns a new Position that is the result of offsetting the current
position by p2 -X and -Y.

#### type Preferences

```go
type Preferences interface {
	// Bool looks up a boolean value for the key
	Bool(key string) bool
	// BoolWithFallback looks up a boolean value and returns the given fallback if not found
	BoolWithFallback(key string, fallback bool) bool
	// SetBool saves a boolean value for the given key
	SetBool(key string, value bool)

	// Float looks up a float64 value for the key
	Float(key string) float64
	// FloatWithFallback looks up a float64 value and returns the given fallback if not found
	FloatWithFallback(key string, fallback float64) float64
	// SetFloat saves a float64 value for the given key
	SetFloat(key string, value float64)

	// Int looks up an integer value for the key
	Int(key string) int
	// IntWithFallback looks up an integer value and returns the given fallback if not found
	IntWithFallback(key string, fallback int) int
	// SetInt saves an integer value for the given key
	SetInt(key string, value int)

	// String looks up a string value for the key
	String(key string) string
	// StringWithFallback looks up a string value and returns the given fallback if not found
	StringWithFallback(key, fallback string) string
	// SetString saves a string value for the given key
	SetString(key string, value string)

	// RemoveValue removes a value for the given key (not currently supported on iOS)
	RemoveValue(key string)
}
```

Preferences describes the ways that an app can save and load user preferences

#### type Resource

```go
type Resource interface {
	Name() string
	Content() []byte
}
```

Resource represents a single binary resource, such as an image or font. A
resource has an identifying name and byte array content. The serialised path of
a resource can be obtained which may result in a blocking filesystem write
operation.

#### func  LoadResourceFromPath

```go
func LoadResourceFromPath(path string) (Resource, error)
```
LoadResourceFromPath creates a new StaticResource in memory using the contents
of the specified file.

#### func  LoadResourceFromURLString

```go
func LoadResourceFromURLString(urlStr string) (Resource, error)
```
LoadResourceFromURLString creates a new StaticResource in memory using the body
of the specified URL.

#### type ScrollEvent

```go
type ScrollEvent struct {
	PointEvent
	DeltaX, DeltaY int
}
```

ScrollEvent defines the parameters of a pointer or other scroll event. The
DeltaX and DeltaY represent how large the scroll was in two dimensions.

#### type Scrollable

```go
type Scrollable interface {
	Scrolled(*ScrollEvent)
}
```

Scrollable describes any CanvasObject that can also be scrolled. This is mostly
used to implement the widget.ScrollContainer.

#### type SecondaryTappable

```go
type SecondaryTappable interface {
	TappedSecondary(*PointEvent)
}
```

SecondaryTappable describes a CanvasObject that can be right-clicked or
long-tapped.

#### type Settings

```go
type Settings interface {
	Theme() Theme
	SetTheme(Theme)
	Scale() float32

	AddChangeListener(chan Settings)
}
```

Settings describes the application configuration available.

#### type Shortcut

```go
type Shortcut interface {
	ShortcutName() string
}
```

Shortcut is the interface used to describe a shortcut action

#### type ShortcutCopy

```go
type ShortcutCopy struct {
	Clipboard Clipboard
}
```

ShortcutCopy describes a shortcut copy action.

#### func (*ShortcutCopy) ShortcutName

```go
func (se *ShortcutCopy) ShortcutName() string
```
ShortcutName returns the shortcut name

#### type ShortcutCut

```go
type ShortcutCut struct {
	Clipboard Clipboard
}
```

ShortcutCut describes a shortcut cut action.

#### func (*ShortcutCut) ShortcutName

```go
func (se *ShortcutCut) ShortcutName() string
```
ShortcutName returns the shortcut name

#### type ShortcutHandler

```go
type ShortcutHandler struct {
}
```

ShortcutHandler is a default implementation of the shortcut handler for the
canvasObject

#### func (*ShortcutHandler) AddShortcut

```go
func (sh *ShortcutHandler) AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut))
```
AddShortcut register an handler to be executed when the shortcut action is
triggered

#### func (*ShortcutHandler) TypedShortcut

```go
func (sh *ShortcutHandler) TypedShortcut(shortcut Shortcut)
```
TypedShortcut handle the registered shortcut

#### type ShortcutPaste

```go
type ShortcutPaste struct {
	Clipboard Clipboard
}
```

ShortcutPaste describes a shortcut paste action.

#### func (*ShortcutPaste) ShortcutName

```go
func (se *ShortcutPaste) ShortcutName() string
```
ShortcutName returns the shortcut name

#### type ShortcutSelectAll

```go
type ShortcutSelectAll struct{}
```

ShortcutSelectAll describes a shortcut selectAll action.

#### func (*ShortcutSelectAll) ShortcutName

```go
func (se *ShortcutSelectAll) ShortcutName() string
```
ShortcutName returns the shortcut name

#### type Shortcutable

```go
type Shortcutable interface {
	TypedShortcut(Shortcut)
}
```

Shortcutable describes any CanvasObject that can respond to shortcut commands
(quit, cut, copy, and paste).

#### type Size

```go
type Size struct {
	Width  int // The number of units along the X axis.
	Height int // The number of units along the Y axis.
}
```

Size describes something with width and height.

#### func  MeasureText

```go
func MeasureText(text string, size int, style TextStyle) Size
```
MeasureText uses the current driver to calculate the size of text when rendered.

#### func  NewSize

```go
func NewSize(w int, h int) Size
```
NewSize returns a newly allocated Size of the specified dimensions.

#### func (Size) Add

```go
func (s Size) Add(s2 Size) Size
```
Add returns a new Size that is the result of increasing the current size by s2
Width and Height.

#### func (Size) IsZero

```go
func (s Size) IsZero() bool
```
IsZero returns whether the Size has zero width and zero height.

#### func (Size) Max

```go
func (s Size) Max(s2 Size) Size
```
Max returns a new Size that is the maximum of the current Size and s2.

#### func (Size) Min

```go
func (s Size) Min(s2 Size) Size
```
Min returns a new Size that is the minimum of the current Size and s2.

#### func (Size) Subtract

```go
func (s Size) Subtract(s2 Size) Size
```
Subtract returns a new Size that is the result of decreasing the current size by
s2 Width and Height.

#### func (Size) Union

```go
func (s Size) Union(s2 Size) Size
```
Union returns a new Size that is the maximum of the current Size and s2.
Deprecated: use Max() instead

#### type StaticResource

```go
type StaticResource struct {
	StaticName    string
	StaticContent []byte
}
```

StaticResource is a bundled resource compiled into the application. These
resources are normally generated by the fyne_bundle command included in the Fyne
toolkit.

#### func  NewStaticResource

```go
func NewStaticResource(name string, content []byte) *StaticResource
```
NewStaticResource returns a new static resource object with the specified name
and content. Creating a new static resource in memory results in sharable binary
data that may be serialised to the location returned by CachePath().

#### func (*StaticResource) Content

```go
func (r *StaticResource) Content() []byte
```
Content returns the bytes of the bundled resource, no compression is applied but
any compression on the resource is retained.

#### func (*StaticResource) GoString

```go
func (r *StaticResource) GoString() string
```
GoString converts a Resource object to Go code. This is useful if serialising to
a Go file for compilation into a binary.

#### func (*StaticResource) Name

```go
func (r *StaticResource) Name() string
```
Name returns the unique name of this resource, usually matching the file it was
generated from.

#### type Tappable

```go
type Tappable interface {
	Tapped(*PointEvent)
}
```

Tappable describes any CanvasObject that can also be tapped. This should be
implemented by buttons etc that wish to handle pointer interactions.

#### type TextAlign

```go
type TextAlign int
```

TextAlign represents the horizontal alignment of text within a widget or canvas
object.

```go
const (
	// TextAlignLeading specifies a left alignment for left-to-right languages.
	TextAlignLeading TextAlign = iota
	// TextAlignCenter places the text centrally within the available space.
	TextAlignCenter
	// TextAlignTrailing will align the text right for a left-to-right language.
	TextAlignTrailing
)
```

#### type TextStyle

```go
type TextStyle struct {
	Bold      bool // Should text be bold
	Italic    bool // Should text be italic
	Monospace bool // Use the system monospace font instead of regular
}
```

TextStyle represents the styles that can be applied to a text canvas object or
text based widget.

#### type TextWrap

```go
type TextWrap int
```

TextWrap represents how text longer than the widget's width will be wrapped.

```go
const (
	// TextWrapOff extends the widget's width to fit the text, no wrapping is applied.
	TextWrapOff TextWrap = iota
	// TextTruncate trims the text to the widget's width, no wrapping is applied.
	TextTruncate
	// TextWrapBreak trims the line of characters to the widget's width adding the excess as new line.
	TextWrapBreak
	// TextWrapWord trims the line of words to the widget's width adding the excess as new line.
	TextWrapWord
)
```

#### type Theme

```go
type Theme interface {
	BackgroundColor() color.Color
	ButtonColor() color.Color
	DisabledButtonColor() color.Color
	HyperlinkColor() color.Color
	TextColor() color.Color
	DisabledTextColor() color.Color
	IconColor() color.Color
	DisabledIconColor() color.Color
	PlaceHolderColor() color.Color
	PrimaryColor() color.Color
	HoverColor() color.Color
	FocusColor() color.Color
	ScrollBarColor() color.Color
	ShadowColor() color.Color

	TextSize() int
	TextFont() Resource
	TextBoldFont() Resource
	TextItalicFont() Resource
	TextBoldItalicFont() Resource
	TextMonospaceFont() Resource

	Padding() int
	IconInlineSize() int
	ScrollBarSize() int
	ScrollBarSmallSize() int
}
```

Theme defines the requirements of any Fyne theme.

#### type Widget

```go
type Widget interface {
	CanvasObject

	CreateRenderer() WidgetRenderer
}
```

Widget defines the standard behaviours of any widget. This extends the
CanvasObject - a widget behaves in the same basic way but will encapsulate many
child objects to create the rendered widget.

#### type WidgetRenderer

```go
type WidgetRenderer interface {
	Layout(Size)
	MinSize() Size

	Refresh()
	BackgroundColor() color.Color
	Objects() []CanvasObject
	Destroy()
}
```

WidgetRenderer defines the behaviour of a widget's implementation. This is
returned from a widget's declarative object through the Render() function and
should be exactly one instance per widget in memory.

#### type Window

```go
type Window interface {
	// Title returns the current window title.
	// This is typically displayed in the window decorations.
	Title() string
	// SetTitle updates the current title of the window.
	SetTitle(string)

	// FullScreen returns whether or not this window is currently full screen.
	FullScreen() bool
	// SetFullScreen changes the requested fullScreen property
	// true for a fullScreen window and false to unset this.
	SetFullScreen(bool)

	// Resize this window to the requested content size.
	// The result may not be exactly as desired due to various desktop or
	// platform constraints.
	Resize(Size)

	// RequestFocus attempts to raise and focus this window.
	// This should only be called when you are sure the user would want this window
	// to steal focus from any current focused window.
	RequestFocus()

	// FixedSize returns whether or not this window should disable resizing.
	FixedSize() bool
	// SetFixedSize sets a hint that states whether the window should be a fixed
	// size or allow resizing.
	SetFixedSize(bool)

	// CenterOnScreen places a window at the center of the monitor
	// the Window object is currently positioned on.
	CenterOnScreen()

	// Padded, normally true, states whether the window should have inner
	// padding so that components do not touch the window edge.
	Padded() bool
	// SetPadded allows applications to specify that a window should have
	// no inner padding. Useful for fullscreen or graphic based applications.
	SetPadded(bool)

	// Icon returns the window icon, this is used in various ways
	// depending on operating system.
	// Most commonly this is displayed on the window border or task switcher.
	Icon() Resource

	// SetIcon sets the icon resource used for this window.
	// If none is set should return the application icon.
	SetIcon(Resource)

	// SetMaster indicates that closing this window should exit the app
	SetMaster()

	// MainMenu gets the content of the window's top level menu.
	MainMenu() *MainMenu

	// SetMainMenu adds a top level menu to this window.
	// The way this is rendered will depend on the loaded driver.
	SetMainMenu(*MainMenu)

	SetOnClosed(func())

	// Show the window on screen.
	Show()
	// Hide the window from the user.
	// This will not destroy the window or cause the app to exit.
	Hide()
	// Close the window.
	// If it is the only open window, or the "master" window the app will Quit.
	Close()

	// ShowAndRun is a shortcut to show the window and then run the application.
	// This should be called near the end of a main() function as it will block.
	ShowAndRun()

	// Content returns the content of this window.
	Content() CanvasObject
	// SetContent sets the content of this window.
	SetContent(CanvasObject)
	// Canvas returns the canvas context to render in the window.
	// This can be useful to set a key handler for the window, for example.
	Canvas() Canvas

	//Clipboard returns the system clipboard
	Clipboard() Clipboard
}
```

Window describes a user interface window. Depending on the platform an app may
have many windows or just the one.
