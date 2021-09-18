---
layout: page
tags: [api]
title: Upgrading to v2.1
---

The 2.1 release is fully backward compatible with 2.0.4 and earlier, so upgrading
is as simple as updating the version of code you compile with.
This is different depending on whether or not you use go modules.

## Modules

If your project has a `go.mod` file then you can edit the `require` line to use
version `v2.1.0`, or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne@v2.1.0
```

The next time you build or run your app it will be using the 2.1 release

## GOPATH

If you are not using modules then you will need to update the Fyne checkout in
your go source code. To do this execute the following command:

```bash
go get -u fyne.io/fyne
```

Any apps without a module file will now use the 2.1 release.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may notice.

* Padding in the collection widgets (List, Table, Tree) has been removed, so cells will be the size of the containing widget
* theme.LightTheme() and DarkTheme() have been deprecated, in preference to theme.DefaultTheme()
  if you require overriding user choice you can create a custom theme that extends DefaultTheme
  and in your Color method delegate to the DefaultTheme.Color and pass the variant you require
* fyne.TextStyle now has a TabWidth field, if you use unnamed struct initialisation this may require changes in your code
