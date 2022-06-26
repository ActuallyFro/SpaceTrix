package main

import (
	"log" //stdout Debugging
	"math"

	"image/color"

	"fyne.io/fyne/v2/canvas" // https://developer.fyne.io/container/grid
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"github.com/actuallyfro/SpaceTrix/include"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type boardElement struct {
	value        string
	elementColor color.RGBA
}

func main() {

	newFyneApp := app.New()
	newFyneApp.Settings().SetTheme(theme.DarkTheme())
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	totalBoardObjects := 1024
	totalBoardObjectsInRow := int(math.Sqrt(float64(totalBoardObjects)))

	centerCell := totalBoardObjects/2 + totalBoardObjectsInRow/2

	//create array boardElement to hold the board, set default as '-'
	board := make([]boardElement, totalBoardObjects)
	for i := 0; i < totalBoardObjects; i++ {
		if i == centerCell {
			board[i].value = "[X]"
			board[i].elementColor = color.RGBA{0, 0, 255, 255}
		} else {
			board[i].value = "[-]"
			board[i].elementColor = color.RGBA{255, 255, 255, 255}

		}

	}

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { newFyneApp.Quit() }),
	)

	//https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1
	//https://blogvali.com/menu-items-fyne-gui-golang-tutorial-35/

	helpMenuAbout := fyne.NewMenuItem("About", func() {
		dialog.ShowCustom("About", "Close", container.NewVBox(
			widget.NewLabel("SpaceTrix - a WASD Adventure"),
			widget.NewLabel("Version: v0.0.1"),
			widget.NewLabel("Author: actuallyfro"),
		), mainSpaceTrixWindow)
	})

	clock := widget.NewLabel("")
	include.InitRecurringFunctionUpdateClock(clock)

	include.UpdateTime(clock)

	helpMenuSeeTime := fyne.NewMenuItem("See Time", func() {
		dialog.ShowCustom("Current Time", "Close", container.NewVBox(
			clock,
		), mainSpaceTrixWindow)
	})

	helpMenu := fyne.NewMenu("Help", helpMenuAbout, helpMenuSeeTime)

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

	grid := container.New(layout.NewGridLayout(totalBoardObjectsInRow))

	/*
	   want this ... but for the BOARD rendering
	   func UpdateTime(clock *widget.Label) {
	   	formatted := time.Now().Format("Time: 03:04:05")
	   	clock.SetText(formatted)
	   }

	   func InitRecurringFunctionUpdateClock(clock *widget.Label) {
	   	// w.SetContent(clock)
	   	go func() {
	   		for range time.Tick(time.Second) {
	   			UpdateTime(clock)
	   		}
	   	}()
	   }


	*/

	for index := 0; index < totalBoardObjects; index++ {
		// 1: textObjects = append(textObjects, canvas.NewText(fmt.Sprint(index), color.White))
		// 2: grid.Add(canvas.NewText(fmt.Sprint(index), color.White))
		// val := "[_]" <-- EMPTY CELL
		text := canvas.NewText(board[index].value, board[index].elementColor)
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
