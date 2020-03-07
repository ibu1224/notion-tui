package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Slide func() (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

func main() {
	title, content := Tree()
	fmt.Println(title)
	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			fmt.Println("test")
		} else if event.Key() == tcell.KeyCtrlP {
			fmt.Println("hoge")
		}
		return event
	})

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(content, 0, 1, true)

	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
