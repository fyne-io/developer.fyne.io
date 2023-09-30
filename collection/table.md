---
title: Table

redirect_from:
  - /tour/widget/table
  - /widget/table
--- 

The `Table` collection widget is like the [List](/collection/list) widget (another of the toolkit's collection widgets) with a two-dimensional index.
Like `List` this is designed to help build really performant
interfaces when lots of data is being presented.
Because of this the widget is not created with all the data embedded, but instead calls out to the data source when needed.

The `Table` uses callback functions to ask for data when it is required.
There are 3 main callbacks, `Length`, `CreateCell` and `UpdateCell`. The Length callback (passed first) is the simplest,
it returns how many items are in the data to be presented, the two ints it returns represent the row and colum count.
The other two relate to the content templates.

The `CreateCell` callback returns a new template object, just like list.
The difference being that `MinSize` will define the standard size of each cell, and the minimum size of the table (it shows at least one cell).
As previously the `UpdateCell` is called to apply data to a cell template. The index passed in is the same `(row, col)` int pair.

```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
```

The following code expands on this earlier example. It demonstrates how to enable and customize row and column headers.
```go
package main

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)

const (
	rowNum = 100
	colNum = 5
)

func main() {
	a := app.New()
	w := a.NewWindow("Headed Table")
	t := widget.NewTableWithHeaders(dataLength, dataCreate, dataUpdate)
	t.CreateHeader = headerCreate
	t.UpdateHeader = headerUpdate // Updating headers is mandatory when they have been created
	t.SetColumnWidth(0, 155)
	w.SetContent(t)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}

type ActiveCell struct {
	widget.Label
	OnTapped func()
}

func newActiveCell(label string) *ActiveCell {
	c := &ActiveCell{}
	c.ExtendBaseWidget(c)
	c.Label.SetText(label)
	return c
}

func (h *ActiveCell) Tapped(_ *fyne.PointEvent) {
	if h.OnTapped != nil {
		h.OnTapped()
	}
}

func (h *ActiveCell) TappedSecondary(_ *fyne.PointEvent) {
}

func dataLength() (int, int) {
	return rowNum, colNum
}

func dataCreate() fyne.CanvasObject {
	return newActiveCell("Cell 000, 000")
}

func dataUpdate(id widget.TableCellID, o fyne.CanvasObject) {
	label := o.(*ActiveCell)
	switch id.Col {
	case 0:
		label.SetText("A looooooooonger cell")
	case 1, 2, 3:
		label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
	case 4:
		label.SetText(fmt.Sprintf("Tail cell %d", id.Row+1))
	}
	label.OnTapped = func() {
		log.Printf("Cell tapped %d, %d\n", id.Row, id.Col)
	}
}

type ActiveHeader struct {
	widget.Label
	OnTapped func()
}

func newActiveHeader(label string) *ActiveHeader {
	h := &ActiveHeader{}
	h.ExtendBaseWidget(h)
	h.SetText(label)
	return h
}

func (h *ActiveHeader) Tapped(_ *fyne.PointEvent) {
	if h.OnTapped != nil {
		h.OnTapped()
	}
}

func (h *ActiveHeader) TappedSecondary(_ *fyne.PointEvent) {
}

func headerCreate() fyne.CanvasObject {
	return newActiveHeader("000")
}

func headerUpdate(id widget.TableCellID, o fyne.CanvasObject) {
	header := o.(*ActiveHeader)
	header.TextStyle.Bold = true
	switch id.Col {
	case -1:
		header.SetText(strconv.Itoa(id.Row + 1))
	case 0:
		header.SetText("A longer header")
	case 1, 2, 3:
		header.SetText(fmt.Sprintf("Header %d", id.Col+1))
	case 4:
		header.SetText("Nonsortable")
	}

	header.OnTapped = func() {
		log.Printf("Header %d tapped\n", id.Col)
	}
}

```
