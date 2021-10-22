---
layout: page
title: App Metadata

order: 75
---

## App Metadata

Since release v2.1.0 of the `fyne` command we support a metadata file that allows you to store
information about your appliacation in the repository.
This file is optional, but can help to avoid having to remember specific build parameters for
each package and release command.

The file should be named `FyneApp.toml` in the directory where you run the `fyne` command
(this is normally the `main` package). The contents of the file are as follows:

```toml
Website = "https://example.com"

[Details]
Icon = "Icon.png"
Name = "My App"
ID = "com.example.app"
Version = "1.0.0"
Build = 1
```

The fyne tool will use this file if it is found, many mandatory command parameters are not required
if the metadata is present. You can still override these values by using command line parameters.