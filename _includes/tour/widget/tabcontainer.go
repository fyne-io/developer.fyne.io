package main

import (
	"fyne.io/fyne/app"
	//"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		widget.NewTabItem("Tab 2", widget.NewLabel("World!")))

	//widget.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab"))

	tabs.SetTabLocation(widget.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
