---
layout: tour

title: Window Handling
order: 2

---

Windows are created using `App.NewWindow()` and need to be shown using
the `Show()` function. The helper method `ShowAndRun()` on `fyne.Window`
allows you to show your window and run the application at the same time.

If you wish to show a second window you must only call the `Show()`
function. This is illustrated in the `showAnother` function.

By default a window will be the right size to show its content
by checking the `MinSize()` function (more on that in later examples).
You can set a larger size by calling the `Window.Resize()` function.

Be aware that the desktop environment may have constraints that cause
windows to be smaller than requested.
