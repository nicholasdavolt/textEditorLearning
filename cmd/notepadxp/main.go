package main

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//sets basic app and creates window with title
	a := app.New()
	w := a.NewWindow("TextEditor — (untitled)")
	w.Resize(fyne.NewSize(800, 600))
	//creates the editor, sets wordwrap
	editor := widget.NewMultiLineEntry()
	editor.Wrapping = fyne.TextWrapWord

	//create status bar
	status := widget.NewLabel("Ln 1, Col 1")

	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() {}),
		fyne.NewMenuItem("Open…", func() {}),
		fyne.NewMenuItem("Save", func() {}),
		fyne.NewMenuItem("Save As…", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Close", w.Close))

	editMenu := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Undo", func() {}),
		fyne.NewMenuItem("Redo", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Cut", func() {}),
		fyne.NewMenuItem("Copy", func() {}),
		fyne.NewMenuItem("Paste", func() {}),
		fyne.NewMenuItem("Delete", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Select All", func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Find…", func() {}),
		fyne.NewMenuItem("Replace…", func() {}),
		fyne.NewMenuItem("Go to Line…", func() {}))
	formatMenu := fyne.NewMenu("Format",
		fyne.NewMenuItem("Word Wrap", func() {}),
		fyne.NewMenuItem("Font…", func() {}))

	viewMenu := fyne.NewMenu("View",
		fyne.NewMenuItem("Show Status Bar", func() {}))

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("TextEditor Help", func() {}))
	menuItems := []*fyne.Menu{fileMenu, editMenu, formatMenu, viewMenu, helpMenu}
	mainMenu := fyne.NewMainMenu(menuItems...)
	w.SetMainMenu(mainMenu)
	minSizePlaceholder := canvas.NewImageFromImage(image.NewNRGBA(image.Rect(0, 0, 0, 0)))
	minSizePlaceholder.SetMinSize(fyne.NewSize(400, 300))

	content := container.NewBorder(nil, status, nil, nil, editor)
	stackContainer := container.NewStack(minSizePlaceholder, content)

	w.SetContent(stackContainer)
	w.Canvas().Focus(editor)
	w.ShowAndRun()

}
