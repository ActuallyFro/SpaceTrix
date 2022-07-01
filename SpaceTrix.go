package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	//stdout Debugging

	// https://developer.fyne.io/container/grid

	"github.com/actuallyfro/SpaceTrix/include"

	"github.com/google/uuid"

	"fyne.io/fyne/v2" //fyne.*
	"fyne.io/fyne/v2/app"
)

func GenerateUUIDStr() string {
	randIntStr := uuid.New().String()

	// randIntStr := fmt.Sprintf("%x", rand.Int63())
	fmt.Println("[DEBUG] [GenerateUUID] generated:", randIntStr)

	return randIntStr

}

func RandUUID() uuid.UUID {
	return uuid.New()
}

func GenerateInt64() int64 {

	randIntStr := rand.Int63()

	//cast randIntStr as randInt64 (int64)
	randInt64 := int64(randIntStr)

	// randIntStr := fmt.Sprintf("%x", rand.Int63())
	fmt.Println("[DEBUG] [GenerateInt64] generated:", randIntStr)

	return randInt64

}

func ExtractLowerIntFromUUID(passedUUID uuid.UUID) int64 {
	return int64(binary.BigEndian.Uint64(passedUUID[0:8]))
}

func ExtractUpperIntFromUUID(passedUUID uuid.UUID) int64 {
	return int64(binary.BigEndian.Uint64(passedUUID[8:16]))
}

// func ConvertStringToL33tInt(s string) int64 {
// 	var i int64
// 	for _, c := range s {
// 		i += int64(c)
// 	}
// 	//print int
// 	fmt.Println("[DEBUG] [Str2L33tInt] generated:", i)

// 	return i
// }

func main() {

	newFyneApp := app.New()
	// newFyneApp.Settings().SetTheme(theme.DarkTheme()) //A ppARENTly ... func is deprecated... -_-
	mainSpaceTrixWindow := newFyneApp.NewWindow("SpaceTrix - a WASD Adventure")
	mainSpaceTrixWindow.Resize(fyne.NewSize(640, 480))

	// seed := "SpaceTrix" //returns 915
	// seed := "SpaceTrix++"
	rand.Seed(time.Now().UnixNano()) //Random start, to then create a re-usable seed

	//Old
	// seed := GenerateUUIDStr()
	// seed := "d5acc140-55dc-4e86-8a40-bd7931df3f92" //generates 2470
	// ConvertStringToL33tInt("d5acc140-55dc-4e86-8a40-bd7931df3f92")
	// rand.Seed(ConvertStringToL33tInt(seed))

	//New
	newUUID := RandUUID()

	u1 := ExtractLowerIntFromUUID(newUUID)
	// u2 := ExtractUpperIntFromUUID(newUUID)

	//Hex print of UUID upper/lower
	// fmt.Printf("%x %x\n", u1, u2)

	// u2 := ExtractUpperIntFromUUID(newUUID)
	// fmt.Println("[DEBUG] [UUID] split: ", u1, u2)

	//log int
	fmt.Println("[DEBUG] [main] seedInt:", newUUID)
	fmt.Println("[DEBUG] [main] generating seed with:", u1)

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
