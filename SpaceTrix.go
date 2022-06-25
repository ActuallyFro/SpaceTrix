package main

import (
	"log" //stdout Debugging
	"math"

	"image/color"

	"fyne.io/fyne/v2/canvas" // https://developer.fyne.io/container/grid
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	newFyneApp := app.New()
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	totalBoardObjects := 1024
	totalBoardObjectsInRow := int(math.Sqrt(float64(totalBoardObjects)))

	centerCell := totalBoardObjects/2 + totalBoardObjectsInRow/2

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	// grid := container.New(layout.NewGridLayout(8), widget.NewLabel("Content"))
	grid := container.New(layout.NewGridLayout(totalBoardObjectsInRow))

	// 1: var textObjects []*canvas.Text
	for index := 0; index < totalBoardObjects; index++ {
		// 1: textObjects = append(textObjects, canvas.NewText(fmt.Sprint(index), color.White))
		// 2: grid.Add(canvas.NewText(fmt.Sprint(index), color.White))
		// val := "[_]" <-- EMPTY CELL
		val := "[-]"
		if index == centerCell {
			val = "[x]"
		}

		// text := canvas.NewText(fmt.Sprint(index), color.White)
		text := canvas.NewText(val, color.White)
		text.Alignment = fyne.TextAlignCenter
		// text.TextStyle = fyne.TextStyle{Italic: true}

		grid.Add(text)

		// grid.Add(canvas.NewText(fmt.Sprint(index), color.White))

	}

	ToolbarAndContent := container.NewBorder(toolbar, nil, nil, nil, grid)

	mainSpaceTrixWindow.SetContent(ToolbarAndContent)

	// w.SetContent(widget.NewLabel("Hello World!"))
	// w.ShowAndRun()

	// Close the App when Escape key is pressed
	mainSpaceTrixWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			newFyneApp.Quit()
		}
	})

	mainSpaceTrixWindow.ShowAndRun()
}
