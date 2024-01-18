---
redirect_to:
  - https://docs.fyne.io/explore/dialogs.md

layout: page
title: Dialog List
readonly: true

redirect_from:
- /started/dialogs
---
## Standard Dialogs

---
### Color

Allow users to pick a colour from a standard set (or any color in advanced mode).

{% include dialog.html name="color" %}

### Confirm

Ask for conformation of an action.

{% include dialog.html name="confirm" %}

### FileOpen

Present this to ask user to choose a file to use inside the app.
The actual dialog displayed will depend on the current operating system.

{% include dialog.html name="fileopen" %}

### Form

Get various input elements in a dialog, with validation.

{% include dialog.html name="form" %}

### Information

A simple way to present some information to the app user.

{% include dialog.html name="information" %}

### Custom

Present any content inside a dialog container.

{% include dialog.html name="custom" %}
