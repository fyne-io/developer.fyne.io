---
layout: tour

title: Border Layout
order: 4

---

The border layout is possibly the most widely used to construct user
interfaces as it allows the positioning of items around a central element
which will expand to fill the space. To create a border layout you need
to pass the `fyne.CanvasObject`s that should be positioned in a border
position to the layout (as well as the container as usual). This
syntax is a little different to the other layouts but is basically just
`layout.NewBorderLayout(top, bottom, left, right)` as illustrated in
the example to the right.

Any items passed to the container that do not appear in specific border
locations will be positioned to the central area and will expand to
fill the space available. You can also pass `nil` to border parameters
that you wish to leave empty.

Note that all items in the center will expand to fill the space (as if
they were in a [`layout.MaxLayout`](maxlayout.html) container). To manage
the area yourself you can create a new `fyne.Container` and use any
layout you wish.
