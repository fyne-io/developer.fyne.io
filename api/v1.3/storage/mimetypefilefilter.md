---
redirect_to:
  - https://docs.fyne.io/api/v1.3/storage/mimetypefilefilter.md

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

#### type MimeTypeFileFilter

```go
type MimeTypeFileFilter struct {
	MimeTypes []string
}
```

MimeTypeFileFilter represents a file filter based on the files mime type, for example "image/*", "audio/mp3".

#### func (*MimeTypeFileFilter) Matches

```go
func (mt *MimeTypeFileFilter) Matches(uri fyne.URI) bool
```
Matches returns true if a file URI has one of the filtered mimetypes.
