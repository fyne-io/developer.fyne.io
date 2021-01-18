---
layout: page
tags: [api]
title: Upgrading to v2.0
---

As our first "major" release the 2.0 release contains some breaking changes.
Upgrading to v2.0.0 may break some applications so we recommend reading this
document carefully before performing the update

## Modules

If your project has a `go.mod` file then you can edit the `require` line to use
version `v2.0.0`, or you can execute the following command inside the directory:

```bash
go get fyne.io/fyne@v2.0.0
```

The next time you build or run your app it will be using the 2.0 release

## GOPATH

If you are not using modules then you will need to update the Fyne checkout in
your go source code. To do this execute the following command:

```bash
go get -u fyne.io/fyne
```

Any apps without a module file will now use the 2.0 release.

## Changes

TODO the following section will itemise all the notable changes.

### Breaking changes

These items must be considered before upgrading for your app to compile and run:

*

### Other changes

These changes could cause behaviour inconsistency, so be sure to test your app
after performing the upgrade

*

