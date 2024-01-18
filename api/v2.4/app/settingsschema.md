---
redirect_to:
  - https://docs.fyne.io/_gen/api.md

layout: page
tags: [api]
title: Fyne API "app.SettingsSchema"
package: fyne.io/fyne/v2/app
---
# app.SettingsSchema
---
```go
import "fyne.io/fyne/v2/app"
```

## Usage

#### type SettingsSchema

```go
type SettingsSchema struct {
	// these items are used for global settings load
	ThemeName         string  `json:"theme"`
	Scale             float32 `json:"scale"`
	PrimaryColor      string  `json:"primary_color"`
	CloudName         string  `json:"cloud_name"`
	CloudConfig       string  `json:"cloud_config"`
	DisableAnimations bool    `json:"no_animations"`
}
```

SettingsSchema is used for loading and storing global settings

#### func (*SettingsSchema) StoragePath

```go
func (sc *SettingsSchema) StoragePath() string
```
StoragePath returns the location of the settings storage
