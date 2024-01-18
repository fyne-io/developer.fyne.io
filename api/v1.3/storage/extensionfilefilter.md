---
redirect_to:
  - https://docs.fyne.io/api/v1.3/storage/extensionfilefilter.md

layout: page
tags: [api]
title: Fyne API storage
---


# storage
---
```go
import "fyne.io/fyne/storage"
```

## Usage

#### type ExtensionFileFilter

```go
type ExtensionFileFilter struct {
	Extensions []string
}
```

ExtensionFileFilter represents a file filter based on the the ending of file names, for example ".txt" and ".png".

#### func (*ExtensionFileFilter) Matches

```go
func (e *ExtensionFileFilter) Matches(uri fyne.URI) bool
```
Matches returns true if a file URI has one of the filtered extensions.
