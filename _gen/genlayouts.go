package main

import (
	"image/png"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
)

type drawItem struct {
	name string
	lay  fyne.CanvasObject
}

var (
	imgDir string
)

func makeDrawList() []drawItem {
	bObjs := makeObjs()
	border := fyne.NewContainerWithLayout(layout.NewBorderLayout(bObjs[0], nil, bObjs[1], nil), bObjs...)

	smaller := makeObjs()
	for _, s := range smaller {
		s.(*canvas.Rectangle).SetMinSize(fyne.NewSize(s.MinSize().Width, 25))
	}
	top := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), smaller...)
	left := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), makeObjs()...)
	content := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(40, 40)), makeObjs()...)
	combined := fyne.NewContainerWithLayout(layout.NewBorderLayout(top, nil, left, nil), top, left, content)

	return []drawItem{
		{"hbox", fyne.NewContainerWithLayout(layout.NewHBoxLayout(), makeObjs()...)},
		{"vbox", fyne.NewContainerWithLayout(layout.NewVBoxLayout(), makeObjs()...)},
		{"border", border},
		{"center", fyne.NewContainerWithLayout(layout.NewCenterLayout(), makeObjs()...)},
		{"grid", fyne.NewContainerWithLayout(layout.NewGridLayout(2), makeObjs()...)},
		{"gridwrap", fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(75, 75)), makeObjs()...)},
		{"max", fyne.NewContainerWithLayout(layout.NewMaxLayout(), makeObjs()...)},
		{"combined", combined},
	}
}

func makeObjs() []fyne.CanvasObject {
	prop1 := canvas.NewRectangle(theme.ButtonColor())
	prop1.SetMinSize(fyne.NewSize(50, 50))
	prop2 := canvas.NewRectangle(theme.ButtonColor())
	prop2.SetMinSize(fyne.NewSize(25, 25))
	prop3 := canvas.NewRectangle(theme.ButtonColor())
	prop3.SetMinSize(fyne.NewSize(75, 75))
	return []fyne.CanvasObject{prop1, prop2, prop3}
}

func draw(scene fyne.CanvasObject, name string, c test.WindowlessCanvas, themeName string) {
	fileName := filepath.Join(imgDir, name+"-"+themeName+".png")
	file, err := os.Create(fileName)
	if err != nil {
		fyne.LogError("err", err)
		file, err = os.Open(fileName)
		if err != nil {
			fyne.LogError("Unable to open file for writing", err)
			return
		}
	}

	c.SetScale(2.0) // get HiDPI output so we can render nicely on fancy screens :)
	c.SetContent(scene)
	c.Resize(fyne.NewSize(250, 200))
	img := c.Capture()
	err = png.Encode(file, img)
	if err != nil {
		fyne.LogError("Unable to write image", err)
	}
}

func main() {
	w := test.NewWindow(nil)
	c := w.Canvas().(test.WindowlessCanvas)

	pwd, _ := os.Getwd()
	imgDir = filepath.Join(pwd, "images", "layouts")

	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	for _, item := range makeDrawList() {
		draw(item.lay, item.name, c, "light")
	}

	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	for _, item := range makeDrawList() {
		draw(item.lay, item.name, c, "dark")
	}
}
