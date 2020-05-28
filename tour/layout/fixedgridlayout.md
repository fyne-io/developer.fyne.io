---
layout: tour

title: Fixed Grid Layout
order: 3

---

like the previous grid layout, the fixed grid creates an arrangement
of elements in a grid pattern. However this grid does not have a set
number of columns, instead it uses a fixed size for each cell and
then flows the content to as many rows as is needed to display the items.

You create a fixed grid layout using `layout.NewFixedGridLayout(size)`
where size specifies the size to apply to all child elements.
This layout is then passed as the first parameter to
`fyne.NewContainerWithLayout(...)`.
The number of columns and rows will be calculated based on the current
size of the container.

Initially a fixed grid will have a single column, if you resize it
(as illustrated in the code comment to the right) it will rearrange
the child elements to fill the space.
