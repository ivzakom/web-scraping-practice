package lot_usecase

import (
	"context"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
)

type LotService interface {
	Create(lot entity.Lot) entity.Lot
	GetAll(context.Context) ([]entity.LotView, error)
	ScrapLot() ([]entity.Lot, error)
}

type lotUseCase struct {
	lotService LotService
}

func NewLotUseCase(lotService LotService) *lotUseCase {
	return &lotUseCase{lotService}
}

func (u lotUseCase) CreateLot(ctx context.Context, dto CreateLotDto) (string, error) {
	// map: DTO -> lot
	//u.lotService.Create(lot)
	return "", nil
}

func (u lotUseCase) GetAllLots(ctx context.Context) ([]entity.LotView, error) {
	all, err := u.lotService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (u lotUseCase) UpdateLots(ctx context.Context) error {

	lots, err := u.lotService.ScrapLot()
	if err != nil {
		return err
	}
	for _, lot := range lots {
		u.lotService.Create(lot)
	}
	return err
}
