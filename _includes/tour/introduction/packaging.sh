#!/bin/sh

go get fyne.io/fyne/cmd/fyne

go build
fyne package -icon mylogo.png

# result is a platform specific package
# for the current operating system.