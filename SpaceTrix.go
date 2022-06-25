package main

import (
	"log" //stdout Debugging
	"math"
	"time"

	"image/color"

	"fyne.io/fyne/v2/canvas" // https://developer.fyne.io/container/grid
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	newFyneApp := app.New()
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	totalBoardObjects := 1024
	totalBoardObjectsInRow := int(math.Sqrt(float64(totalBoardObjects)))

	centerCell := totalBoardObjects/2 + totalBoardObjectsInRow/2

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { newFyneApp.Quit() }),
	)

	//https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1
	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("SpaceTrix - a WASD Adventure"),
				widget.NewLabel("Version: v0.0.1"),
				widget.NewLabel("Author: actuallyfro"),
			), mainSpaceTrixWindow)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	mainSpaceTrixWindow.SetMainMenu(mainMenu)

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

	clock := widget.NewLabel("")
	updateTime(clock)

	// w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	grid := container.New(layout.NewGridLayout(totalBoardObjectsInRow))

	for index := 0; index < totalBoardObjects; index++ {
		// 1: textObjects = append(textObjects, canvas.NewText(fmt.Sprint(index), color.White))
		// 2: grid.Add(canvas.NewText(fmt.Sprint(index), color.White))
		// val := "[_]" <-- EMPTY CELL
		val := "[-]"
		if index == centerCell {
			val = "[x]"
		}

		text := canvas.NewText(val, color.White)
		text.Alignment = fyne.TextAlignCenter

		grid.Add(text)

	}

	ToolbarAndContent := container.NewBorder(toolbar, nil, nil, nil, grid)

	mainSpaceTrixWindow.SetContent(ToolbarAndContent)

	mainSpaceTrixWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			newFyneApp.Quit()
		}
	})

	mainSpaceTrixWindow.ShowAndRun()
}
