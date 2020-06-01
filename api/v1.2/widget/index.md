---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
--
    import "fyne.io/fyne/widget"

Package widget defines the UI widgets within the Fyne toolkit

## Usage

#### func  DestroyRenderer

```go
func DestroyRenderer(wid fyne.Widget)
```
DestroyRenderer frees a render implementation for a widget. This is typically for internal use only. Deprecated: Access to widget renderers is being removed, render details should be private to a WidgetRenderer.

#### func  Refresh

```go
func Refresh(wid fyne.Widget)
```
Refresh instructs the containing canvas to refresh the specified widget. Deprecated: Call Widget.Refresh() instead.

#### func  Renderer

```go
func Renderer(wid fyne.Widget) fyne.WidgetRenderer
```
Renderer looks up the render implementation for a widget Deprecated: Access to widget renderers is being removed, render details should be private to a WidgetRenderer.

#### types

 * [BaseWidget](basewidget.html)
 * [Box](box.html)
 * [Button](button.html)
 * [ButtonStyle](buttonstyle.html)
 * [Check](check.html)
 * [DisableableWidget](disableablewidget.html)
 * [Entry](entry.html)
 * [Form](form.html)
 * [FormItem](formitem.html)
 * [Group](group.html)
 * [Hyperlink](hyperlink.html)
 * [Icon](icon.html)
 * [Label](label.html)
 * [Orientation](orientation.html)
 * [PopUp](popup.html)
 * [ProgressBar](progressbar.html)
 * [ProgressBarInfinite](progressbarinfinite.html)
 * [Radio](radio.html)
 * [ScrollContainer](scrollcontainer.html)
 * [ScrollDirection](scrolldirection.html)
 * [Select](select.html)
 * [Slider](slider.html)
 * [TabContainer](tabcontainer.html)
 * [TabItem](tabitem.html)
 * [TabLocation](tablocation.html)
 * [Toolbar](toolbar.html)
 * [ToolbarAction](toolbaraction.html)
 * [ToolbarItem](toolbaritem.html)
 * [ToolbarSeparator](toolbarseparator.html)
 * [ToolbarSpacer](toolbarspacer.html)
