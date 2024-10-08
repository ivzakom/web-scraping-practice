package lot

type LotStorage interface {
	GetOne(id string) *Lot
	GetAll(limit, offset int) []*Lot
	Create(lot Lot) *Lot
	Delete(lot Lot) error
}

type lotService struct {
	lotStorage LotStorage
}

func NewLotService(storage LotStorage) *lotService {
	return &lotService{storage}
}

func (s *lotService) GetOne(id string) *Lot {
	return &Lot{}
}

func (s *lotService) GetAll(limit, offset int) []*Lot {
	return s.lotStorage.GetAll(limit, offset)
}

func (s *lotService) Create(lot Lot) *Lot {
	return s.lotStorage.Create(lot)
}

func (s *lotService) Delete(lot Lot) error {
	return s.lotStorage.Delete(lot)
}
