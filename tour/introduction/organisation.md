---
layout: tour

title: Organisation and Packages
order: 5

---

The Fyne project is split into many packages, each providing different
types of functionality. They are as follows:

`fyne.io/fyne`
: This import provides the basic definitions common to all Fyne code 
: including data types and interfaces.

`fyne.io/fyne/app`
: The app package provides the APIs that start a new application.
: Normally you only require `app.New()`.

`fyne.io/fyne/canvas`
: The canvas package provides all of the drawing APIs within Fyne.
: The complete Fyne toolkit is made up of these primitive graphical types.

`fyne.io/fyne/dialog`
: Dialog windows such as confirm or error are handled by this package.

`fyne.io/fyne/layout`
: The layout package provides various layout implementations for use
: with containers (discussed in a later tutorial).

`fyne.io/fyne/test`
: Applications can be tested more easily using the tools within the test
: package.

`fyne.io/fyne/widget`
: Most graphical applications are created using a collection of widgets.
: All the widgets and interactive elements within Fyne are in this package.
