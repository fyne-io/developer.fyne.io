---
layout: tour

title: Label
order: 1

permalink: /tour/widget/

redirect_from:
- /tour/widget/label

---

Widgets are the main components of a Fyne application GUI, they can be
used in an place that a basic `fyne.CanvasObject` can. They manage user
interactions but also will always match the current theme.

The Label widget is the simplest of them - it presents text to the user.
Unlike `canvas.Text` it can handle some simple formatting (such as `\n`).
You can create a label by calling `widget.NewLabel("some text")`, the
result can be assigned to a variable or passed directly into a container.
