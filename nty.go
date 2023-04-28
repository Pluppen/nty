package main

import (
	"fmt"
    "flag"
    "strings"
	"log"
    "net/http"
)

type Notification struct {
	Message string
	Title string
	Priority string
	Tags string
	Destination string
	Server string
}

func main() {
    DEFAULT_DESTINATION := "<Default Destination>"

	message := flag.String("msg", "", "notification message")

	title := flag.String("title", "", "notification title")
	priority := flag.String("priority", "", "notification priority")
	tags := flag.String("tags", "", "notification tags")
	dest := flag.String("dest", DEFAULT_DESTINATION, "notification destination uid")
    server := flag.String("server", "https://ntfy.sh/", "ntfy.sh server")

	flag.Parse()

	if *message == "" {
		fmt.Println("Error: message flag is required")
		return
	}

	notification := &Notification{
		Message: *message,
		Title: *title,
		Priority: *priority,
		Tags: *tags,
        Destination: *dest,
        Server: *server,
	}

    sendNotification(*notification)
}

func sendNotification(notification Notification) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", notification.Server + notification.Destination, strings.NewReader(notification.Message))
	if err != nil {
		log.Fatal(err)
	}

    if notification.Title != "" {
        req.Header.Set("Title", notification.Title)
    }

    if notification.Priority != "" {
        req.Header.Set("Priority", notification.Priority)
    }

    if notification.Tags != "" {
        req.Header.Set("Tags", notification.Tags)
    }

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to send notification. Status code: %d\n", resp.StatusCode)
	}

	fmt.Printf("Sent notification: %s\n", notification.Message)
}
