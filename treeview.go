package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func treeView() tview {

	// newPrimitive := func(text string) tview.Primitive {
	// 	return tview.NewTextView().
	// 		SetTextAlign(tview.AlignCenter).
	// 		SetText(text)
	// }
	menu := newPrimitive("Menu")
	// main := newPrimitive("Main content")

	// grid := tview.NewGrid().
	// 	SetRows(3, 0, 3).
	// 	SetColumns(30, 0, 30).
	// 	SetBorders(true).
	// 	AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
	// 	AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).

	rootDir := "."
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorSlateGray)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		mydata := [3]string{"hoge", "fuga", "piyo"}
		for _, file := range mydata {
			node := tview.NewTreeNode(file)
			node.SetColor(tcell.ColorMediumAquamarine)
			target.AddChild(node)
		}
	}
	// }
	// Add the current directory to the root node.
	add(root, rootDir)
	return tree

	// if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
	// 	panic(err)
	// }
}
