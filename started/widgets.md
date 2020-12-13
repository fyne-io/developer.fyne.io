---
layout: page
title: Widget List

order: 30
---

## Standard Widgets (in `widget` package)

---

### Accordion

Accordion displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.

{% include widget.html name="accordion" %}

### Button

Button widget has a text label and icon, both are optional.

{% include widget.html name="button" %}

### Card

Card widget groups elements with a header and subheader, all are optional.

{% include widget.html name="card" %}

### Check

Check widget has a text label and a checked (or unchecked) icon.

{% include widget.html name="check" %}

### Entry

Entry widget allows simple text to be input when focused.

{% include widget.html name="entry" %}
{% include widget.html name="entry-valid" %}
{% include widget.html name="entry-invalid" %}

PasswordEntry widget hides text input and adds a button to display the text.

{% include widget.html name="password" %}

### FileIcon

FileIcon provides helpful standard icons for various types of file.
It displays the type of file as an indicator icon and shows the extension of the file type.

{% include widget.html name="fileicon" %}

### Form

Form widget is two column grid where each row has a label and a widget (usually an input). The last row of the grid will contain the appropriate form control buttons if any should be shown.

{% include widget.html name="form" %}

### Hyperlink

Hyperlink widget is a text component with appropriate padding and layout. When clicked, the URL opens in your default web browser.

{% include widget.html name="hyperlink" %}

### Icon

Icon widget is a basic image component that load's its resource to match the theme.

{% include widget.html name="icon" %}

### Label

Label widget is a label component with appropriate padding and layout.

{% include widget.html name="label" %}

### Progress bar

ProgressBar widget creates a horizontal panel that indicates progress.

{% include widget.html name="progress" %}

ProgressBarInfinite widget creates a horizontal panel that indicates waiting indefinitely An infinite progress bar loops 0% -> 100% repeatedly until Stop() is called.

{% include widget.html name="progressinf" %}

### Radio

Radio widget has a list of text labels and radio check icons next to each.

{% include widget.html name="radio" %}

### Select

Select widget has a list of options, with the current one shown, and triggers an event function when clicked.

{% include widget.html name="select" %}

### SelectEntry

Select entry widget adds an editable component to the select widget.
Users can select an option or enter their own value.

{% include widget.html name="selectentry" %}

### Separator

Separator widget shows a dividing line between other elements.

{% include widget.html name="separator" %}

### Slider

Slider if a widget that can slide between two fixed values.

{% include widget.html name="slider" %}

### TextGrid

TextGrid is a monospaced grid of characters. This is designed to be used by a text editor, code preview or terminal emulator.

{% include widget.html name="textgrid" %}

### Toolbar

Toolbar widget creates a horizontal list of tool buttons.

{% include widget.html name="toolbar" %}


## Collection Widgets (in `widget` package)

Collection widgets provide advanced caching functionality to provide high performance rendering of massive data. This does lead to a more complex constructor,
but is a good balance for the outcome it enables.
Each of these widgets uses a series of callbacks, the minimum set is defined by their constructor function, which includes the data size, the creation of template items that can be re-used and finally the function that applies data to a widget as it is about to be added to the display.

### List

List provides a high performance vertical scroll of many sub-items.

{% include widget.html name="list" %}

### Table

Table provides a high performance scrolled two dimensional display of many sub-items.

{% include widget.html name="table" %}

### Tree

Tree provides a high performance vertical scroll of items that can be expanded to reveal child elements..

{% include widget.html name="tree" %}


## Container Widgets (in `container` package)

Container widgets are like regular containers but they provide some additional functionality.

### AppTabs

AppTabs widget allows switching visible content from a list of TabItems. Each item is represented by a button at the top of the widget.

{% include widget.html name="apptabs" %}

### ScrollContainer

ScrollContainer defines a container that is smaller than the Content.

{% include widget.html name="scrollcontainer" %}

### SplitContainer

SplitContainer defines a container whose size is split between two children.

{% include widget.html name="splitcontainer" %}
