package notifier

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

// TODO: Add more notifier services, for now only telegram is added.

type Announcement struct {
	Title       string `json:"title,omitempty"`
	Link        string `json:"link,omitempty"`
	Date        string `json:"date,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	BOT_TOKEN = ""
	CHAT_ID   = ""
)

func init() {
	if os.Getenv("BOT_TOKEN") != "" {
		BOT_TOKEN = os.Getenv("BOT_TOKEN")
	} else {
		log.Println("Notifier: BOT_TOKEN not found in environment variables.")
	}

	if os.Getenv("CHAT_ID") != "" {
		CHAT_ID = os.Getenv("CHAT_ID")
	} else {
		log.Println("Notifier: CHAT_ID not found in environment variables.")
	}

	if BOT_TOKEN == "" || CHAT_ID == "" {
		log.Println("Notifier: Telegram notifier not initialized.")
	}
}

func formatMessage(announcements Announcement) string {
	msg := fmt.Sprintf("<b>%s</b>\n\n", announcements.Title)
	if announcements.Description != "" {
		msg += fmt.Sprintf("<code>%s</code>\n", announcements.Description)
	}
	msg += fmt.Sprintf("<a href=\"%s\">Read more</a>\n\n", announcements.Link)
	msg += fmt.Sprintf("<i>%s</i>", announcements.Date)

	return msg
}

func Notify(announcements Announcement) {
	if BOT_TOKEN == "" || CHAT_ID == "" {
		fmt.Println("Notifier: Telegram notifier not initialized.")
		return
	}

	msg := formatMessage(announcements)
	uri := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BOT_TOKEN)
	values := map[string]string{"chat_id": CHAT_ID, "text": msg, "parse_mode": "HTML"}
	urx := buildQuery(uri, values)
	req, err := http.NewRequest("GET", urx, nil)

	if err != nil {
		log.Println("Notifier: An error occured while sending notification:", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Notifier: An error occured while sending notification:", err)
		return
	}

	if res.StatusCode != 200 {
		log.Println("Notifier: An error occured while sending notification: Status code not 200.")
		return
	}

	log.Println("Notifier: Notification sent.")
}

func buildQuery(uri string, values map[string]string) string {
	query := ""
	for key, value := range values {
		query += fmt.Sprintf("%s=%s&", key, url.QueryEscape(value))
	}

	return fmt.Sprintf("%s?%s", uri, query)
}
