package main

import (
	//stdout Debugging

	// https://developer.fyne.io/container/grid

	"github.com/actuallyfro/SpaceTrix/include"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {

	newFyneApp := app.New()
	newFyneApp.Settings().SetTheme(theme.DarkTheme())
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	BoardInfo := include.InitBoard(13, 21)

	var board [][]include.BoardElement
	board = include.CreateBoard(BoardInfo)

	include.CreateMenus(&newFyneApp, &mainSpaceTrixWindow)

	toolbar := include.InitToolbar()
	grid := include.CreateBoardGrid(board)
	include.UpdateBoard(mainSpaceTrixWindow, toolbar, grid)

	include.InitBoardInput(&newFyneApp, &mainSpaceTrixWindow, toolbar, &BoardInfo, &board)

	mainSpaceTrixWindow.ShowAndRun()
}
