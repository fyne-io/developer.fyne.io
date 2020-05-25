---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
--
    import "fyne.io/fyne"

## Usage

#### type FileReadCloser

```go
type FileReadCloser interface {
	io.ReadCloser
	Name() string
	URI() URI
}
```

FileReadCloser represents a cross platform data stream from a file or provider
of data. It may refer to an item on a filesystem or data in another application
that we have access to.
