package main

import (
	"KeyEvent/ui"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

type customTheme struct {
	fyne.Theme
}

func (t customTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		return theme.DefaultTheme().Font(s)
	}
	// 通常のテキストも太字フォントを使用
	return theme.DefaultTheme().Font(fyne.TextStyle{Bold: true})
}

func main() {
	os.Setenv("FYNE_FONT", "NotoSansJP-VariableFont_wght.ttf")
	myApp := app.New()
	myApp.Settings().SetTheme(&customTheme{theme.DefaultTheme()})

	myWindow := myApp.NewWindow("KeyEvent")
	myWindow.Resize(fyne.NewSize(400, 400))
	ui.SetupUI(myWindow)
	myWindow.ShowAndRun()
}
