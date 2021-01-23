---
layout: tour

title: Form Layout
order: 5

---

The `layout.FormLayout` is like a 2 column [grid layout](gridlayout.html)
but tweaked to lay out forms in an application.
The height of each item will be the larger of the two minimum heights
in each row. The width of the left item will be the largest minimum
width of all items in the first column whilst the second item in each
row will expand to fill the space.

This layout is more typically used within the `widget.Form` (for validation, submit and cancel buttons, etc) but it can
also be used directly with `layout.NewFormLayout()` passed to the first
parameter of `container.New(...)`.
