---
layout: tour

title: List
order: 8

---

The `List` widget is one of the toolkit's collection widgets. 
These widgets are designed to help build really performant
interfaces when lots of data is being presented. 
You can also see the `Table` and `Tree` widgets which have a 
similar API. Because of this design they are a little more
complicated to use.

The list uses callback functions to ask for data when it is required.
There are 3 main callbacks, `Length`, `CreateItem` and `UpdateItem`. The Length callback (passed first) is the simplest,
it returns how many items are in the data to be presented. The
second two relate to templates - how graphical elements are
created, cached and re-used.

The `CreateItem` callback returns a new template object. This
will be re-used with real data when the widget is presented.
The `MinSize` of this object will influence the `List` minimum size.
Lastly `UpdateItem` is called to apply an item of data to a 
cached template. Use this to set the content ready for display.

If you want to read more we will move on to [data binding](/tour/binding/).
