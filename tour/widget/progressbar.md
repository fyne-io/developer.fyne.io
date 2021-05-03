---
layout: tour

title: ProgressBar
order: 6

---

The progress bar widget has two forms, the standard progress bar
shows the user which `Value` has been reached, from `Min` to
`Max`. The default min is `0.0` and the max defaults to `1.0`.
To use the default values just call `widget.NewProgressBar()`.
After creating you can set the `Value` field.

To set up a custom range you can set `Min` and `Max` fields
manually. The label will always show the percentage completion.

The other form of progress widget is the infinite progress bar.
This version simply shows that some activity is ongoing by
moving a segment of the bar from left to right and back again.
You create this using `widget.NewProgressBarInfinite()` and
it will start animating as soon as it is shown.
