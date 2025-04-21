package utils

import (
	"context"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {
	// url := `https://www.setlist.fm/setlists/billie-eilish-1bc3b540.html`
	url := `https://www.setlist.fm/setlists/oasis-bd6bd7e.html`

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var upcoming string
	var stats string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML(`.upcomingSetlistsList`, &upcoming, chromedp.ByQuery),
		chromedp.OuterHTML(`.artistStatsTeaser`, &stats, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	upcomingHTML, err := goquery.NewDocumentFromReader(strings.NewReader(upcoming))

	if err != nil {
		log.Fatal(err)
	}

	upcomingHTML.Find(".setlist:not(.hidden)").Each(func(i int, s *goquery.Selection) {
		day := s.Find("strong.big").Text()
		month := s.Find("strong.text-uppercase").Text()
		year := strings.TrimSpace(s.Find("span.smallDateBlock span").Text())
		id := ""

		url, exists := s.Find(".content a").Attr("href")

		if exists {
			trimmed := strings.TrimSuffix(url, ".html")
			split := strings.Split(trimmed, "-")
			id = split[len(split)-1]
		}

		venue := s.Find(".content a span strong").Text()
		location := s.Find(".content span.subline span").Text()

		log.Println(day, month, year, venue, location, id)
	})

	statsHTML, err := goquery.NewDocumentFromReader(strings.NewReader(stats))

	if err != nil {
		log.Fatal(err)
	}

	statsHTML.Find("li").Each(func(i int, s *goquery.Selection) {
		song := s.Find("a").Text()
		count := s.Find("span").Text()
		log.Println(song, count)
	})
}
