---
redirect_to:
  - https://docs.fyne.io/api/v1.4/storage.md

layout: page
tags: [api]
title: Fyne API "fyne.Storage"
---


# fyne.Storage
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Storage

```go
type Storage interface {
	RootURI() URI
}
```

Storage is used to manage file storage inside an application sandbox
