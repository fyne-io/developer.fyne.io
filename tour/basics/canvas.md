---
layout: tour

title: Canvas and CanvasObject
order: 3

---

In Fyne a `Canvas` is the area which an application is drawn within.
Each window has a canvas which you can access with `Window.Canvas()`
but usually you will find functions on `Window` that avoid accessing
the canvas.

Everything that can be drawn in Fyne is a type of `CanvasObject`.
The example here opens a new window and then shows different types of
primitive graphical element by setting the content of the window canvas.
There are many ways that each type of object can be customised as
shown with the text and circle examples. 

As well as changing the content shown using `Canvas.SetContent()` it is
possible to change the content that is shown. If, for example, you
change the colour of a rectangle you can request a refresh of this
existing component using `canvas.Refresh(rect)`.
