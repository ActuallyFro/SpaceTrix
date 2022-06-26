package main

import (
	"log" //stdout Debugging

	"image/color"

	// https://developer.fyne.io/container/grid

	"fyne.io/fyne/v2/dialog"
	"github.com/actuallyfro/SpaceTrix/include"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type currentPositionCoord struct {
	x int
	y int
}

type boardElement struct {
	value        string
	elementColor color.RGBA
}

func CreateBoardGrid(passedBoard [][]boardElement) *fyne.Container {
	boardGrid := container.New(layout.NewGridLayout(len(passedBoard[0])))

	for i := 0; i < len(passedBoard); i++ { //Y
		for j := 0; j < len(passedBoard[0]); j++ { //X

			text := canvas.NewText(passedBoard[i][j].value, passedBoard[i][j].elementColor)
			text.Alignment = fyne.TextAlignCenter

			boardGrid.AddObject(text)
		}
	}

	return boardGrid
}

func UpdateBoard(passedWindow fyne.Window, passedToolbar *widget.Toolbar, passedContent *fyne.Container) {
	ToolbarAndContent := container.NewBorder(passedToolbar, nil, nil, nil, passedContent)
	passedWindow.SetContent(ToolbarAndContent)
}

//For AI:
// func InitRecurringFunctionUpdateBoard(passedWindow fyne.Window, passedToolbar *widget.Toolbar, passedContent *fyne.Container) {
// 	// w.SetContent(clock)
// 	go func() {
// 		//updat every 100ms
// 		for range time.Tick(time.Millisecond * 500) {
// 			UpdateBoard(passedWindow, passedToolbar, passedContent)
// 		}
// 	}()
// }

func main() {

	newFyneApp := app.New()
	newFyneApp.Settings().SetTheme(theme.DarkTheme())
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	totalBoardObjectsX := 13
	totalBoardObjectsY := 21
	totalBoardObjects := totalBoardObjectsX * totalBoardObjectsY

	centerCell := currentPositionCoord{x: totalBoardObjectsX / 2, y: totalBoardObjectsY / 2}
	currentPosition := centerCell

	currentIndexPos := currentPosition.x + currentPosition.y*totalBoardObjectsX
	hasBoardBeenInitialized := false

	var board [][]boardElement
	var grid *fyne.Container

	var helpMenuAbout *fyne.MenuItem
	var clock *widget.Label
	var helpMenuSeeTime *fyne.MenuItem
	var helpMenu *fyne.Menu
	var fileMenu *fyne.Menu
	var mainMenu *fyne.MainMenu

	var toolbar *widget.Toolbar

	if !hasBoardBeenInitialized {
		board = make([][]boardElement, totalBoardObjectsY)
		for i := 0; i < totalBoardObjectsY; i++ {
			board[i] = make([]boardElement, totalBoardObjectsX)
		}

		for indexY := 0; indexY < totalBoardObjectsY; indexY++ {
			for indexX := 0; indexX < totalBoardObjectsX; indexX++ {
				if indexX == centerCell.x && indexY == centerCell.y { //Copilot generation
					board[indexY][indexX].value = "[X]"
					board[indexY][indexX].elementColor = color.RGBA{0, 0, 255, 255}
				} else {
					board[indexY][indexX].value = "[-]"
					board[indexY][indexX].elementColor = color.RGBA{255, 255, 255, 255}

				}
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

		grid = CreateBoardGrid(board)
		UpdateBoard(mainSpaceTrixWindow, toolbar, grid)

		hasBoardBeenInitialized = true
	}

	mainSpaceTrixWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		posUpdate := false
		oldPos := currentPosition

		if keyEvent.Name == fyne.KeyEscape {
			newFyneApp.Quit()

		} else if keyEvent.Name == fyne.KeyUp || keyEvent.Name == fyne.KeyW {
			if currentIndexPos >= totalBoardObjectsX {
				currentIndexPos -= totalBoardObjectsX
				currentPosition.y = currentIndexPos / totalBoardObjectsX
				posUpdate = true
			}

			// log.Println("[DEBUG] Up Pressed")

		} else if keyEvent.Name == fyne.KeyDown || keyEvent.Name == fyne.KeyS {
			if currentIndexPos < totalBoardObjects-totalBoardObjectsX {
				currentIndexPos += totalBoardObjectsX
				currentPosition.y = currentIndexPos / totalBoardObjectsX
				posUpdate = true
			}
			// log.Println("[DEBUG] Down Pressed")

		} else if keyEvent.Name == fyne.KeyLeft || keyEvent.Name == fyne.KeyA {
			if currentIndexPos%totalBoardObjectsX != 0 {
				currentIndexPos--
				currentPosition.x--
				posUpdate = true
			}
			// log.Println("[DEBUG] Left Pressed")

		} else if keyEvent.Name == fyne.KeyRight || keyEvent.Name == fyne.KeyD {
			if currentIndexPos%totalBoardObjectsX != totalBoardObjectsX-1 {
				currentIndexPos++
				currentPosition.x++
				posUpdate = true
			}
			// log.Println("[DEBUG] Right Pressed")
		}

		if posUpdate {
			board[currentPosition.y][currentPosition.x].value = "[X]"
			board[currentPosition.y][currentPosition.x].elementColor = color.RGBA{0, 0, 255, 255}

			log.Println("[DEBUG] Current Pos: ", currentPosition.x, currentPosition.y)

			board[oldPos.y][oldPos.x].value = "[_]"
			board[oldPos.y][oldPos.x].elementColor = color.RGBA{0, 0, 0, 255}
			grid = CreateBoardGrid(board)
			UpdateBoard(mainSpaceTrixWindow, toolbar, grid)

		}

	})

	mainSpaceTrixWindow.ShowAndRun()
}
