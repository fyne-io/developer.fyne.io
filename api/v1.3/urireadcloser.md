---
redirect_to:
  - https://docs.fyne.io/api/v1.3/urireadcloser.md

layout: page
tags: [api]
title: Fyne API fyne
---


# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type URIReadCloser

```go
type URIReadCloser interface {
	io.ReadCloser
	Name() string
	URI() URI
}
```

URIReadCloser represents a cross platform data stream from a file or provider of data. It may refer to an item on a filesystem or data in another application that we have access to.
