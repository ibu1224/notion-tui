package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// The application.
var app = tview.NewApplication()

func main() {
	content := Tree()
	// Shortcuts to navigate the slides.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(content, 0, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			fmt.Println("test")
		} else if event.Key() == tcell.KeyCtrlP {
			fmt.Println("hoge")
		} else if event.Key() == tcell.KeyEsc {
			app.SetFocus(layout)
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
