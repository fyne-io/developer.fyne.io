---
layout: page
title: Theme Icons

---

## Icons

<ul class="theme-icon-list">
{% for file in site.static_files %}
{% if file.path contains "/images/theme/icons/" and file.extname == ".svg" %}
    {% capture asset_goname %}{{ file.basename | replace: "-", "" }}IconRes{% endcapture %}
    {% capture asset_filename %}theme/icons/{{ file.name }}{% endcapture %}
    {% capture svg %}{% include {{ asset_filename }} %}{% endcapture %}
  <li class="icon-item"><figure>{{ svg | split: 'svg11.dtd">' | last }}<figcaption>{{ asset_goname }}</figcaption></figure></li>
{% endif %}
{% endfor %}
</ul>
