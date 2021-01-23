---
layout: tour

title: AppTabs
order: 1

permalink: /tour/container/

redirect_from:
- /tour/container/apptabs

---

The AppTabs container is used to allow the user to switch
between different content panels. Tabs are either just text
or text and an icon. It is recommended not to mix some tabs
having icons and some without. A tab container is created
using `container.NewAppTabs(...)` and passing
`container.TabItem` items (that can be created using
`container.NewTabItem(...)`).

The tab container can be configured by setting the location
of tabs, one of `container.TabLocationTop`, `container.TabLocationBottom`,
`container.TabLocationLeading` and `container.TabLocationTrailing`.
The default location is top.

When loaded on a mobile device the tab location may be ignored.
In a portrait orientation a leading or trailing position will
be changed to bottom. When in landscape orientation, the top or bottom
positions will be moved to leading.
