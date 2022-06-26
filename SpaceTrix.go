package main

import (
	"log" //stdout Debugging
	"math"
	"time"

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

func UpdateGrid(passedBoardTotalObjects int, passedBoardTotalRowObjects int, passedBoard []boardElement) *fyne.Container {

	grid := container.New(layout.NewGridLayout(passedBoardTotalRowObjects))

	for index := 0; index < passedBoardTotalObjects; index++ {
		// 1: textObjects = append(textObjects, canvas.NewText(fmt.Sprint(index), color.White))
		// 2: grid.Add(canvas.NewText(fmt.Sprint(index), color.White))
		// val := "[_]" <-- EMPTY CELL
		text := canvas.NewText(passedBoard[index].value, passedBoard[index].elementColor)
		text.Alignment = fyne.TextAlignCenter

		grid.Add(text)

	}
	return grid
}

//    want this ... but for the BOARD rendering
func UpdateBoard(passedWindow fyne.Window, passedToolbar *widget.Toolbar, passedContent *fyne.Container) {
	ToolbarAndContent := container.NewBorder(passedToolbar, nil, nil, nil, passedContent)
	passedWindow.SetContent(ToolbarAndContent)

	// log.Println("[DEBUG] Updating board...")

}

func InitRecurringFunctionUpdateBoard(passedWindow fyne.Window, passedToolbar *widget.Toolbar, passedContent *fyne.Container) {
	// w.SetContent(clock)
	go func() {
		//updat every 100ms
		for range time.Tick(time.Millisecond * 500) {
			UpdateBoard(passedWindow, passedToolbar, passedContent)
		}
	}()
}

func main() {

	newFyneApp := app.New()
	newFyneApp.Settings().SetTheme(theme.DarkTheme())
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	totalBoardObjects := 1024
	totalBoardObjectsInRow := int(math.Sqrt(float64(totalBoardObjects)))
	totalBoardObjectsInRowLessOne := totalBoardObjectsInRow - 1

	centerCell := totalBoardObjects/2 + totalBoardObjectsInRow/2

	currentPos := centerCell
	hasBoardBeenInitialized := false

	var board []boardElement
	var grid *fyne.Container

	var helpMenuAbout *fyne.MenuItem
	var clock *widget.Label
	var helpMenuSeeTime *fyne.MenuItem
	var helpMenu *fyne.Menu
	var fileMenu *fyne.Menu
	var mainMenu *fyne.MainMenu

	var toolbar *widget.Toolbar

	if !hasBoardBeenInitialized {
		//create array boardElement to hold the board, set default as '-'
		board = make([]boardElement, totalBoardObjects)
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
		fileMenu = fyne.NewMenu("File",
			fyne.NewMenuItem("Quit", func() { newFyneApp.Quit() }),
		)

		//https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1
		//https://blogvali.com/menu-items-fyne-gui-golang-tutorial-35/

		helpMenuAbout = fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("SpaceTrix - a WASD Adventure"),
				widget.NewLabel("Version: v0.0.1"),
				widget.NewLabel("Author: actuallyfro"),
			), mainSpaceTrixWindow)
		})

		clock = widget.NewLabel("")
		include.InitRecurringFunctionUpdateClock(clock)

		include.UpdateTime(clock)

		helpMenuSeeTime = fyne.NewMenuItem("See Time", func() {
			dialog.ShowCustom("Current Time", "Close", container.NewVBox(
				clock,
			), mainSpaceTrixWindow)
		})

		helpMenu = fyne.NewMenu("Help", helpMenuAbout, helpMenuSeeTime)

		mainMenu = fyne.NewMainMenu(
			fileMenu,
			helpMenu,
		)
		mainSpaceTrixWindow.SetMainMenu(mainMenu)

		toolbar = widget.NewToolbar(
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

		grid = UpdateGrid(totalBoardObjects, totalBoardObjectsInRow, board)

		// ToolbarAndContent := container.NewBorder(toolbar, nil, nil, nil, grid)
		// mainSpaceTrixWindow.SetContent(ToolbarAndContent)
		InitRecurringFunctionUpdateBoard(mainSpaceTrixWindow, toolbar, grid)

		hasBoardBeenInitialized = true
	}

	mainSpaceTrixWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		posUpdate := false
		oldPos := currentPos

		if keyEvent.Name == fyne.KeyEscape {
			newFyneApp.Quit()

		} else if keyEvent.Name == fyne.KeyUp {
			if currentPos > totalBoardObjectsInRow {
				currentPos -= totalBoardObjectsInRow
				posUpdate = true
			}

			// log.Println("[DEBUG] Up Pressed")

		} else if keyEvent.Name == fyne.KeyDown {
			if currentPos < (totalBoardObjects - totalBoardObjectsInRow) {
				currentPos += totalBoardObjectsInRow
				posUpdate = true
			}
			// log.Println("[DEBUG] Down Pressed")

		} else if keyEvent.Name == fyne.KeyLeft {
			if (currentPos % totalBoardObjectsInRow) != 0 {
				currentPos--
				posUpdate = true
			}
			// log.Println("[DEBUG] Left Pressed")

		} else if keyEvent.Name == fyne.KeyRight {
			if (currentPos % totalBoardObjectsInRow) != totalBoardObjectsInRowLessOne {
				currentPos++
				posUpdate = true
			}
			// log.Println("[DEBUG] Right Pressed")
		}

		if posUpdate {
			board[currentPos].value = "[X]"
			board[currentPos].elementColor = color.RGBA{0, 0, 255, 255}
			// log.Println("[DEBUG] Current Pos: ", currentPos)
			//position as X,Y
			x := currentPos % totalBoardObjectsInRow
			y := currentPos / totalBoardObjectsInRow
			log.Println("[DEBUG] Current Pos: ", x, y)

			board[oldPos].value = "[_]"
			board[oldPos].elementColor = color.RGBA{255, 255, 255, 255}

			grid = UpdateGrid(totalBoardObjects, totalBoardObjectsInRow, board)
			UpdateBoard(mainSpaceTrixWindow, toolbar, grid)
		}

	})

	mainSpaceTrixWindow.ShowAndRun()
}
