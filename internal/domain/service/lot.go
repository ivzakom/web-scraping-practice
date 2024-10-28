package service

import (
	"context"
	"errors"
	pkkRosreestr "github.com/ivzakom/web-scraping-practice/internal/adapters/api/pkkRosreestr/lot"
	torgiGov "github.com/ivzakom/web-scraping-practice/internal/adapters/api/torgiGov/lot"
	"github.com/ivzakom/web-scraping-practice/internal/apperror"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
	"github.com/jinzhu/copier"
	"strconv"
	"time"
)

type LotStorage interface {
	GetOne(ctx context.Context, num int, docUrl string) (entity.Lot, error)
	GetAll(ctx context.Context) ([]entity.LotView, error)
	Create(lot entity.Lot) error
}

type PkkScraper interface {
	GetLocationPoint(Decription string) (pkkRosreestr.PkkRosreestrLotDto, error)
}

type GurievskLotScraper interface {
	Scrap() ([]entity.Lot, error)
}

type TorgiGovScraper interface {
	ScrapNotices(ctx context.Context, params map[string]string) ([]torgiGov.TorgiGovLotDto, error)
}

type lotService struct {
	lotStorage         LotStorage
	pkkScraper         PkkScraper
	gurievskLotScraper GurievskLotScraper
	torgiGovScraper    TorgiGovScraper
}

func NewLotService(storage LotStorage, pkkScraper PkkScraper, gurievskLotScraper GurievskLotScraper, torgiGovScraper TorgiGovScraper) *lotService {
	return &lotService{
		storage,
		pkkScraper,
		gurievskLotScraper,
		torgiGovScraper}
}

func (s *lotService) GetOne(ctx context.Context, num int, docUrl string) (entity.Lot, error) {
	return s.lotStorage.GetOne(ctx, num, docUrl)
}

func (s *lotService) GetAll(ctx context.Context) ([]entity.LotView, error) {

	allLotView, err := s.lotStorage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return allLotView, nil
}

func (s *lotService) Create(lot entity.Lot) error {
	return s.lotStorage.Create(lot)
}

func (s *lotService) ScrapNewNotices(ctx context.Context, lastUpdateDate time.Time) ([]entity.Lot, error) {

	var (
		findedLots []entity.Lot
		err        error
	)

	prams := map[string]string{
		"page":    "0",
		"pubFrom": lastUpdateDate.Format("2006-01-02"),
	}

	for i := 0; ; i++ {

		prams["page"] = strconv.Itoa(i)

		lotsDto, scrapErr := s.torgiGovScraper.ScrapNotices(ctx, prams)
		if scrapErr != nil {
			if errors.Is(scrapErr, apperror.ErrorEOL) {
				return findedLots, nil
			} else {
				return nil, err
			}
		}

		for _, lotDto := range lotsDto {

			lot, convertErr := convertLot(lotDto)
			if convertErr != nil {
				return nil, convertErr
			}

			findedLots = append(findedLots, lot)

		}

	}

	return findedLots, nil

}

func (s *lotService) EnrichLotByPkkRosreestr(lot *entity.Lot) error {

	rosreestrLotDto, err := s.pkkScraper.GetLocationPoint(lot.Description)
	if err != nil {
		return err
	}

	copier.Copy(lot.RosreestrData, &rosreestrLotDto)
	lot.Address = rosreestrLotDto.Features[0].Attrs.Address
	lot.CadastreNumber = rosreestrLotDto.CadastreNumber

	return nil

}

func (s *lotService) UpdateCreateLot(lot entity.Lot) error {
	//TODO update
	return s.lotStorage.Create(lot)
}

func convertLot(lotDto torgiGov.TorgiGovLotDto) (entity.Lot, error) {

	lot := entity.Lot{
		Num:             0,
		NoticeNumber:    lotDto.NoticeNumber,
		Description:     lotDto.LotName,
		Address:         "",
		CadastreNumber:  "",
		Square:          0,
		DocURL:          lotDto.Url,
		PublicationDate: time.Time{},
		Price:           lotDto.PriceMin,
		RosreestrData: struct {
			Total         int    `json:"total"`
			TotalRelation string `json:"total_relation"`
			Features      []struct {
				Center struct {
					Y float64 `json:"y"`
					X float64 `json:"x"`
				} `json:"center"`
				Extent struct {
					Xmin float64 `json:"xmin"`
					Xmax float64 `json:"xmax"`
					Ymax float64 `json:"ymax"`
					Ymin float64 `json:"ymin"`
				} `json:"extent"`
				Sort  int64 `json:"sort"`
				Type  int   `json:"type"`
				Attrs struct {
					Address      string `json:"address"`
					CategoryType string `json:"category_type"`
					Cn           string `json:"cn"`
					Id           string `json:"id"`
				} `json:"attrs"`
			} `json:"features"`
		}{},
	}

	err := copier.Copy(&lot.TorgiGovData, &lotDto)
	if err != nil {
		return entity.Lot{}, err
	}

	return lot, nil

}
