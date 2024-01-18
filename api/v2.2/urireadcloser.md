---
redirect_to:
  - https://docs.fyne.io/api/v2.2/urireadcloser.md

layout: page
tags: [api]
title: Fyne API "fyne.URIReadCloser"
---


# fyne.URIReadCloser
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type URIReadCloser

```go
type URIReadCloser interface {
	io.ReadCloser

	URI() URI
}
```

URIReadCloser represents a cross platform data stream from a file or provider of data. It may refer to an item on a filesystem or data in another application that we have access to.
