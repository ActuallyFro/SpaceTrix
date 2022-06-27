package include

import (
	// "time"

	"math/rand"

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

	StartingRocks       int
	StartingEnemies     int
	StartingNPCs        int
	StartingElementsSet bool

	Rocks   []BoardCoord
	Enemies []BoardCoord
	NPCs    []BoardCoord

	CurrentPositionIndex int
}

func InitBoard(totalX int, totalY int) BoardPositions {
	var tempBoardPos BoardPositions

	tempBoardPos.TotalX = totalX
	tempBoardPos.TotalY = totalY
	tempBoardPos.TotalElements = totalX * totalY
	//log
	// fmt.Println("[DEBUG][InitBoard] Total Elements:", tempBoardPos.TotalElements)

	tempBoardPos.CenterPosition.x = totalX / 2
	tempBoardPos.CenterPosition.y = totalY / 2

	tempBoardPos.CurrentPosition.x = tempBoardPos.CenterPosition.x
	tempBoardPos.CurrentPosition.y = tempBoardPos.CenterPosition.y

	tempBoardPos.CurrentPositionIndex = tempBoardPos.CurrentPosition.x + tempBoardPos.CurrentPosition.y*tempBoardPos.TotalX

	tempBoardPos.StartingRocks = (tempBoardPos.TotalElements / 10) //always 10%
	// fmt.Println("[DEBUG][InitBoard] Rocks: ", tempBoardPos.StartingRocks)
	tempBoardPos.StartingEnemies = tempBoardPos.TotalElements / 20 //always 5%
	// fmt.Println("[DEBUG][InitBoard] Enemies: ", tempBoardPos.StartingEnemies)

	tempBoardPos.StartingNPCs = 1 //-- want it as a 1 in 50 chance!

	// tempBoardPos.StartingElementsSet = false
	tempBoardPos.Rocks = tempBoardPos.SetupStartingElements(tempBoardPos.StartingRocks)
	tempBoardPos.Enemies = tempBoardPos.SetupStartingElements(tempBoardPos.StartingEnemies)
	tempBoardPos.NPCs = tempBoardPos.SetupStartingElements(tempBoardPos.StartingNPCs)
	//print sizes of starting elements:
	// fmt.Println("[DEBUG][InitBoard] Rocks: ", len(tempBoardPos.Rocks))
	// fmt.Println("[DEBUG][InitBoard] Enemies: ", len(tempBoardPos.Enemies))
	// fmt.Println("[DEBUG][InitBoard] NPCs: ", len(tempBoardPos.NPCs))

	return tempBoardPos
}

// func (x BoardPositions) EXAMPLE-FUNCT() {
// }
func GenerateRandomCoord(minX int, maxX int, minY int, maxY int) BoardCoord { ///<COPILOT CODE>
	var tempCoord BoardCoord ///<COPILOT CODE>
	tempCoord.x = 0          ///<COPILOT CODE>
	tempCoord.y = 0          ///<COPILOT CODE>

	//get random number between min and max
	tempCoord.x = minX + rand.Intn(maxX-minX) ///<COPILOT CODE>
	tempCoord.y = minY + rand.Intn(maxY-minY) ///<COPILOT CODE>

	return tempCoord ///<COPILOT CODE>
}

//Interfaces are DOPE -- but are NOT OO! -- https://golangbot.com/structs-instead-of-classes/
// func (board BoardPositions) SetupStartingElements() {
// 	for i := 0; i < board.StartingRocks; i++ {
// 		randCoord := GenerateRandomCoord(0, board.TotalX, 0, board.TotalY)
// 		board.Rocks = append(board.Rocks, BoardCoord{x: randCoord.x, y: randCoord.y})
// 		fmt.Println("[DEBUG][SetupStartElems] Rock Added: ", randCoord.x, randCoord.y, " Total Rocks: ", len(board.Rocks))
// 	}

// 	for i := 0; i < board.StartingEnemies; i++ {
// 		randCoord := GenerateRandomCoord(0, board.TotalX, 0, board.TotalY)
// 		board.Enemies = append(board.Enemies, BoardCoord{x: randCoord.x, y: randCoord.y})
// 	}

// 	for i := 0; i < board.StartingNPCs; i++ {
// 		randCoord := GenerateRandomCoord(0, board.TotalX, 0, board.TotalY)
// 		board.NPCs = append(board.NPCs, BoardCoord{x: randCoord.x, y: randCoord.y})
// 	} ///<COPILOT CODE> -- FRAME for fors; rand() mine

