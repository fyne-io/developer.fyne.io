---
layout: tour

title: Line
order: 3

---

The `canvas.Line` object draws a line from the `Position1` (default
is top, left) to `Position2` (default is bottom, right).
You specify it's colour and can vary the stroke width which otherwise
defaults to `1`.

A line position can be manipulated using the `Position1` or `Position2`
fields or by using the `Move()` and `Resize()` functions. For example 
a 0 width area will show a vertical line whereas 0 height would be
horizontal.

Lines are typically used in a custom layout or controlled manually.
Unlike text they have no natural (minimum) size but can be
used to great effect in complex layouts.
