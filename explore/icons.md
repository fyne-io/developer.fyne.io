---
layout: page
title: Theme Icons

---

## Icons

Each of the following icons is available via the `theme` package as a function.
For example `theme.InfoIcon()`.

The icons are also available via their source icon name by using the `ThemeIconName` 
with the `safeIconLookup` function. For example `theme.SafeIconLookup(theme.IconNameInfo)`.

### List

<ul class="theme-icon-list">
{% for icon in site.data.icons.themedIcons %}
  {% capture svg %}{% include {{ icon.includePath }} %}{% endcapture %}
  <li class="icon-item" data-filepath="{{ icon.includePath }}" data-sourceIcon="{{ icon.sourceIcon }}" data-icon-theme-method="{{ icon.Name }}Icon()" data-icon-safeName="IconName{{ icon.Name }}" id="IconName{{ icon.Name }}"><figure>{{ svg | split: 'svg11.dtd">' | last }}<figcaption>{{ icon.Name }}Icon()</figcaption></figure></li>
{% endfor %}
</ul>
