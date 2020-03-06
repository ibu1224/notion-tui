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
	// slides := []Slide{
	// 	Tree,
	// }
	// pages := tview.NewPages()
	// Create the main layout.
	title,content := Tree()
	fmt.Print(title)
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(content, 0, 1, true)
		// AddItem(info, 1, 1, false)
	
	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			fmt.Println("hoge")
		} else if event.Key() == tcell.KeyCtrlP {
			fmt.Println("hoge")
		}
		return event
	})

		// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
			panic(err)
	}
}