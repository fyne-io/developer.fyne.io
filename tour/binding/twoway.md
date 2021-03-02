---
layout: tour

title: Two-Way Binding
order: 3

permalink: /tour/binding/twoway

---

So far we have looked at data binding as a way of keeping
our user interface elements up to date. Far more common,
however, is the need to update valules from the UI
widgets and keep the data up to date everywhere.
Thankfully the bindings provided in Fyne are "two-way"
which means that values can be pushed into them as well
as read out. The change in data will be communicated to
all connected code without any additional code.

To see this in action we can update the last test app
to display a `Label` and an `Entry` that are bound to
the same value. By setting this up you can see that
editing the value through the entry will update the
text in the label as well. This is all possible without
calling refresh or referencing the widgets in our code.

By moving your app to use data binding you can stop
saving pointers to all your widgets. By instead 
capturing your data as a set of bound values your
user interface can be completely separate code.
Cleaner to read and easier to manage.

Next we will look at how to add [conversions](/tour/binding/conversion) in our data.
