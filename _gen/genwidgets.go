package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type drawItem struct {
	name string
	obj  fyne.CanvasObject
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
		{"accordion", widget.NewAccordion(
			&widget.AccordionItem{Title: "A", Detail: widget.NewLabel("Hidden")},
			widget.NewAccordionItem("B", widget.NewLabel("Shown item")),
			widget.NewAccordionItem("C", widget.NewLabel("2")))},
		{"button", widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {})},
		{"card", &widget.Card{Title: "Card Title", Subtitle: "Subtitle", Image: canvas.NewImageFromResource(theme.FyneLogo())}},
		{"check", &widget.Check{Text: "Check", Checked: true}},
		{"entry", &widget.Entry{PlaceHolder: "Entry"}},
		{"entry-invalid", makeInvalidEntry()},
		{"entry-valid", &widget.Entry{Validator: func(_ string) error { return nil }, Text: "Valid"}},
		{"form", &widget.Form{Items: []*widget.FormItem{
			{Text: "Username", Widget: widget.NewEntry()},
			{Text: "Password", Widget: widget.NewPasswordEntry()}},
			OnSubmit: func() {}, OnCancel: func() {}}},
		{"group", widget.NewGroup("Group", prop)},
		{"hyperlink", widget.NewHyperlink("fyne.io", nil)},
		{"icon", widget.NewIcon(theme.ContentPasteIcon())},
		{"label", widget.NewLabel("Text label")},
		{"list", makeList()},
		{"table", makeTable()},
		{"tree", makeTree()},
		{"password", &widget.Entry{PlaceHolder: "Password", Password: true}},
		{"popupmenu", makePopUpMenu()},
		{"progress", &widget.ProgressBar{Value: 0.74}},
		{"progressinf", widget.NewProgressBarInfinite()},
		{"radio", &widget.Radio{Options: []string{"Item 1", "Item 2"}, OnChanged: func(string) {}, Selected: "Item 1"}},
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

func makeInvalidEntry() *widget.Entry {
	e := widget.NewEntry()
	e.Validator = func(_ string) error { return fmt.Errorf("reason") }
	test.Type(e, "Invalid")
	e.FocusLost()
	return e
}

func makeList() *widget.List {
	l := widget.NewList(func() int { return 5 },
		func() fyne.CanvasObject {
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("TemplateItm"))
		},
		func(i int, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(fmt.Sprintf("List item %d", i+1))
		})

	return l
}

func makePopUpMenu() fyne.CanvasObject {
	m := widget.NewMenu(fyne.NewMenu("",
		fyne.NewMenuItem("Item 1", func() {}),
		fyne.NewMenuItem("Item 2", func() {})))

	m.Items[0].(desktop.Hoverable).MouseIn(nil)
	return m
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

func makeTable() *widget.Table {
	return widget.NewTable(
		func() (int, int) {
			return 3, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 0, 0")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			text := fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1)
			obj.(*widget.Label).SetText(text)
		})
}

func makeTree() *widget.Tree {
	data := make(map[string][]string)
	data[""] = []string{"Cars", "Trains"}
	data["Cars"] = []string{"Ford", "Tesla", "Jaguar"}
	data["Trains"] = []string{"Rocket", "TGV", "Eurotar"}

	t := widget.NewTreeWithStrings(data)
	t.OpenBranch("Trains")
	return t
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
	if name == "progressinf" {
		time.Sleep(time.Second)
	} else if name == "list" || name == "table" || name == "tree" || name == "accordion" {
		obj.Resize(fyne.NewSize(136, 120))
		c.(test.WindowlessCanvas).Resize(fyne.NewSize(136, 120))
		test.TapCanvas(c, fyne.NewPos(50, 50))
	}
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
