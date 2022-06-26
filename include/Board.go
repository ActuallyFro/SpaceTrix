package include

import (
	// "time"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type BoardElement struct {
	Value        string
	ElementColor color.RGBA
}

type BoardCoord struct {
	x int
	y int
}

type BoardPositions struct {
	TotalX        int
	TotalY        int
	TotalElements int

	CenterPosition  BoardCoord
	CurrentPosition BoardCoord

	CurrentPositionIndex int
}

func InitBoard(totalX int, totalY int) BoardPositions {
	var tempBoardPos BoardPositions

	tempBoardPos.TotalX = totalX
	tempBoardPos.TotalY = totalY
	tempBoardPos.TotalElements = totalX * totalY

	tempBoardPos.CenterPosition.x = totalX / 2
	tempBoardPos.CenterPosition.y = totalY / 2

	tempBoardPos.CurrentPosition.x = tempBoardPos.CenterPosition.x
	tempBoardPos.CurrentPosition.y = tempBoardPos.CenterPosition.y

	tempBoardPos.CurrentPositionIndex = tempBoardPos.CurrentPosition.x + tempBoardPos.CurrentPosition.y*tempBoardPos.TotalX

	return tempBoardPos
}

// func (x BoardPositions) EXAMPLE-FUNCT() {
// }

func CreateBoard(totalRows int, totalColumns int) [][]BoardElement {
	var board [][]BoardElement
	board = make([][]BoardElement, totalColumns)
	for i := 0; i < totalColumns; i++ {
		board[i] = make([]BoardElement, totalRows)
	}

	for indexY := 0; indexY < totalColumns; indexY++ {
		for indexX := 0; indexX < totalRows; indexX++ {
			if indexX == totalRows/2 && indexY == totalColumns/2 { //Copilot generation
				board[indexY][indexX].Value = "[X]"
				board[indexY][indexX].ElementColor = color.RGBA{0, 0, 255, 255}
			} else {
				board[indexY][indexX].Value = "[-]"
				board[indexY][indexX].ElementColor = color.RGBA{255, 255, 255, 255}

			}
		}
	}
	return board
}

func CreateBoardGrid(passedBoard [][]BoardElement) *fyne.Container {
	boardGrid := container.New(layout.NewGridLayout(len(passedBoard[0])))

	for i := 0; i < len(passedBoard); i++ { //Y
		for j := 0; j < len(passedBoard[0]); j++ { //X

			text := canvas.NewText(passedBoard[i][j].Value, passedBoard[i][j].ElementColor)
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
