---
layout: tour

title: Form Layout
order: 5

---

The `layout.FormLayout` is like a 2 column [grid layout](gridlayout.html)
but tweaked for layout out forms in an application.
The height of each item will be the larger of the two minimum heights
in each row. The width of the left item will be the largest minimum
width of all items in the first column whilst the second item in each
row will expand to fill the space.

This layout is more typically used within the `widget.Form` but it can
be used directly with `layout.NewFormLayout()` passed to the first
parameter of `fyne.NewContainerWithLayout(...)`.
