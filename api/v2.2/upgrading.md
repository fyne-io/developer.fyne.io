---
redirect_to:
  - https://docs.fyne.io/api/v2.2/upgrading

layout: page
tags: [api]
title: Upgrading to v2.2
---


The 2.2 release is fully backward compatible with 2.0.4 and earlier, so upgrading
is as simple as updating the version of code you compile with.
This is different depending on whether or not you use go modules.

## Modules

If your project has a `go.mod` file then you can edit the `require` line to use
version `v2.2.0`, or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne/v2@v2.2.0
```

The next time you build or run your app it will be using the 2.2 release

## GOPATH

If you are not using modules then you will need to update the Fyne checkout in
your go source code. To do this execute the following command:

```bash
go get -u fyne.io/fyne/v2
```

Any apps without a module file will now use the 2.2 release.

## Fyne command

You should update the `fyne` tool for v2.2.0 to get the web driver "fyne serve" support.
You can make the upgrade by using the `go get` command similarly to above:

```bash
go get -u fyne.io/fyne/v2/cmd/fyne
```

After that completes, check you have the new version installed by running `fyne version`.

## Changes

Although this release is backwards compatible so your code will compile and
run as expected, there are some changes which you may notice.

* Toolbar item constructors now return concrete types instead of ToolbarItem
* Low importance buttons no longer draw button color as a background
* ProgressBar widget height is now consistent with other widgets
