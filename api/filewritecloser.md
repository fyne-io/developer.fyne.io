---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
--
    import "fyne.io/fyne"

## Usage

#### type FileWriteCloser

```go
type FileWriteCloser interface {
	io.WriteCloser
	Name() string
	URI() URI
}
```

FileWriteCloser represents a cross platform data writer for a file resource.
This will normally refer to a local file resource.
