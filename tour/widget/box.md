---
layout: tour

title: Box
order: 3

---

The box widget is a simple horizontal or vertical container
that uses the [box layout](../layout/boxlayout.md) to lay out
the child components. You can pass the objects to include in
the `widget.NewHBox()` or `widget.NewVBox()`constructor functions.

It is also possible to add items to a box widget after it is
created using `Append()` (to add after existing content)
or `Prepend()` (to add before the content).
