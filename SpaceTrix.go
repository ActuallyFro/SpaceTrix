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

	// seed := "SpaceTrix" //returns 915
	// seed := "SpaceTrix++"

	// seed := "d5acc140-55dc-4e86-8a40-bd7931df3f92" //generates 2470

	newUUID := include.RandUUID()

	u1 := include.UUIDInt64Lower(newUUID)

	fmt.Println("[DEBUG] [main] seedUUID:", newUUID)
	fmt.Println("[DEBUG] [main] seeding lower int64 of:", u1)

	rand.Seed(u1)
	BoardInfo := include.InitBoard(13, 21)

	board := include.CreateBoard(BoardInfo)

	include.CreateMenus(&newFyneApp, &mainSpaceTrixWindow)

	toolbar := include.InitToolbar()
	grid := include.CreateBoardGrid(board)
	include.UpdateBoard(mainSpaceTrixWindow, toolbar, grid)

	include.InitBoardInput(&newFyneApp, &mainSpaceTrixWindow, toolbar, &BoardInfo, &board)

	mainSpaceTrixWindow.ShowAndRun()
}
