// Demo code for the List primitive.
package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	list := tview.NewList()
	list.AddItem(NewItem("one", "first"))
	list.AddItem(NewItem("two", "second"))
	list.AddItem(NewItem("three", "third"))
	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

type Item struct {
	txt string
	sec string
}

func NewItem(txt string, sec string) *Item {
	return &Item{txt: txt, sec: sec}
}

func (i *Item) GetText() string {
	return i.txt
}

func (i *Item) GetSecondaryText() string {
	return i.sec
}

func (i *Item) GetSelected() func() {
	return func() {

	}
}
