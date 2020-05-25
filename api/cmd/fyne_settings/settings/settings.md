---
layout: page
tags: [api]
title: Fyne API settings
---

# settings
--
    import "fyne.io/fyne/cmd/fyne_settings/settings"

## Usage

#### type Settings

```go
type Settings struct {
}
```

Settings gives access to user interfaces to control Fyne settings

#### func  NewSettings

```go
func NewSettings() *Settings
```
NewSettings returns a new settings instance with the current configuration
loaded

#### func (*Settings) AppearanceIcon

```go
func (s *Settings) AppearanceIcon() fyne.Resource
```
AppearanceIcon returns the icon for appearance settings

#### func (*Settings) LoadAppearanceScreen

```go
func (s *Settings) LoadAppearanceScreen(w fyne.Window) fyne.CanvasObject
```
LoadAppearanceScreen creates a new settings screen to handle appearance
configuration
