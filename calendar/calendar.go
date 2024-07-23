package calendar

import (
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func CreateEvent(title, date, startTime, endTime, color string) error {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %w", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return fmt.Errorf("unable to parse client secret file to config: %w", err)
	}
	client := GetClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		return fmt.Errorf("unable to retrieve Calendar client: %w", err)
	}

	event := &calendar.Event{
		Summary: title,
		Start: &calendar.EventDateTime{
			DateTime: date + "T" + startTime + ":00-00:00",
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: date + "T" + endTime + ":00-00:00",
			TimeZone: "UTC",
		},
		ColorId: color,
	}

	calendarId := "primary"
	_, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		return fmt.Errorf("unable to create event: %w", err)
	}

	return nil
}
