---
layout: tour

title: Data Conversion
order: 4

permalink: /tour/binding/conversion

---

So far we have used data binding where the type of
data matches the output type (for example `String` and
`Label` or `Entry`). Often it will be desirable to
present data that is not already in the correct format.
To do this the `binding` package provides a number
of helpful conversion functions.

Most commonly this will be used to convert different
types of data into strings for displaying in `Label`
or `Entry` widgets. See in the code how we can convert
a `Float` to `String` using `binding.FloatToString`.
The original value can be edited by moving the slider.
Each time the data changes it will run the conversion
code and update any connected widgets.

It is also possible to use format strings to add more
natural output for the user interface.
You can see that our `short` binding is also converting
a `Float` to `String` but by using a `WithFormat` helper
we can pass a format string (similar to the `fmt` package)
to provide a custom output.

Lastly in this section we will look at [list](/tour/binding/list) data.