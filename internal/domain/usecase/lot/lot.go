package lot_usecase

import (
	"context"
	"errors"
	"github.com/ivzakom/web-scraping-practice/internal/apperror"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
)

type LotService interface {
	Create(lot entity.Lot) error
	GetAll(context.Context) ([]entity.LotView, error)
	ScrapLot() ([]entity.Lot, error)
	GetOne(context.Context, int, string) (entity.Lot, error)
}

type lotUseCase struct {
	lotService LotService
}

func NewLotUseCase(lotService LotService) lotUseCase {
	return lotUseCase{lotService}
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

		_, findOneErr := u.lotService.GetOne(ctx, lot.Num, lot.DocURL)
		if findOneErr != nil {
			if errors.Is(findOneErr, apperror.ErrorNotFound) {
				err = u.lotService.Create(lot)
				if err != nil {
					return err
				}
			} else {
				return findOneErr
			}
		}

	}
	return err
}
