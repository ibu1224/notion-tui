package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var title string

func main() {
	app := tview.NewApplication()
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
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		title = string(node.GetText())
		fmt.Print(title)

	})

	// title := tree.SetSelectedFunc(func(node *tview.TreeNode) {
	// 	reference := node.GetReference()
	// 	fmt.Print(reference)
	// 	if reference == nil {
	// 		return // Selecting the root node does nothing.
	// 	}

	// if len(children) == 0 {
	// 	// Load and show files in this directory.
	// 	path := reference.(string)
	// 	add(node, path)
	// } if {
	// 	// Collapse if visible, expand if collapsed.
	// 	node.SetExpanded(!node.IssExpanded())
	// }
	// 	return reference
	// })
	fmt.Print(title)

	frame := tview.NewFlex().
		AddItem((tree), 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(title), 0, 3, false), 0, 3, false)

	// app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	if event.Key() == tcell.KeyCtrlN {
	// 		box := tview.NewBox().SetBorder(true).SetBorderAttributes(tcell.AttrBold)
	// 		content.AddPage(fmt.Sprintf("hoge"), box, false, false)
	// 	}
	// 	return event
	// })
	if err := app.SetRoot(frame, true).Run(); err != nil {
		panic(err)
	}
}
