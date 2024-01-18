---
redirect_to:
  - https://docs.fyne.io/api/v2.4/dialog/confiialog

layout: page
tags: [api]
title: Fyne API "dialog.ConfirmDialog"
package: fyne.io/fyne/v2/dialog
---
# dialog.ConfirmDialog
---

```go
import "fyne.io/fyne/v2/dialog"
```

## Usage

#### type ConfirmDialog

```go
type ConfirmDialog struct {
}
```

ConfirmDialog is like the standard Dialog but with an additional confirmation button

#### func  NewConfirm

```go
func NewConfirm(title, message string, callback func(bool), parent fyne.Window) *ConfirmDialog
```
NewConfirm creates a dialog over the specified window for user confirmation. The title is used for the dialog window and message is the content. The callback is executed when the user decides. After creation you should call Show().

#### func  NewCustomConfirm

```go
func NewCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window) *ConfirmDialog
```
NewCustomConfirm creates and returns a dialog over the specified application using custom content. The cancel button will have the dismiss text set and the "OK" will use the confirm text. The response callback is called on user action. The MinSize() of the CanvasObject passed will be used to set the size of the window.

#### func (ConfirmDialog) Hide

```go
func (d ConfirmDialog) Hide()
```

#### func (ConfirmDialog) MinSize

```go
func (d ConfirmDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below


<div class="since">Since: <code>
2.1</code></div>

#### func (ConfirmDialog) Refresh

```go
func (d ConfirmDialog) Refresh()
```

#### func (ConfirmDialog) Resize

```go
func (d ConfirmDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

#### func (*ConfirmDialog) SetConfirmImportance

```go
func (d *ConfirmDialog) SetConfirmImportance(importance widget.Importance)
```
SetConfirmImportance sets the importance level of the confirm button.

Since 2.4

#### func (*ConfirmDialog) SetConfirmText

```go
func (d *ConfirmDialog) SetConfirmText(label string)
```
SetConfirmText allows custom text to be set in the confirmation button

#### func (ConfirmDialog) SetDismissText

```go
func (d ConfirmDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

#### func (ConfirmDialog) SetOnClosed

```go
func (d ConfirmDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (ConfirmDialog) Show

```go
func (d ConfirmDialog) Show()
```
