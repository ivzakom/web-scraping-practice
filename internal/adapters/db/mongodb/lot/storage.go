package mongodb

import (
	. "github.com/ivzakom/web-scraping-practice/internal/domain/entity"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type lotStorage struct {
	db *mongo.Database
}

func NewLotStorage(db *mongo.Database) *lotStorage {
	return &lotStorage{db: db}
}

func (bs *lotStorage) GetOne(id string) *Lot {
	return nil
}
func (bs *lotStorage) GetAll(limit, offset int) []*Lot {
	return nil
}
func (bs *lotStorage) Create(lot Lot) *Lot {
	return nil
}
func (bs *lotStorage) Delete(lot Lot) error {
	return nil
}
