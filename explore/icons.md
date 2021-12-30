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
  <li class="icon-item">{{ svg | split: 'svg11.dtd">' | last }}<figure><figcaption>{{ asset_goname }}</figcaption></figure></li>
{% endif %}
{% endfor %}
</ul>

<script>
    if (isDarkPreference()) {
        document.body.classList.add('dark-theme')
    } else {
        document.body.classList.add('light-theme')
    }
</script>
<style>
.theme-icon-list {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-evenly;
}
.icon-item {
    list-style: none;
    padding: 1rem;
    text-align: center;
}
.theme-icon-list > .icon-item svg {
    transform: scale(2);
}
.dark-theme .theme-icon-list > .icon-item svg {
    fill: var(white); /* hack cascade fallback for now */
    fill: var(--foreground-color); /* Replace this with hex codes or define CSS vars */
}
</style>
