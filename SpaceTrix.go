package main

import (
	"fmt"
	"math/rand"

	//stdout Debugging

	// https://developer.fyne.io/container/grid

	"github.com/actuallyfro/SpaceTrix/include"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
)

func ConvertStringToL33tInt(s string) int64 {
	var i int64
	for _, c := range s {
		i += int64(c)
	}
	//print int
	fmt.Println("[DEBUG] [Str2L33tInt] generated:", i)

	return i
}

func main() {

	newFyneApp := app.New()
	// newFyneApp.Settings().SetTheme(theme.DarkTheme()) //A ppARENTly ... func is deprecated... -_-
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	// seed := "SpaceTrix"
	seed := "SpaceTrix++"
	rand.Seed(ConvertStringToL33tInt(seed))
	//convert string "seed" to int64

	// rand.Seed("SpaceTrix")
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