// 	board.StartingElementsSet = true
// }

//this will replicate the calls, without needing to make a method for each call...
func (board BoardPositions) SetupStartingElements(amount int) []BoardCoord {
	var returnCoords []BoardCoord

	for i := 0; i < amount; i++ {
		randCoord := GenerateRandomCoord(0, board.TotalX, 0, board.TotalY)
		returnCoords = append(returnCoords, BoardCoord{x: randCoord.x, y: randCoord.y})
		// fmt.Println("[DEBUG][SetupStartElems] Element Added: ", randCoord.x, randCoord.y, " Total Elements: ", len(returnCoords))
	}

	return returnCoords
}

func BoardRemoveElement(passedBoardElementArray *[]BoardCoord, elementPos int) {
	tempArray := *passedBoardElementArray
	//Copilot Code:
	// (*passedBoardElementArray)[elementPos] = (*passedBoardElementArray)[len(*passedBoardElementArray)-1]
	// *passedBoardElementArray = (*passedBoardElementArray)[:len(*passedBoardElementArray)-1]

	tempArray = append(tempArray[:elementPos], tempArray[elementPos+1:]...)
	// passedBoard.Rocks = append(passedBoard.Rocks[:i], passedBoard.Rocks[i+1:]...)
}

func CreateBoard(passedBoard BoardPositions) [][]BoardElement {
	// func CreateBoard(totalRows int, totalColumns int) [][]BoardElement {
	// tempBoardInfo := *passedBoard
	var board [][]BoardElement
	board = make([][]BoardElement, passedBoard.TotalY)
	for i := 0; i < passedBoard.TotalY; i++ {
		board[i] = make([]BoardElement, passedBoard.TotalX)
	}

	for indexY := 0; indexY < passedBoard.TotalY; indexY++ {
		for indexX := 0; indexX < passedBoard.TotalX; indexX++ {
			if indexX == passedBoard.TotalX/2 && indexY == passedBoard.TotalY/2 { //Copilot generation
				board[indexY][indexX].Value = "[X]"
				board[indexY][indexX].ElementColor = color.RGBA{0, 0, 255, 255}
			} else {
				board[indexY][indexX].Value = "[-]"
				board[indexY][indexX].ElementColor = color.RGBA{255, 255, 255, 255}
			}
		}
	}

	//log total rocks
	// fmt.Println("[DEBUG][CreateBoard] Total Rocks: ", len(passedBoard.Rocks))
	for i := 0; i < len(passedBoard.Rocks); i++ {
		if board[passedBoard.Rocks[i].y][passedBoard.Rocks[i].x].Value == "[X]" {
			BoardRemoveElement(&passedBoard.Rocks, i) //Co-pilot and https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
		} else {
			board[passedBoard.Rocks[i].y][passedBoard.Rocks[i].x].Value = "[R]"
			board[passedBoard.Rocks[i].y][passedBoard.Rocks[i].x].ElementColor = color.RGBA{0, 255, 0, 255}
		}
		//Log creation
		// fmt.Println("[R] at x:", passedBoard.Rocks[i].x, "y:", passedBoard.Rocks[i].y)
	}

	for i := 0; i < len(passedBoard.Enemies); i++ {
		if board[passedBoard.Enemies[i].y][passedBoard.Enemies[i].x].Value == "[X]" {
			BoardRemoveElement(&passedBoard.Enemies, i) //Co-pilot and https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
		} else {
			board[passedBoard.Enemies[i].y][passedBoard.Enemies[i].x].Value = "[E]"
			board[passedBoard.Enemies[i].y][passedBoard.Enemies[i].x].ElementColor = color.RGBA{255, 0, 0, 255}
		}
	}
	for i := 0; i < len(passedBoard.NPCs); i++ {
		if board[passedBoard.NPCs[i].y][passedBoard.NPCs[i].x].Value == "[X]" {
			BoardRemoveElement(&passedBoard.NPCs, i) //Co-pilot and https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
		} else {

			board[passedBoard.NPCs[i].y][passedBoard.NPCs[i].x].Value = "[N]"
			board[passedBoard.NPCs[i].y][passedBoard.NPCs[i].x].ElementColor = color.RGBA{250, 220, 90, 255}
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
