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

func SetupUI(window fyne.Window) {
	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Enter event title")

	yearEntry := widget.NewEntry()
	yearEntry.SetPlaceHolder("YYYY")

	monthEntry := widget.NewEntry()
	monthEntry.SetPlaceHolder("MM")

	dayEntry := widget.NewEntry()
	dayEntry.SetPlaceHolder("DD")

	startHourEntry := widget.NewEntry()
	startHourEntry.SetPlaceHolder("HH")

	startMinuteEntry := widget.NewEntry()
	startMinuteEntry.SetPlaceHolder("MM")

	endHourEntry := widget.NewEntry()
	endHourEntry.SetPlaceHolder("HH")

	endMinuteEntry := widget.NewEntry()
	endMinuteEntry.SetPlaceHolder("MM")

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
			{Text: "Title", Widget: titleEntry},
			{Text: "Year", Widget: yearEntry},
			{Text: "Month", Widget: monthEntry},
			{Text: "Day", Widget: dayEntry},
			{Text: "Start Hour", Widget: startHourEntry},
			{Text: "Start Minute", Widget: startMinuteEntry},
			{Text: "End Hour", Widget: endHourEntry},
			{Text: "End Minute", Widget: endMinuteEntry},
		},
	}

	submitButton := widget.NewButton("Submit", func() {
		title := titleEntry.Text
		date := fmt.Sprintf("%s-%s-%s", yearEntry.Text, monthEntry.Text, dayEntry.Text)
		startTime := fmt.Sprintf("%s:%s", startHourEntry.Text, startMinuteEntry.Text)
		endTime := fmt.Sprintf("%s:%s", endHourEntry.Text, endMinuteEntry.Text)

		err := calendar.CreateEvent(title, date, startTime, endTime, selectedColor)
		if err != nil {
			statusLabel.SetText(fmt.Sprintf("Failed to create event: %v", err))
			log.Printf("Error: %v", err)
		} else {
			statusLabel.SetText("Event created successfully!")
		}
	})

	window.SetContent(container.NewVBox(
		form,
		colorGrid,
		submitButton,
		statusLabel,
	))
}
