---
redirect_to:
  - https://docs.fyne.io/started/webapp
layout: page
title: Run in a Browser

order: 78
---
Fyne applications can also be run over the web using a standard web browser!
Due to the wide variety of standards and strenghts of different engines this is a little more complex.

A web app created with Fyne will deliver a WASM runtime as well as a JavaScript code bundle,
this allows the web page generated to look at the current browser and pick the right implementation
for it's strengths. This leads to a great user experience on most systems.

To prepare your app to be used over the web we use the "fyne" cli app again, which has a
"serve" command for quick testing

```
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne serve
```

You will see, after a few moments, that a web server has been started on port :8080.
Just open your web browser to https://localhost:8080 and you can use your app!

The tools used in this process are very sensitive to Go versions being used, so it is possible there will be an error relating to a mis-matched version.
If you are using Go more recent than 1.18 you should ensure you have a copy of the 1.18 release
and reference it in your environment, such as (for a [homebrew](https://brew.sh) based install):

```
export GOPHERJS_GOROOT=/opt/homebrew/Cellar/go@1.18/1.18.10/libexec
```

You can find out more about this on the [GopherJS](https://github.com/gopherjs/gopherjs) documentation.

### Packaging for web distribution

The `fyne serve` command is great for local testing, but just like other platforms you'll want
to be able to distribute your app as well. To prepare the files for upload just use the
`fyne package` command like with regular [Packaging](/started/packaging).

```
fyne package -os web
```

You can also choose to package only for WASM or JavaScript instead of the auto-detecting setup:

```
fyne package -os wasm
fyne package -os js
```

### Demo

You can see a Fyne app in action to test on any of your devices by visiting [demo.fyne.io](https://demo.fyne.io/).

### Limitations

As of release v2.4.0 the web driver is not 100% complete, so your app may not be able to use
the following features:

* Multiple windows (but dialogs all function inside the current window)
* Storage of Documents and Preferences

These issues are being worked on and will be resolved in a future release.
