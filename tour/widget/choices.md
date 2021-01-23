---
layout: tour

title: Choices
order: 4

---

There are various widgets available to present the user with
a choice, these include a checkbox, radio group and select popup.

The `widget.Check` provides a simple yes/no choice and is created
using a string label. Each of these widgets also takes a
"changed" `func(...)` where the parameter is of the appropriate
type. `widget.NewCheck(..)` therefore takes a `string` parameter for
the label and a `func(bool)` param for the change handler.
You can also use the `Checked` field to get the boolean value.

The radio widget is similar, but the first parameter is a
slice of `string`s that represents each of the options.
The change function expects a `string` parameter this time
to return the currently selected value. Call `widget.NewRadioGroup(...)`
to construct the radio group widget, you can use this reference
later to read the `Selected` field instead of using the change
callback.

The select widget is identical in the constructor signature
as the radio widget. Calling `widget.NewSelect(...)` will
instead show a button that displays a popup when tapped from
which the user can make a selection. This is better for long
lists of options.
