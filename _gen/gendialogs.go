package main

import (
	"image/png"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type drawItem struct {
	name string
	dia  func(fyne.Window)
}

var (
	imgDir string
)

func makeDrawList() []drawItem {
	return []drawItem{
		{"information", func(w fyne.Window) { dialog.ShowInformation("Some Information", "This is a thing to know", w) }},
		{"confirm", func(w fyne.Window) { dialog.ShowConfirm("Please Confirm", "Are you sure..?", func(bool) {}, w) }},
		{"progress", func(w fyne.Window) {
			d := dialog.NewProgress("Please Wait", "Doing something...", w)
			d.SetValue(0.45)
			d.Show()
		}},
		{"progressinf", func(w fyne.Window) {
			d := dialog.NewProgressInfinite("Please Wait", "Doing something...", w)
			time.Sleep(time.Millisecond * 800)
			d.Show()
		}},
		{"entry", func(w fyne.Window) { dialog.ShowEntryDialog("Entry Input", "Enter a string...", func(string) {}, w) }},
		{"fileopen", func(w fyne.Window) { dialog.ShowFileOpen(func(fyne.URIReadCloser, error) {}, w) }},
		{"custom", func(w fyne.Window) {
			text := widget.NewTextGrid()
			text.SetText("Custom content")
			text.SetStyleRange(0, 7, 0, 14, widget.TextGridStyleWhitespace)
			dialog.ShowCustom("Custom Dialog", "Cancel", text, w)
		}},
	}
}

func draw(scene func(fyne.Window), name string, w fyne.Window, c test.WindowlessCanvas, themeName string) {
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
	for _, over := range c.Overlays().List() {
		c.Overlays().Remove(over)
	}

	c.SetContent(canvas.NewRectangle(theme.BackgroundColor()))
	scene(w)
	if name == "fileopen" {
		w.Resize(fyne.NewSize(420, 320))
		c.Overlays().Top().Resize(fyne.NewSize(412, 312))
	} else {
		w.Resize(fyne.NewSize(280, 160))
		c.Overlays().Top().Resize(fyne.NewSize(272, 152))
	}

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
	imgDir = filepath.Join(pwd, "images", "dialogs")

	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	for _, item := range makeDrawList() {
		draw(item.dia, item.name, w, c, "light")
	}

	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	for _, item := range makeDrawList() {
		draw(item.dia, item.name, w, c, "dark")
	}
}
