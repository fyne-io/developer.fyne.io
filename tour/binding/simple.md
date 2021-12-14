---
layout: tour

title: Binding Simple Widgets
order: 2

---

The simplest form of binding a widget is to pass it
a bound item as a value instead of a static value.
Many widgets provide a `WithData` constructor that will
accept a typed data binding item. To set up the binding
all you need to do is pass the right type in.

Although this may not look like much of a benefit in the
initial code you can then see how it ensures that the
displayed content is always up to date with the source
of the data.
You will notice how we did not need to call `Refresh()`
on the `Label` widget or even keep a reference to it
and yet it updates appropriately.

In the next step we look at how to set up widgets 
that edit values through [two way](/tour/binding/twoway) binding.
