---
layout: tour

title: Packaging and Distribution
order: 6

include: sh
---

Packaging for multiple operating systems can be a complex task.
Graphical applications typically have icons and metadata associated
with them as well as specific formats required to integrate with each
environment.

They `fyne` command provides support for preparing applications to be
distributed across all the platforms the toolkit supports.
Running "fyne package" will create an application ready to be installed
on your computer and to be distributed to other computers by simply
copying the created files from the current directory.

For Windows it will create a `.exe` file with icons embedded. For a
macOS computer it will create a `.app` bundle and for Linux it will
generate a `.tar.xz` file that can be installed in the usual manner.

Of course you can still run your applications using the standard Go
tools if you prefer.
