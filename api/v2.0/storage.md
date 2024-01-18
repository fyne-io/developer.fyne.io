---
redirect_to:
  - https://docs.fyne.io/api/v2.0/storage.md

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
}
```

Storage is used to manage file storage inside an application sandbox
