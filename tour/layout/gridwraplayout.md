---
layout: tour

title: Grid Wrap Layout
order: 3

---

Like the previous grid layout, the grid wrap layout creates an arrangement
of elements in a grid pattern. However this grid does not have a set
number of columns, instead it uses a fixed size for each cell and
then flows the content to as many rows as is needed to display the items.

You create a grid wrap layout using `layout.NewGridWrapLayout(size)`
where size specifies the size to apply to all child elements.
This layout is then passed as the first parameter to
`fyne.NewContainerWithLayout(...)`.
The number of columns and rows will be calculated based on the current
size of the container.

Initially a grid wrap layout will have a single column, if you resize it
(as illustrated in the code comment to the right) it will rearrange
the child elements to fill the space.
