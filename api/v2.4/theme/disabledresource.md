---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "theme.DisabledResource"
package: fyne.io/fyne/v2/theme
---
# theme.DisabledResource
---
```go
import "fyne.io/fyne/v2/theme"
```

## Usage

#### type DisabledResource

```go
type DisabledResource struct {
}
```

DisabledResource is a resource wrapper that will return an appropriate resource colorized by the current theme's `DisabledColor` color.

#### func  NewDisabledResource

```go
func NewDisabledResource(res fyne.Resource) *DisabledResource
```
NewDisabledResource creates a resource that adapts to the current theme's DisabledColor setting.

#### func (*DisabledResource) Content

```go
func (res *DisabledResource) Content() []byte
```
Content returns the disabled style content of the correct resource for the current theme

#### func (*DisabledResource) Name

```go
func (res *DisabledResource) Name() string
```
Name returns the resource source name prefixed with `disabled_` (used for caching)
