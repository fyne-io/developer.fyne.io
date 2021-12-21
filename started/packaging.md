---
layout: page
title: Packaging

order: 50
---

## Packaging for Desktop

---

Packaging a graphical app for distribution can be complex - Windows executables
need embedded icons, macOS apps are bundles and with Linux there are various
metadata files that should get installed. What a hassle!

Thankfully the "fyne" app has a "package" command that can handle this automatically. Just specifying the target OS and any required metadata (such as icon) will generate the appropriate package. The icon conversion will be done automatically for .icns or .ico so just provide a .png file :). All you need is to have the application already built for the target platform...

```
go get fyne.io/fyne/v2/cmd/fyne
fyne package -os darwin -icon myapp.png
```

Will create myapp.app, a complete bundle structure for distribution to macOS users. You could then build the linux and Windows versions too...

```
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png
```

These commands will create:

  * myapp.tar.gz that contains a folder structure starting at usr/local/ that a Linux user could expand to the root of their system.
  * myapp.exe (after the second build, which is required for a windows package) will have the icon and app metadata embedded.

If you just want to install the desktop app on your computer then you can make
use of the helpful install subcommand. For example to install your current
application system wide you could simply execute the following:

```
fyne install -icon myapp.png
```

All of these commands also support a default icon file of `Icon.png` so that you
can avoid typing the parameter for each execution.

// TODO merge this in
Packaging for multiple operating systems can be a complex task.
Graphical applications typically have icons and metadata associated
with them as well as specific formats required to integrate with each
environment.

The `fyne` command provides support for preparing applications to be
distributed across all the platforms the toolkit supports.
Running "fyne package" will create an application ready to be installed
on your computer and to be distributed to other computers by simply
copying the created files from the current directory.

For Windows it will create a `.exe` file with icons embedded. For a
macOS computer it will create a `.app` bundle and for Linux it will
generate a `.tar.xz` file that can be installed in the usual manner
(or by running `make install` inside the extracted folder).

Of course you can still run your applications using the standard Go
tools if you prefer.
