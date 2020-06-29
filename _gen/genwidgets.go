package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type drawItem struct {
	name string
	obj fyne.CanvasObject
}

var (
	imgDir string
)

func makeDrawList() []drawItem {
	prop := canvas.NewRectangle(color.Transparent)
	prop.SetMinSize(fyne.NewSize(100, 0))
	se := widget.NewSelectEntry([]string{"1", "2"})
	se.SetPlaceHolder("Select one or type")
	return []drawItem{
		{"accordioncontainer", widget.NewAccordionContainer(
			&widget.AccordionItem{Title: "A", Detail: widget.NewLabel("Shown item"), Open: true},
			widget.NewAccordionItem("B", widget.NewLabel("Hidden")))},
		{"button", widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {})},
		{"check", &widget.Check{Text: "Check", Checked: true}},
		{"entry", &widget.Entry{PlaceHolder: "Entry"}},
		{"form", &widget.Form{Items:
		[]*widget.FormItem{
			{Text: "Username", Widget: widget.NewEntry()},
			{Text: "Password", Widget: widget.NewPasswordEntry()}},
			OnSubmit: func() {}, OnCancel: func() {}}},
		{"group", widget.NewGroup("Group", prop)},
		{"hyperlink", widget.NewHyperlink("fyne.io", nil)},
		{"icon", widget.NewIcon(theme.ContentPasteIcon())},
		{"label", widget.NewLabel("Text label")},
		{"password", &widget.Entry{PlaceHolder: "Password", Password: true}},
		{"progress", &widget.ProgressBar{Value: 0.74}},
		{"progressinf", widget.NewProgressBarInfinite()},
		{"radio", widget.NewRadio([]string{"Item 1"}, func(string) {})},
		{"scrollcontainer", widget.NewScrollContainer(widget.NewLabel("Scroll"))},
		{"select", widget.NewSelect([]string{"1", "2"}, func(string) {})},
		{"selectentry", se},
		{"slider", widget.NewSlider(-5, 25)},
		{"splitcontainer", widget.NewHSplitContainer(widget.NewLabel("Line1\nLine2"),
			widget.NewVSplitContainer(widget.NewLabel("Top"), widget.NewLabel("Bottom")))},
		{"tabcontainer", widget.NewTabContainer(
			widget.NewTabItem("Tab1", canvas.NewRectangle(color.Transparent)),
			widget.NewTabItem("Tab2", canvas.NewRectangle(color.Transparent)))},
		{"textgrid", makeTextGrid()},
		{"toolbar", widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() {}),
			widget.NewToolbarSeparator(),
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		)},
	}
}

func makeTextGrid() *widget.TextGrid {
	grid := widget.NewTextGridFromString("TextGrid\n  Content  ")
	grid.SetStyleRange(0, 4, 0, 7,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})
	grid.Rows[1].Style = &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}}

	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true

	return grid
}

func draw(obj fyne.CanvasObject, name string, c fyne.Canvas, themeName string) {
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
	c.SetContent(obj)
	img := c.Capture()
	err = png.Encode(file, img)
	if err != nil {
		fyne.LogError("Unable to write image", err)
	}
}

func main() {
	w := test.NewWindow(nil)
	c := w.Canvas()

	pwd, _ := os.Getwd()
	imgDir = filepath.Join(pwd, "images", "widgets")

	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	for _, item := range makeDrawList() {
		draw(item.obj, item.name, c, "light")
	}

	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	for _, item := range makeDrawList() {
		draw(item.obj, item.name, c, "dark")
	}
}

