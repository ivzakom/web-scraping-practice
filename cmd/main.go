package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Plot struct {
	Description     string
	Address         string
	CadastreNumber  string
	Square          int
	DocURL          string
	PublicationDate time.Time
}

func main() {

	endOfList := false

	BaseUrl := "gurievsk.gov39.ru/grazhdanam/land-lease/"
	Area := "SECTION_ID=7247" //&PAGEN_1=161

	fName := "lots.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("Cannot close file %q: %s\n", fName, err)
			return
		}
	}()

	c := colly.NewCollector(
		colly.AllowedDomains("gurievsk.gov39.ru"),
		colly.CacheDir("./cash/gurievsk_cache"),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36"

	detailCollector := c.Clone()

	Lots := make([]Plot, 0, 200)

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML(`a.inner-link`, func(e *colly.HTMLElement) {
		plotURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(plotURL, BaseUrl) != -1 {
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
			if strings.HasPrefix(paragraph, "Лот") {

				newPlot := Plot{
					Description:     paragraph,
					DocURL:          e.Request.AbsoluteURL(e.Attr("href")),
					PublicationDate: date,
				}

				re := regexp.MustCompile(`(\d+\.\d+|\d+)\s*кв\.?\s*м\.?`)

				if matches := re.FindStringSubmatch(paragraph); len(matches) > 0 {
					newPlot.Square, _ = strconv.Atoi(matches[1])
				}

				re = regexp.MustCompile(`КН\s*\d{2}:\d{2}:\d{6,7}:\d+`)
				if matches := re.FindStringSubmatch(paragraph); len(matches) > 0 {
					newPlot.CadastreNumber = matches[0][5:]
				}

				re = regexp.MustCompile(`по адресу:\s*(.*?),?\s*(КН|площадью|$)`)
				if matches := re.FindStringSubmatch(paragraph); len(matches) > 0 {
					newPlot.Address = matches[0]
				}

				Lots = append(Lots, newPlot)

			}
		}

	})

	err = c.Visit(fmt.Sprintf("https://%s?%s", BaseUrl, Area))
	if err != nil {
		return
	}
	for i := 2; i < 161; i++ {

		err = c.Visit(fmt.Sprintf("https://%s?%s&PAGEN_1=%d", BaseUrl, Area, i))
		if err != nil {
			fmt.Println(err)
			return
		}
		if endOfList {
			log.Println("endOfList")
			break
		}

	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	err = enc.Encode(Lots)
	if err != nil {
		log.Fatal(err)
		return
	}
}
