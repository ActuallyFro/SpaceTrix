package include

import (
	// "time"

	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateMenus(passedApp *fyne.App, passedWindow *fyne.Window) {
	newFyneApp := *passedApp
	mainSpaceTrixWindow := *passedWindow

	var helpMenuAbout *fyne.MenuItem
	var helpMenuSeeTime *fyne.MenuItem
	var helpMenu *fyne.Menu
	var fileMenu *fyne.Menu
	var mainMenu *fyne.MainMenu

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

	var clock *widget.Label
	clock = widget.NewLabel("")
	InitRecurringFunctionUpdateClock(clock)
	UpdateTime(clock)

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
}

/////
//https://github.com/fyne-io/fyne/issues/2257

type toolbarLabel struct {
	*widget.Label
}

func NewToolbarLabel(label string) widget.ToolbarItem {
	l := widget.NewLabelWithStyle(label, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	l.MinSize()
	return &toolbarLabel{l}
}

func (t *toolbarLabel) ToolbarObject() fyne.CanvasObject {
	return t.Label
}

/////

func InitToolbar(BoardCellInformation widget.ToolbarItem) *widget.Toolbar {
	// tempText := canvas.NewText("<empty>", color.RGBA64{0, 0, 0, 0})
	// tempText.Alignment = fyne.TextAlignCenter

	// var BoardInfo widget.ToolbarItem
	// BoardInfo.ToolbarObject() = tempText

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),

		widget.NewToolbarSeparator(),
		//Print current board item/filler info here
		BoardCellInformation,
		widget.NewToolbarSeparator(),

		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	// text := canvas.NewText(passedBoard[i][j].Value, passedBoard[i][j].ElementColor)
	// text.Alignment = fyne.TextAlignCenter

	// boardGrid.AddObject(text)
	return toolbar
}
