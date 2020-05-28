---
layout: tour

title: Grid Layout
order: 2

---

The grid layout lays out the elements of a container in a grid pattern
with a fixed number of columns. Items will fill a single row until the
number of columns is met, after this a new row will be created.
Vertical space will be split equally between each of the rows of objects.

You create a grid layout using `layout.NewGridLayout(cols)` where cols
is the number of items (columns) you wish to have in each row. This
layout is then passed as the first parameter to
`fyne.NewContainerWithLayout(...)`.

If you resize the container then each of the cells will resize equally
to share the available space.
