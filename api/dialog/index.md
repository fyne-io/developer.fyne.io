---
layout: page
tags: [api]
title: Fyne API dialog
---

# dialog
--
    import "fyne.io/fyne/dialog"

Package dialog defines standard dialog windows for application GUIs

## Usage

#### func  ShowConfirm

```go
func ShowConfirm(title, message string, callback func(bool), parent fyne.Window)
```
ShowConfirm shows a dialog over the specified window for a user confirmation.
The title is used for the dialog window and message is the content. The callback
is executed when the user decides.

#### func  ShowCustom

```go
func ShowCustom(title, dismiss string, content fyne.CanvasObject, parent fyne.Window)
```
ShowCustom shows a dialog over the specified application using custom content.
The button will have the dismiss text set. The MinSize() of the CanvasObject
passed will be used to set the size of the window.

#### func  ShowCustomConfirm

```go
func ShowCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window)
```
ShowCustomConfirm shows a dialog over the specified application using custom
content. The cancel button will have the dismiss text set and the "OK" will use
the confirm text. The response callback is called on user action. The MinSize()
of the CanvasObject passed will be used to set the size of the window.

#### func  ShowError

```go
func ShowError(err error, parent fyne.Window)
```
ShowError shows a dialog over the specified window for an application error. The
title and message are extracted from the provided error.

#### func  ShowFileOpen

```go
func ShowFileOpen(callback func(fyne.FileReadCloser, error), parent fyne.Window)
```
ShowFileOpen shows a file dialog allowing the user to choose a file to open. The
dialog will appear over the window specified.

#### func  ShowFileSave

```go
func ShowFileSave(callback func(fyne.FileWriteCloser, error), parent fyne.Window)
```
ShowFileSave shows a file dialog allowing the user to choose a file to save to
(new or overwrite). If the user chooses an existing file they will be asked if
they are sure. The dialog will appear over the window specified.

#### func  ShowInformation

```go
func ShowInformation(title, message string, parent fyne.Window)
```
ShowInformation shows a dialog over the specified window for user information.
The title is used for the dialog window and message is the content.

#### type ConfirmDialog

```go
type ConfirmDialog struct {
}
```

ConfirmDialog is like the standard Dialog but with an additional confirmation
button

#### func  NewConfirm

```go
func NewConfirm(title, message string, callback func(bool), parent fyne.Window) *ConfirmDialog
```
NewConfirm creates a dialog over the specified window for user confirmation. The
title is used for the dialog window and message is the content. The callback is
executed when the user decides. After creation you should call Show().

#### func (ConfirmDialog) Hide

```go
func (d ConfirmDialog) Hide()
```

#### func (ConfirmDialog) Layout

```go
func (d ConfirmDialog) Layout(obj []fyne.CanvasObject, size fyne.Size)
```

#### func (ConfirmDialog) MinSize

```go
func (d ConfirmDialog) MinSize(obj []fyne.CanvasObject) fyne.Size
```

#### func (*ConfirmDialog) SetConfirmText

```go
func (d *ConfirmDialog) SetConfirmText(label string)
```
SetConfirmText allows custom text to be set in the confirmation button

#### func (ConfirmDialog) SetDismissText

```go
func (d ConfirmDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (ConfirmDialog) SetOnClosed

```go
func (d ConfirmDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is
closed

#### func (ConfirmDialog) Show

```go
func (d ConfirmDialog) Show()
```

#### type Dialog

```go
type Dialog interface {
	Show()
	Hide()
	SetDismissText(label string)
	SetOnClosed(closed func())
}
```

Dialog is the common API for any dialog window with a single dismiss button

#### func  NewCustom

```go
func NewCustom(title, dismiss string, content fyne.CanvasObject, parent fyne.Window) Dialog
```
NewCustom creates and returns a dialog over the specified application using
custom content. The button will have the dismiss text set. The MinSize() of the
CanvasObject passed will be used to set the size of the window.

#### func  NewCustomConfirm

```go
func NewCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window) Dialog
```
NewCustomConfirm creates and returns a dialog over the specified application
using custom content. The cancel button will have the dismiss text set and the
"OK" will use the confirm text. The response callback is called on user action.
The MinSize() of the CanvasObject passed will be used to set the size of the
window.

#### func  NewError

```go
func NewError(err error, parent fyne.Window) Dialog
```
NewError creates a dialog over the specified window for an application error.
The title and message are extracted from the provided error. After creation you
should call Show().

#### func  NewInformation

```go
func NewInformation(title, message string, parent fyne.Window) Dialog
```
NewInformation creates a dialog over the specified window for user information.
The title is used for the dialog window and message is the content. After
creation you should call Show().

#### type ProgressDialog

```go
type ProgressDialog struct {
}
```

ProgressDialog is a simple dialog window that displays text and a progress bar.

#### func  NewProgress

```go
func NewProgress(title, message string, parent fyne.Window) *ProgressDialog
```
NewProgress creates a progress dialog and returns the handle. Using the returned
type you should call Show() and then set its value through SetValue().

#### func (ProgressDialog) Hide

```go
func (d ProgressDialog) Hide()
```

#### func (ProgressDialog) Layout

```go
func (d ProgressDialog) Layout(obj []fyne.CanvasObject, size fyne.Size)
```

#### func (ProgressDialog) MinSize

```go
func (d ProgressDialog) MinSize(obj []fyne.CanvasObject) fyne.Size
```

#### func (ProgressDialog) SetDismissText

```go
func (d ProgressDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (ProgressDialog) SetOnClosed

```go
func (d ProgressDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is
closed

#### func (*ProgressDialog) SetValue

```go
func (p *ProgressDialog) SetValue(v float64)
```
SetValue updates the value of the progress bar - this should be between 0.0 and
1.0.

#### func (ProgressDialog) Show

```go
func (d ProgressDialog) Show()
```

#### type ProgressInfiniteDialog

```go
type ProgressInfiniteDialog struct {
}
```

ProgressInfiniteDialog is a simple dialog window that displays text and a
infinite progress bar.

#### func  NewProgressInfinite

```go
func NewProgressInfinite(title, message string, parent fyne.Window) *ProgressInfiniteDialog
```
NewProgressInfinite creates a infinite progress dialog and returns the handle.
Using the returned type you should call Show().

#### func (*ProgressInfiniteDialog) Hide

```go
func (d *ProgressInfiniteDialog) Hide()
```
Hide this dialog and stop the infinite progress goroutine

#### func (ProgressInfiniteDialog) Layout

```go
func (d ProgressInfiniteDialog) Layout(obj []fyne.CanvasObject, size fyne.Size)
```

#### func (ProgressInfiniteDialog) MinSize

```go
func (d ProgressInfiniteDialog) MinSize(obj []fyne.CanvasObject) fyne.Size
```

#### func (ProgressInfiniteDialog) SetDismissText

```go
func (d ProgressInfiniteDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (ProgressInfiniteDialog) SetOnClosed

```go
func (d ProgressInfiniteDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is
closed

#### func (ProgressInfiniteDialog) Show

```go
func (d ProgressInfiniteDialog) Show()
```
