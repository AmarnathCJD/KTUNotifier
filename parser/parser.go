package parser

import (
	"log"
	"net/http"
	"strings"
	"time"

	"main/notifier"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL = "https://ktu.edu.in/eu/core/announcements.htm"
const UPDATE_INTERVAL = 10                                              // seconds
var ANNOUNCE []notifier.Announcement = make([]notifier.Announcement, 0) // global variable

func fetchSoup() (*goquery.Document, error) {
	req, _ := http.NewRequest("GET", BASE_URL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func Parse() ([]notifier.Announcement, error) {
	soup, err := fetchSoup()
	if err != nil {
		return nil, err
	}

	announcements := make([]notifier.Announcement, 0)

	soup.Find(".c-details").Eq(0).Find("tr").Each(func(i int, s *goquery.Selection) {
		if i == 20 {
			return
		}

		ann := notifier.Announcement{}
		li := s.Find("li").Eq(0)
		liHtml, _ := li.Html()
		ann.Title = strings.TrimSpace(strings.Split(strings.TrimPrefix(li.Find("b").Text(), ": "), "\n")[0])
		if desc := strings.Split(liHtml, "</b>"); len(desc) > 1 {
			ann.Description = strings.Replace(strings.TrimSpace(strings.Split(strings.Split(desc[1], ".<!-- </a> -->")[0], "\n")[0]), "<!-- &lt;/a&gt; -->", "", -1)
		}

		ann.Link = "https://ktu.edu.in" + strings.Split(li.Find("a").AttrOr("href", ""), "\n")[0]

		ann.Date = strings.Split(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(s.Find("td").Eq(0).Text(), "\n", ""), "\t", "")), "  ")[0]
		announcements = append(announcements, ann)
	})

	return announcements, nil
}

func FetchPoller() {
	ticker := time.NewTicker(UPDATE_INTERVAL * time.Second)
	for range ticker.C {
		aq, err := Parse()
		if err != nil {
			log.Println("Poller: An error occured while fetching announcements:", err)
			continue
		}

		for _, ann := range aq {
			if !checkDuplicate(ann) {
				if checkIfOld(ann) {
					ANNOUNCE = append(ANNOUNCE, ann)
					notifier.Notify(ann)
				}
			}
		}
	}
}

func checkDuplicate(ann notifier.Announcement) bool {
	for _, a := range ANNOUNCE {
		if a.Title == ann.Title {
			return true
		}
	}

	return false
}

func checkIfOld(ann notifier.Announcement) bool {
	t, _ := time.Parse("Mon Jan 2 15:04:05 MST 2006", ann.Date)
	return t.Before(time.Now())
}
