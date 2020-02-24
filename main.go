package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	root := tview.NewTreeNode(".").SetColor(tcell.ColorLawnGreen)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		// files, err := ioutil.ReadDir(path)
		// if err != nil {
		// panic(err)
		// }
		test := [3]string{"abc", "hoge", "pup"}
		for _, file := range test {
			node := tview.NewTreeNode(file).SetColor(tcell.Color100).
				SetReference(file)
			// SetSelectable(file)
			// if file.IsDir() {
			// 	node.SetColor(tcell.ColorRebeccaPurple)
			// }
			target.AddChild(node)
		}
	}

	add(root, ".")
	frame := tview.NewFlex().
		AddItem(tree, 0, 1, true).
		// AddItem(tview.NewFlex.SetBorder(true).SetBorderColor(tcell.ColorLawnGreen).SetTitle("treeView").SetTitleColor(tcell.Color100), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Main content"), 0, 3, false), 0, 3, false)

	if err := app.SetRoot(frame, true).Run(); err != nil {
		panic(err)
	}
}
