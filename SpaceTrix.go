package main

import (
	"log" //stdout Debugging

	"image/color"

	// https://developer.fyne.io/container/grid

	"github.com/actuallyfro/SpaceTrix/include"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

type currentPositionCoord struct {
	x int
	y int
}

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
	// hasBoardBeenInitialized := false

	var board [][]include.BoardElement

	// if !hasBoardBeenInitialized {
	board = include.CreateBoard(totalBoardObjectsX, totalBoardObjectsY)

	include.CreateMenus(&newFyneApp, &mainSpaceTrixWindow)

	toolbar := include.InitToolbar()
	grid := include.CreateBoardGrid(board)
	include.UpdateBoard(mainSpaceTrixWindow, toolbar, grid)

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
			board[currentPosition.y][currentPosition.x].Value = "[X]"
			board[currentPosition.y][currentPosition.x].ElementColor = color.RGBA{0, 0, 255, 255}

			log.Println("[DEBUG] Current Pos: ", currentPosition.x, currentPosition.y)

			board[oldPos.y][oldPos.x].Value = "[_]"
			board[oldPos.y][oldPos.x].ElementColor = color.RGBA{0, 0, 0, 255}
			grid = include.CreateBoardGrid(board)
			include.UpdateBoard(mainSpaceTrixWindow, toolbar, grid)

		}

	})

	mainSpaceTrixWindow.ShowAndRun()
}
