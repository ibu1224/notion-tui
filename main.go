package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// The application.
var app = tview.NewApplication()

var layout = tview.NewFlex().
	SetDirection(tview.FlexRow).
	AddItem(Tree(), 0, 1, true)

func main() {
	// Shortcuts to navigate the slides.

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
