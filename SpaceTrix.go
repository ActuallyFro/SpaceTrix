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

func main() {

	newFyneApp := app.New()
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	include.InitRandSeedTime()

	BoardSeedStarting := include.UUIDInt64Lower(include.RandUUID())
	fmt.Println("[DEBUG] Starting Seed: ", BoardSeedStarting)

	rand.Seed(BoardSeedStarting)
	BoardInfo := include.InitBoard(13, 21)

	board := include.CreateBoard(BoardInfo)

	include.CreateMenus(&newFyneApp, &mainSpaceTrixWindow)

	toolbar := include.InitToolbar()
	grid := include.CreateBoardGrid(board)
	include.UpdateBoard(mainSpaceTrixWindow, toolbar, grid)

	include.InitBoardInput(&newFyneApp, &mainSpaceTrixWindow, toolbar, &BoardInfo, &board)

	mainSpaceTrixWindow.ShowAndRun()
}
