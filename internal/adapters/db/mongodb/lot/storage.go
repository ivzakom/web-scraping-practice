package mongodb

import (
	"context"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	lotsCollection = "lots"
)

type lotStorage struct {
	db *mongo.Database
}

func NewLotStorage(db *mongo.Database) *lotStorage {
	return &lotStorage{db: db}
}

func (bs *lotStorage) GetOne(id string) entity.Lot {
	return entity.Lot{}
}
func (bs *lotStorage) GetAll(ctx context.Context) ([]entity.LotView, error) {

	cursor, err := bs.db.Collection(lotsCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Закрываем курсор в конце
	defer cursor.Close(ctx)

	// Объявляем слайс для результатов
	var lots []entity.LotView

	// Читаем результаты в слайс
	if err = cursor.All(ctx, &lots); err != nil {
		return nil, err
	}

	return lots, nil
}
func (bs *lotStorage) Create(lot entity.Lot) entity.Lot {
	return entity.Lot{}
}
func (bs *lotStorage) Delete(lot entity.Lot) error {
	return nil
}
