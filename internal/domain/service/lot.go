package service

import (
	"context"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
)

type LotStorage interface {
	GetOne(id string) entity.Lot
	GetAll(ctx context.Context) ([]entity.LotView, error)
	Create(lot entity.Lot) entity.Lot
	Delete(lot entity.Lot) error
}

type LotScraper interface {
	ScrapLot() ([]entity.Lot, error)
}

type lotService struct {
	lotStorage LotStorage
	lotScraper LotScraper
}

func NewLotService(storage LotStorage, scraper LotScraper) *lotService {
	return &lotService{storage, scraper}
}

func (s lotService) GetOne(id string) entity.Lot {
	return entity.Lot{}
}

func (s lotService) GetAll(ctx context.Context) ([]entity.LotView, error) {
	return s.lotStorage.GetAll(ctx)
}

func (s lotService) Create(lot entity.Lot) entity.Lot {
	return s.lotStorage.Create(lot)
}

func (s lotService) Delete(lot entity.Lot) error {
	return s.lotStorage.Delete(lot)
}

func (s lotService) ScrapLot() ([]entity.Lot, error) {
	return s.lotScraper.ScrapLot()
}
