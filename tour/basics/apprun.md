---
layout: tour

title: Application and RunLoop
order: 1

permalink: /tour/basics/

redirect_from:
- /tour/basics/apprun

---

For a GUI application to work it needs to run an event loop
(sometimes called a runloop) that processes user interactions
and drawing events. In Fyne this is started using the `App.Run()`
or `Window.ShowAndRun()` functions. One of these must be called
from the end of your setup code in the `main()` function.

An application should only have one runloop and so you should only
call `Run()` once in your code. Calling it a second time will cause
errors.

For desktop runtimes an app can be quit directly by calling `App.Quit()`
(mobile apps do not support this) - normally not needed in developer code.
An application will also quit once all the windows are closed.
See also that functions executed after `Run()` will not be called
until the application exits.
