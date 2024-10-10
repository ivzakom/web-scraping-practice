package service

import "github.com/ivzakom/web-scraping-practice/internal/domain/entity"

type LotStorage interface {
	GetOne(id string) entity.Lot
	GetAll(limit, offset int) []entity.Lot
	Create(lot entity.Lot) entity.Lot
	Delete(lot entity.Lot) error
}

type lotService struct {
	lotStorage LotStorage
}

func NewLotService(storage LotStorage) *lotService {
	return &lotService{storage}
}

func (s lotService) GetOne(id string) entity.Lot {
	return entity.Lot{}
}

func (s lotService) GetAll(limit, offset int) []entity.Lot {
	return s.lotStorage.GetAll(limit, offset)
}

func (s lotService) Create(lot entity.Lot) entity.Lot {
	return s.lotStorage.Create(lot)
}

func (s lotService) Delete(lot entity.Lot) error {
	return s.lotStorage.Delete(lot)
}
