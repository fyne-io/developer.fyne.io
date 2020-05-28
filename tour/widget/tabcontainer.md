---
layout: tour

title: TabContainer
order: 8

---

The tab container widget is used to allow the user to switch
between different content panels. Tabs are either just text
or text and icon. It is recommended not to mix some tabs
having icons and some without. A tab container is created
using `widget.NewTabContainer(...)` and passing
`widget.TabItem` items (that can be created using
`widget.NewTabItem(...)`).

The tab container can be configured by setting the location
of tabs, one of `widget.TabLocationTop`, `widget.TabLocationBottom`,
`widget.TabLocationLeading` and `widget.TabLocationTrailing`.
The default location is top.

When loaded on a mobile device the tab location may be ignored.
In a portrait orientation a leading or trailing position will
be changed to bottom. When in landscape then top or bottom
positions will be moved to leading.
