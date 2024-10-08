package scraper

import (
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity/lot"
)

type Scraper interface {
	ScrapLot() ([]lot.Lot, error)
}
