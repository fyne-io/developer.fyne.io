---
layout: tour

title: Text
order: 2

---

`canvas.Text` is used for all text rendering within Fyne.
It is created by specifying the text and colour for the text.
Text is rendered using the default font, specified by the current theme.

The text object allows certain configuration like the `Alignment`
and `TextStyle` field. as illustrated in the example here.
To use a monospaced font instead you can specify
`fyne.TextStyle{Monospace: true}`.

It is possible to use an alternative font by specifying a `FYNE_FONT`
environment variable. Use this to set a `.ttf` file to use instead of
the one provided in the Fyne toolkit or the current theme.
