---
layout: tour

title: Entry
order: 4

---

The entry widget is used for user input of simple text content.
An entry can be created with a simple `widget.NewEntry()`
constructing function. When you create the widget keep a
reference so that you can access it's `Text` field later.
It is also possible to use the `OnChanged` callback function
to be notified every time the content changes.

Entry widgets can also be set to read only content using
the `ReadOnly` field or calling `SetReadOnly(true)`.
You can also set a `PlaceHolder` text and also set the entry
to `MultiLine` to accept more than one line of text.

You can also create a password entry (where the content is
obscured) using the `NewPasswordEntry()` function.
