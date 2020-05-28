---
layout: tour

title: Form
order: 6

---

The form widget is used to lay out many input fields with
labels and optional cancel and submit buttons. In it's most
bare form it aligns labels to the left of each input widget.
By setting OnCancel or OnSubmit the form will add a button
bar with the specified handlers called when appropriate.

The widget can be created with `widget.NewForm(...)` passing
a list of `widget.FormItem`s, or by using the
`&widget.Form{}` syntax illustrated in the example.
There is also a helpful 'Form.Append(label, widget)` that
can be used for an alternative syntax.

In this example we create two entries, one of which is a
"multiline" (like HTML TextArea) to hold values.
There is an OnSubmit handler which prints the information
before closing the window and therefore the application.
