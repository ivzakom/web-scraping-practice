package lot_usecase

import (
	"context"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
)

type LotService interface {
	Create(lot entity.Lot) entity.Lot
}

type lotUseCase struct {
	lotService LotService
}

func (u lotUseCase) CreateLot(ctx context.Context, dto CreateLotDto) (string, error) {
	// map: DTO -> lot
	//u.lotService.Create(lot)
	return "", nil
}

func (u lotUseCase) GetAllLots(ctx context.Context) ([]entity.LotView, error) {
	return nil, nil
}
func (u lotUseCase) UpdateLots(ctx context.Context) error {
	return nil
}
