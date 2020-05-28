---
layout: tour

title: Toolbar
order: 9

---

The toolbar widget creates a row of action buttons using
icons to represent each. The `widget.NewToolbar(...)`
constructor function takes a list of `widget.ToolbarItem`
parameters. The builtin types of toolbar items are action,
separator and spacer.

The most used item is an action that is created using the
`widget.NewToolbarItemAction(..)` function. An action takes
two parameters, first being the icon resource to draw and
the latter is the `func()` to call when tapped. This creates
a standard toolbar button.

You can use `widget.NewToolbarSeparator()` to create a
small divider between items in a toolbar (usually a thin
vertical line). Lastly you can use `widget.NewToolbarSpacer()`
to create a flexible space between elements. This is most
useful to right align the toolbar items that are listed
after the spacer.

A toolbar should always be at the top of the content area
so it's normal to add it to a `fyne.Container` using the
`layout.BorderLayout` to align it above other content.
