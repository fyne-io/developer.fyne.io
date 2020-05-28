---
layout: tour

title: Container and Layouts
order: 4

---

In the previous example we saw how to set a `CanvasObject` to the
content of a `Canvas`, but it is not very useful to only show
one visual element. To show more than one item we use the `Container` type.

As the `fyne.Container` is also a `fyne.CanvasObject` we can set it to be
the content of a `fyne.Canvas`. In this example we create 3 text objects
and then place them in a container using the `NewContainer()` function.
As there is no layout set we can move the elements around like you see
with `text2.Move()`.

A `fyne.Layout` implements a method for organising items within a container.
By uncommenting the `NewContainerWithLayout()` line in this example you
alter the container to use a grid layout with 2 columns. Run this code
and try resizing the window to see how the layout automatically configures
the contents of the window. Notice also that the manual position
of `text2` is ignored by the layout code.

To find out more jump to the [`Layout`](../layout/) section of this tour.
