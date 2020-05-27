---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
--
    import "fyne.io/fyne/widget"

## Usage

#### type SelectEntry

```go
type SelectEntry struct {
	Entry
}
```

SelectEntry is an input field which supports selecting from a fixed set of options.

#### func  NewSelectEntry

```go
func NewSelectEntry(options []string) *SelectEntry
```
NewSelectEntry creates a SelectEntry.

#### func (*SelectEntry) MinSize

```go
func (e *SelectEntry) MinSize() fyne.Size
```
MinSize returns the minimal size of the select entry. Implements: fyne.Widget

#### func (*SelectEntry) Resize

```go
func (e *SelectEntry) Resize(size fyne.Size)
```
Resize changes the size of the select entry. Implements: fyne.Widget

#### func (*SelectEntry) SetOptions

```go
func (e *SelectEntry) SetOptions(options []string)
```
SetOptions sets the options the user might select from.
