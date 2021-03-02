---
layout: tour

title: List Data
order: 5

permalink: /tour/binding/list

---

To demonstrate how more complex types can be connected
we will look at the `List` widget and how data binding
can make it easier to use. Firstly we create a `StringList`
data binding, which is a list of `String` data type.
Once we have a data of list type we can connect this to
the standard `List` widget. To do so we use the
`widget.NewListWithData` constructor, much like other
widgets.

Comparing this code to the [list tour](/tour/widget/list)
You will see 2 main changes, the first is that we pass
the data type as the first parameter instead of a length
callback function. The second change is the last parameter,
our `UpdateItem` callback. The revised version takes
a `binding.DataItem` value instead of `widget.ListIndexID`.
When using this callback structure we should `Bind`
to the template label widget instead of calling `SetText`.
This means that if any of the strings change in the
data source each affected row of the table will refresh.

In our demo code there is an "Append" `Button`, when
tapped it will append a new value to the data source.
Doing so will automatically trigger the data change
handlers and expand the `List` widget to display the
new data.