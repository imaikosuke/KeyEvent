package ui

import (
	"KeyEvent/calendar"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type colorButton struct {
	button  *widget.Button
	rect    *canvas.Rectangle
	border  *canvas.Rectangle
	colorID string
}

type customEntry struct {
	widget.Entry
}

func newCustomEntry() *customEntry {
	entry := &customEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *customEntry) TypedRune(r rune) {
	e.Entry.TypedRune(r)
	e.Refresh()
}

func (e *customEntry) TypedKey(event *fyne.KeyEvent) {
	e.Entry.TypedKey(event)
	e.Refresh()
}

func SetupUI(window fyne.Window) {
	titleEntry := newCustomEntry()
	titleEntry.SetPlaceHolder("イベントのタイトルを入力")

	yearEntry := newCustomEntry()
	yearEntry.SetPlaceHolder("年 (YYYY)")

	monthEntry := newCustomEntry()
	monthEntry.SetPlaceHolder("月 (MM)")

	dayEntry := newCustomEntry()
	dayEntry.SetPlaceHolder("日 (DD)")

	startHourEntry := newCustomEntry()
	startHourEntry.SetPlaceHolder("開始時 (HH)")

	startMinuteEntry := newCustomEntry()
	startMinuteEntry.SetPlaceHolder("開始分 (MM)")

	endHourEntry := newCustomEntry()
	endHourEntry.SetPlaceHolder("終了時 (HH)")

	endMinuteEntry := newCustomEntry()
	endMinuteEntry.SetPlaceHolder("終了分 (MM)")

	colorOptions := []struct {
		id    string
		color color.Color
	}{
		{"1", color.NRGBA{R: 66, G: 133, B: 244, A: 255}},
		{"2", color.NRGBA{R: 15, G: 157, B: 88, A: 255}},
		{"3", color.NRGBA{R: 156, G: 39, B: 176, A: 255}},
		{"4", color.NRGBA{R: 216, G: 27, B: 96, A: 255}},
		{"5", color.NRGBA{R: 244, G: 180, B: 0, A: 255}},
	}

	var colorButtons []*colorButton
	var selectedColor string
	colorGrid := container.New(layout.NewGridLayoutWithColumns(len(colorOptions)))

	for _, colorOption := range colorOptions {
		colorOption := colorOption // avoid closure capture issue
		rect := canvas.NewRectangle(colorOption.color)
		rect.SetMinSize(fyne.NewSize(30, 30))
		border := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
		border.SetMinSize(fyne.NewSize(34, 34))

		btn := &colorButton{
			button: widget.NewButton("", func() {
				selectedColor = colorOption.id
				for _, btn := range colorButtons {
					btn.border.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
					btn.border.Refresh()
				}
				border.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				border.Refresh()
			}),
			rect:    rect,
			border:  border,
			colorID: colorOption.id,
		}

		colorButtons = append(colorButtons, btn)
		buttonContainer := container.NewMax(rect, border, btn.button)
		colorGrid.Add(buttonContainer)
	}

	statusLabel := widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "タイトル", Widget: titleEntry},
			{Text: "年", Widget: yearEntry},
			{Text: "月", Widget: monthEntry},
			{Text: "日", Widget: dayEntry},
			{Text: "開始時", Widget: startHourEntry},
			{Text: "開始分", Widget: startMinuteEntry},
			{Text: "終了時", Widget: endHourEntry},
			{Text: "終了分", Widget: endMinuteEntry},
		},
	}

	submitButton := widget.NewButton("送信", func() {
		title := titleEntry.Text
		date := fmt.Sprintf("%s-%s-%s", yearEntry.Text, monthEntry.Text, dayEntry.Text)
		startTime := fmt.Sprintf("%s:%s", startHourEntry.Text, startMinuteEntry.Text)
		endTime := fmt.Sprintf("%s:%s", endHourEntry.Text, endMinuteEntry.Text)

		err := calendar.CreateEvent(title, date, startTime, endTime, selectedColor)
		if err != nil {
			statusLabel.SetText(fmt.Sprintf("イベントの作成に失敗しました: %v", err))
			log.Printf("Error: %v", err)
		} else {
			statusLabel.SetText("イベントが正常に作成されました！")
		}
	})

	window.SetContent(container.NewVBox(
		form,
		colorGrid,
		submitButton,
		statusLabel,
	))
}
