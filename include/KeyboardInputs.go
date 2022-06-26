package include

import (
	// "time"

	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func InitBoardInput(passedApp *fyne.App, passedWindow *fyne.Window, passedToolbar *widget.Toolbar, passedPositions *BoardPositions, passedBoard *[][]BoardElement) {
	newFyneApp := *passedApp
	mainSpaceTrixWindow := *passedWindow
	BoardInfo := *passedPositions
	board := *passedBoard
	// toolbar := *passedToolbar

	mainSpaceTrixWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		posUpdate := false
		oldPos := BoardInfo.CurrentPosition

		if keyEvent.Name == fyne.KeyEscape {
			newFyneApp.Quit()

		} else if keyEvent.Name == fyne.KeyUp || keyEvent.Name == fyne.KeyW {
			if BoardInfo.CurrentPositionIndex >= BoardInfo.TotalX {
				BoardInfo.CurrentPositionIndex -= BoardInfo.TotalX
				BoardInfo.CurrentPosition.y = BoardInfo.CurrentPositionIndex / BoardInfo.TotalX
				posUpdate = true
			}

			// log.Println("[DEBUG] Up Pressed")

		} else if keyEvent.Name == fyne.KeyDown || keyEvent.Name == fyne.KeyS {
			if BoardInfo.CurrentPositionIndex < BoardInfo.TotalElements-BoardInfo.TotalX {
				BoardInfo.CurrentPositionIndex += BoardInfo.TotalX
				BoardInfo.CurrentPosition.y = BoardInfo.CurrentPositionIndex / BoardInfo.TotalX
				posUpdate = true
			}
			// log.Println("[DEBUG] Down Pressed")

		} else if keyEvent.Name == fyne.KeyLeft || keyEvent.Name == fyne.KeyA {
			if BoardInfo.CurrentPositionIndex%BoardInfo.TotalX != 0 {
				BoardInfo.CurrentPositionIndex--
				BoardInfo.CurrentPosition.x--
				posUpdate = true
			}
			// log.Println("[DEBUG] Left Pressed")

		} else if keyEvent.Name == fyne.KeyRight || keyEvent.Name == fyne.KeyD {
			if BoardInfo.CurrentPositionIndex%BoardInfo.TotalX != BoardInfo.TotalX-1 {
				BoardInfo.CurrentPositionIndex++
				BoardInfo.CurrentPosition.x++
				posUpdate = true
			}
			// log.Println("[DEBUG] Right Pressed")
		}

		if posUpdate {
			board[BoardInfo.CurrentPosition.y][BoardInfo.CurrentPosition.x].Value = "[X]"
			board[BoardInfo.CurrentPosition.y][BoardInfo.CurrentPosition.x].ElementColor = color.RGBA{0, 0, 255, 255}

			log.Println("[DEBUG] Current Pos: ", BoardInfo.CurrentPosition.x, BoardInfo.CurrentPosition.y)

			board[oldPos.y][oldPos.x].Value = "[_]"
			board[oldPos.y][oldPos.x].ElementColor = color.RGBA{0, 0, 0, 255}
			grid := CreateBoardGrid(board)
			UpdateBoard(mainSpaceTrixWindow, passedToolbar, grid)

		}

	})

}
