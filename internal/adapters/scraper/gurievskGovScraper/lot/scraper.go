package gurievskGovScraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type gurievskGovScraper struct {
	BaseURL   string
	BaseParam string
}

func NewGurievskGovScraper() *gurievskGovScraper {
	return &gurievskGovScraper{
		BaseURL:   "gurievsk.gov39.ru/grazhdanam/land-lease/",
		BaseParam: "SECTION_ID=7247",
	}
}

func (s *gurievskGovScraper) Scrap() ([]entity.Lot, error) {

	var err error

	c := colly.NewCollector(
		colly.AllowedDomains("gurievsk.gov39.ru"),
		colly.CacheDir("./cash/gurievsk_cache"),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36"

	detailCollector := c.Clone()

	Lots := make([]entity.Lot, 0, 200)

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML(`a.inner-link`, func(e *colly.HTMLElement) {
		plotURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(plotURL, s.BaseURL) != -1 {
			err = detailCollector.Visit(plotURL)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	})

	//<div class="news-detail__date"> <span class="date">30.08.2024</span></div>
	detailCollector.OnHTML(`div.news-detail__date`, func(e *colly.HTMLElement) {

		newsDate := e.ChildText(".date")
		e.Request.Ctx.Put("newsDate", newsDate)

	})

	endOfList := false
	detailCollector.OnHTML(`div.news-detail`, func(e *colly.HTMLElement) {
		log.Println("lots found", e.Request.URL)

		date := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

		newsDate := e.Request.Ctx.Get("newsDate")
		if newsDate != "" {

			date, _ = time.Parse("02.01.2006", newsDate)
			log.Println("date ", date)

		}

		year := date.Year()

		if year < 2022 && year != 1 {
			log.Printf("Дошли до %d года", year)
			e.Request.Abort()
			endOfList = true
		}

		paragraphList := strings.Split(e.Text, "\n")

		for _, paragraph := range paragraphList {
			paragraph = strings.TrimSpace(paragraph)

			if strings.HasPrefix(paragraph, "Лот") {

				newPlot := entity.Lot{
					Description:     paragraph,
					DocURL:          e.Request.AbsoluteURL(e.Attr("href")),
					PublicationDate: date,
				}

				re := regexp.MustCompile(`(\d+\.\d+|\d+)\s*кв\.?\s*м\.?`)
				if matches := re.FindStringSubmatch(paragraph); len(matches) > 0 {
					newPlot.Square, _ = strconv.Atoi(matches[1])
				}

				re = regexp.MustCompile(`по адресу:\s*(.*?),?\s*(КН|площадью|$)`)
				if matches := re.FindStringSubmatch(paragraph); len(matches) > 0 {
					newPlot.Address = matches[0]
				}

				re = regexp.MustCompile(`Лот \d*`)
				if matches := re.FindStringSubmatch(paragraph); len(matches) > 0 {

					lotNum, isFound := strings.CutPrefix(matches[0], "Лот ")
					if isFound {
						num, numErr := strconv.Atoi(lotNum)
						if numErr != nil {
							return
						} else {
							newPlot.Num = num
						}
					}

				}

				Lots = append(Lots, newPlot)

			}
		}

	})

	err = c.Visit(fmt.Sprintf("https://%s?%s", s.BaseURL, s.BaseParam))
	if err != nil {
		return nil, err
	}
	for i := 2; i < 161; i++ {

		err = c.Visit(fmt.Sprintf("https://%s?%s&PAGEN_1=%d", s.BaseURL, s.BaseParam, i))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if endOfList {
			log.Println("endOfList")
			break
		}

	}

	return Lots, nil

}
