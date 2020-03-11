package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

//Tree hoge hoge
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

	form := tview.NewForm().
		AddInputField("Message:", "", 256, nil, nil).
		AddButton("Save", nil)

	fmt.Print(form.GetFormItem(0))

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		for _, page := range pages {
			title := page.(map[string]interface{})["title"].(string)
			if title == node.GetText() {
				pageBlock.Clear()
				id := page.(map[string]interface{})["id"].(string)
				block := getPageContent(id)
				for _, m := range block {
					title := m.(map[string]interface{})["title"].(string)
					fmt.Fprintln(pageBlock, title)
				}
				pageBlock.SetTitle(node.GetText())
			}
		}
		app.SetFocus(form)
	})

	return tview.NewFlex().AddItem((tree), 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(pageBlock, 0, 3, false).
			AddItem(form, 5, 1, false), 0, 4, false)
	// AddItem(input, 3, 5, false)
}
