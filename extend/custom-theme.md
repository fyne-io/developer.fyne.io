---
layout: page
title: Creating a Custom Theme

redirect_from:
 - /tutorial/custom-theme
---

Applications are able to load custom themes that can apply small changes to the standard theme or provide a completely unique look. A theme will need to implement the functions of `fyne.Theme` interface, which is defined as follows:

```go
type Theme interface {
	Color(ThemeColorName, ThemeVariant) color.Color
	Font(TextStyle) Resource
	Icon(ThemeIconName) Resource
	Size(ThemeSizeName) float32
}
```

To apply our theme changes we will first define a new type that that implements this interface.

### Define your theme

We start by defining a new type that will be our theme, a simple empty struct will do:

```go
type myTheme struct {}
```

It is a good idea to assert that we implement an interface so that
compile errors are closer to the defining type.

```go
var _ fyne.Theme = (*myTheme)(nil)
```

At this point you could see compile errors because we still need to 
implement the methods, we start with color.

#### Customising colors

The `Color` function defined in the `Theme` interface asks us to define a
named color and also provides a hint for the variant that the user desires (for example `theme.VariantLight` or `theme.VariantDark`). In our theme we will return a custom background color - using a different value for light and dark.

```go
func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}
```

You will see the last line here references the `theme.DefaultTheme()` to
look up standard values. This allows us to provide custom values and yet
fall back to the standard theme when we don't want to provide our own values.

Of course colors are simpler than resources, we look at that for custom icons.

#### Overriding default icons

Icons (and Fonts) use `fyne.Resource` as values instead of simple types like `int` (for size) or `color.Color` for colors. We can build our own
resource using `fyne.NewStaticResource`, or you could pass in a value
that was created using [resource embedding](https://developer.fyne.io/tutorial/bundle).

```go
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	if name == theme.IconNameHome {
		fyne.NewStaticResource("myHome", homeBytes)
	}
	
	return theme.DefaultTheme().Icon(name)
}
```

As above we return the default theme icon if we don't want to provide
a specific override.


### Load the theme

Before we can load the theme you will need to implement the `Size` and `Font` methods as well. You can just use these empty implementations if
you are happy to use the defaults.

```go
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
```

To set the theme for your app you will need to add the following line of code:

```go
app.Settings().SetTheme(&myTheme{})
```

With these changes you can apply your own style, to make small tweaks or
provide a completely custom looking application!

In the end, the resulting code could look something like this:
```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct{}

func (myTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x0}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0x26, G: 0x26, B: 0x26, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x42}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0x7f}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x99}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x66}
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

func (myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (myTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 2
	default:
		return theme.DefaultTheme().Size(s)
	}
}

func main() {
	a := app.New()

	a.Settings().SetTheme(&myTheme{})

	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	
	w.ShowAndRun()
}```

