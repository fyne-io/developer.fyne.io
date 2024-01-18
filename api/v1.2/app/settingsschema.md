---
redirect_to:
  - https://docs.fyne.io/api/v1.2/app/settingsschema

layout: page
tags: [api]
title: Fyne API app
---


# app
---
```go
import "fyne.io/fyne/app"
```

## Usage

#### type SettingsSchema

```go
type SettingsSchema struct {
	// these items are used for global settings load
	ThemeName string  `json:"theme"`
	Scale     float32 `json:"scale"`
}
```

SettingsSchema is used for loading and storing global settings

#### func (*SettingsSchema) StoragePath

```go
func (sc *SettingsSchema) StoragePath() string
```
StoragePath returns the location of the settings storage
