package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var forcusId = ""

func Tree() (content tview.Primitive) {

	root := tview.NewTreeNode(".").SetColor(tcell.ColorLawnGreen)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	pages := getPages()

	add := func(target *tview.TreeNode, path string) {
		for _, page := range pages {
			title := page.(map[string]interface{})["title"].(string)
			node := tview.NewTreeNode(title).SetColor(tcell.ColorDarkSlateGray).SetReference(title)
			target.AddChild(node)
		}
	}
	add(root, ".")

	pageBlock := tview.NewTextView().SetWrap(false).SetDynamicColors(true)
	pageBlock.SetBorderPadding(1, 1, 2, 0)
	pageBlock.SetBorder(true)

	inputField := tview.NewInputField().SetLabel("Message:").SetFieldWidth(100)

	inputField.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			createPageContent(forcusId, inputField.GetText())
			fmt.Println(strings.TrimSpace(inputField.GetText()))
		}
	})
	// form := tview.NewForm().
	// 	AddInputField("Message:", "", 256, nil, nil).
	// 	AddButton("Save", func() { fmt.Print("button") })

	// fmt.Println(form.GetFormItem(0).GetText())
	// fmt.Println(reflect.TypeOf(form.GetFormItem(0).GetLabel()))

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		for _, page := range pages {
			title := page.(map[string]interface{})["title"].(string)
			if title == node.GetText() {
				pageBlock.Clear()
				forcusId = page.(map[string]interface{})["id"].(string)
				block := getPageContent(forcusId)
				for _, m := range block {
					title := m.(map[string]interface{})["title"].(string)
					fmt.Fprintln(pageBlock, title)
				}
				pageBlock.SetTitle(node.GetText())
			}
		}
		app.SetFocus(inputField)
	})

	return tview.NewFlex().AddItem((tree), 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(pageBlock, 0, 3, false).
			AddItem(inputField, 5, 1, false), 0, 4, false)
}
