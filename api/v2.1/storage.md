---
redirect_to:
  - https://docs.fyne.io/api/v2.1/storage.md

layout: page
tags: [api]
title: Fyne API "fyne.Storage"
---


# fyne.Storage
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Storage

```go
type Storage interface {
	RootURI() URI

	Create(name string) (URIWriteCloser, error)
	Open(name string) (URIReadCloser, error)
	Save(name string) (URIWriteCloser, error)
	Remove(name string) error

	List() []string
}
```

Storage is used to manage file storage inside an application sandbox. The files managed by this interface are unique to the current application.
