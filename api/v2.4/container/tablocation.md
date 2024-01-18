---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "container.TabLocation"
package: fyne.io/fyne/v2/container
---
# container.TabLocation
---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type TabLocation

```go
type TabLocation int
```

TabLocation is the location where the tabs of a tab container should be rendered


<div class="since">Since: <code>
1.4</code></div>

```go
const (
	TabLocationTop TabLocation = iota
	TabLocationLeading
	TabLocationBottom
	TabLocationTrailing
)
```
TabLocation values
