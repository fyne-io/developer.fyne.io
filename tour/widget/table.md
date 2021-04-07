---
layout: tour

title: Table
order: 9
---

The `Table` widget is like the `List` widget (another of the toolkit's collection widgets) with a two-dimensional index. 
Like `List` this is designed to help build really performant
interfaces when lots of data is being presented. 
Because of this the widget is not created with all the data embedded, but instead calls out to the data source when needed.

The `Table` uses callback functions to ask for data when it is required.
There are 3 main callbacks, `Length`, `CreateCell` and `UpdateCell`. The Length callback (passed first) is the simplest,
it returns how many items are in the data to be presented, the two ints it returns represent the row and colum count.
The other two relate to the content templates.

The `CreateItem` callback returns a new template object, just like list.
The difference being that `MinSize` will define the standard size of each cell, and the minimum size of the table (it shows at least one cell).
AAs previously the `UpdateItem` is called to apply data to a cell template. The index passed in is the same `(row, col)` int pair.

In the next tour section we explore [data bindings](/tour/binding/).
