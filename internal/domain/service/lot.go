package service

import (
	"context"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
)

type LotStorage interface {
	GetOne(ctx context.Context, num int, docUrl string) (entity.Lot, error)
	GetAll(ctx context.Context) ([]entity.LotView, error)
	Create(lot entity.Lot) error
}

type LotScraper interface {
	Scrap() ([]entity.Lot, error)
}

type lotService struct {
	lotStorage LotStorage
	lotScraper LotScraper
}

func NewLotService(storage LotStorage, scraper LotScraper) *lotService {
	return &lotService{storage, scraper}
}

func (s *lotService) GetOne(ctx context.Context, num int, docUrl string) (entity.Lot, error) {
	return s.lotStorage.GetOne(ctx, num, docUrl)
}

func (s *lotService) GetAll(ctx context.Context) ([]entity.LotView, error) {
	return s.lotStorage.GetAll(ctx)
}

func (s *lotService) Create(lot entity.Lot) error {
	return s.lotStorage.Create(lot)
}

func (s *lotService) ScrapLot() ([]entity.Lot, error) {
	return s.lotScraper.Scrap()
}
