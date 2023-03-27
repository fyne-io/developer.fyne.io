---
layout: page
title: Theme Icons

---
## Icons

Each of the following icons is available via the `theme` package as a function. 
For example `theme.InfoIcon()`.

The icons are also available via their source icon name by using the `ThemeIconName` 
with the `Icon` method on a struct implementing `fyne.Theme`. For example `theme.Icon(theme.IconNameInfo)`.

## List

<ul class="theme-icon-list">
{% for icon in site.data.icons.themedIcons %}
  {% capture svg %}{% include {{ icon.includePath }} %}{% endcapture %}
  <li class="icon-item" data-icon-theme-method="{{ icon.Name }}Icon()" data-icon-safeName="IconName{{ icon.Name }}" id="IconName{{ icon.Name }}"><figure>{{ svg | split: 'svg11.dtd">' | last }}<figcaption>{{ icon.Name }}Icon()</figcaption></figure></li>
{% endfor %}
</ul>

## Using other color sets

The source icon can be used to wrap a source icon for a particular theme color using the 
various public helper methods:

* `NewDisabledThemedResource`
* `NewErrorThemedResource`
* `NewInvertedThemedResource`
* `NewPrimaryThemedResource`

By default all icons are wrapping source icons using `NewThemedResource`; 
which uses the theme foreground color. All Icons are SVG `width="24"`, `height="24"`.
