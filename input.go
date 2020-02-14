package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetFieldWidth(256).
		SetAcceptanceFunc(tview.InputFieldString).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	fmt.Println(inputField.GetText())
	if err := app.SetRoot(inputField, true).Run(); err != nil {
		panic(err)
	}
}
