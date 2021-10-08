---
layout: page
title: List of struct and ListWithData widget
---


Sometime, you need to display a list of complex item, more complex than a list of string, but you also want to rely on Fyne's databinding. It can be the case when you want to display a list of contact, or books for example.


## Defining the binding


Well first thing first, let's start by defining our binding in our app : 

```go
var ContactListBinding binding.UntypedList = binding.NewUntypedList()
```

In Fyne, the only way to bind with a slice of struct is by using a `binding.UntypedList`.

## Creating the list

Fyne comes with a long list of widget, but the one we want to use is ListWithData. The constructor takes three parameters : 

- A databinding
- A function that create the shape of our list element
- A function to update our list element with the data


```go
widget.NewListWithData(
		ContactListBinding,
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel(""), widget.NewLabel(""), widget.NewLabel(""), widget.NewLabel(""))
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			v, _ := di.(binding.Untyped).Get()
			fn := binding.NewString()
			ln := binding.NewString()
			pn := binding.NewString()
			em := binding.NewString()

			fn.Set(v.(contact.Contact).FirstName)
			ln.Set(v.(contact.Contact).LastName)
			pn.Set(v.(contact.Contact).PhoneNumber)
			em.Set(v.(contact.Contact).Email)
			co.(*fyne.Container).Objects[0].(*widget.Label).Bind(fn)
			co.(*fyne.Container).Objects[1].(*widget.Label).Bind(ln)
			co.(*fyne.Container).Objects[2].(*widget.Label).Bind(pn)
			co.(*fyne.Container).Objects[3].(*widget.Label).Bind(em)
		},
	)
```


Ok that makes a lot of code, but let's get through this together.


Here we create a new ListWithData widget, we pass the binding we created earlier as first argument. 

Then we create our list item design with the second argument : a function that return a CanvasObject. 

Finally, we get the current dataitem in our list (the `v` variable) and pass its properties to different string bindings. Those string bindings are then bound to the label we created earlier.


## Set new values


Our list is ready, but we didn't gave her any value yet. Since we created a `binding.UntypedList` we need to set an untyped list of data, hence an slice of empty interface `[]interface{}`. For this, we need to convert our slice of struct : 


```go
	contacts := []contact.Contact{
		{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "+1-555-2194-532",
			Email:       "john.doe@email.com",
		},
		{
			FirstName:   "Jane",
			LastName:    "Doe",
			PhoneNumber: "+1-555-6064-123",
			Email:       "jane.doe@email.com",
		},
	}

	inter := make([]interface{}, len(contacts))

	for i, v := range contacts {
		inter[i] = v
	}

	ContactListBinding.Set(inter)
```

