package main

import (
	"os"
	"KeyEvent/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	os.Setenv("FYNE_FONT", "NotoSansJP-VariableFont_wght.ttf")
	myApp := app.New()
	myWindow := myApp.NewWindow("KeyEvent")
	myWindow.Resize(fyne.NewSize(400, 350))
	ui.SetupUI(myWindow)
	myWindow.ShowAndRun()
}
