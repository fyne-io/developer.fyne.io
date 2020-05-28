---
layout: tour

title: Max Layout
order: 7

---

The `layout.MaxLayout` is the simplest layout, it sets all items in
the container to be the same size as the container. This is not
often useful in general containers but can be suitable when composing
widgets

The max layout will expand the container to be at least the size of the
largest item's minimum size. The objects will be drawn in the order
the are passed to the container, with the last being drawn top-most.

Now that we know how to lay out a user interface we will
move on to the [`Widget`s](../widget/) that are available
to create your application.
