package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

//Tree hoge hoge
func Tree() (title string, content tview.Primitive) {

	root := tview.NewTreeNode(".").SetColor(tcell.ColorLawnGreen)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		pages := getPages()
		for _, page := range pages {
			title := page.(map[string]interface{})["title"].(string)
			node := tview.NewTreeNode(title).SetColor(tcell.ColorDarkSlateGray).SetReference(title)
			target.AddChild(node)
		}
	}
	add(root, ".")

	box := tview.NewBox().SetBorder(true)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		box.SetTitle(node.GetText())
	})

	return "Tree", tview.NewFlex().AddItem((tree), 0, 1, true).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(box, 0, 3, false), 0, 3, false)
}