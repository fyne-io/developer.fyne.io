---
redirect_to:
  - https://docs.fyne.io/api/v2.0/clipboard

layout: page
tags: [api]
title: Fyne API "fyne.Clipboard"
---


# fyne.Clipboard
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Clipboard

```go
type Clipboard interface {
	// Content returns the clipboard content
	Content() string
	// SetContent sets the clipboard content
	SetContent(content string)
}
```

Clipboard represents the system clipboard interface
