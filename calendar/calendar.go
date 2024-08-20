package calendar

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func CreateEvent(title, date, startTime, endTime, color string) error {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		return fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	client := GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Calendar client: %v", err)
	}

	event := &calendar.Event{
		Summary: title,
		Start: &calendar.EventDateTime{
			DateTime: fmt.Sprintf("%sT%s:00+09:00", date, startTime),
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: fmt.Sprintf("%sT%s:00+09:00", date, endTime),
			TimeZone: "Asia/Tokyo",
		},
		ColorId: color,
	}

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Printf("Error creating event: %v", err)
		return fmt.Errorf("unable to create event: %v", err)
	}

	log.Printf("Event created: %s", event.HtmlLink)
	return nil
}
