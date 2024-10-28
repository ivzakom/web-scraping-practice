package lot_usecase

import (
	"context"
	"github.com/ivzakom/web-scraping-practice/internal/controller/http/dto"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
	"time"
)

type LotService interface {
	Create(lot entity.Lot) error
	GetAll(context.Context) ([]entity.LotView, error)
	//ScrapLot() ([]entity.Lot, error)
	GetOne(context.Context, int, string) (entity.Lot, error)
	GetLastDateUpdate(context.Context) time.Time
	ScrapNewNotices(ctx context.Context, lastUpdateDate time.Time) ([]entity.Lot, error)
	EnrichLotByPkkRosreestr(*entity.Lot) error
	UpdateCreateLot(entity.Lot) error
}

type lotUseCase struct {
	lotService LotService
}

func NewLotUseCase(lotService LotService) *lotUseCase {
	return &lotUseCase{lotService}
}

func (u *lotUseCase) GetAllLots(ctx context.Context) ([]dto.LotViewDto, error) {
	allUseCaseDto, err := u.lotService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var httpDtos []dto.LotViewDto
	for _, UseCaseDto := range allUseCaseDto {

		dateShows := UseCaseDto.PublicationDate.String()

		httpDtos = append(httpDtos, dto.LotViewDto{
			Description:     UseCaseDto.Description,
			Address:         UseCaseDto.Address,
			CadastreNumber:  UseCaseDto.CadastreNumber,
			Square:          UseCaseDto.Square,
			DocURL:          UseCaseDto.DocURL,
			PublicationDate: dateShows,
		})

	}

	return httpDtos, nil
}

func (u *lotUseCase) UpdateLots(ctx context.Context) error {

	//lots, err := u.lotService.ScrapLot()
	//if err != nil {
	//	return err
	//}
	//for _, lot := range lots {
	//
	//	_, findOneErr := u.lotService.GetOne(ctx, lot.Num, lot.DocURL)
	//	if findOneErr != nil {
	//		if errors.Is(findOneErr, apperror.ErrorNotFound) {
	//			err = u.lotService.Create(lot)
	//			if err != nil {
	//				return err
	//			}
	//		} else {
	//			return findOneErr
	//		}
	//	}
	//
	//}
	//return err
	return nil
}

func (u *lotUseCase) GetNewLots(ctx context.Context) error {

	var err error

	lastDateUpdate := u.lotService.GetLastDateUpdate(ctx).Add(1 * time.Second)
	noticeLots, ScrapErr := u.lotService.ScrapNewNotices(ctx, lastDateUpdate)
	if ScrapErr != nil {
		return ScrapErr
	}
	for _, lot := range noticeLots {
		err = u.lotService.EnrichLotByPkkRosreestr(&lot)
		if err != nil {
			return err
		}
		err = u.lotService.UpdateCreateLot(lot)
		if err != nil {
			return err
		}
	}

	return nil
}
