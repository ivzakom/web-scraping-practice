package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/ivzakom/web-scraping-practice/internal/apperror"
	"github.com/ivzakom/web-scraping-practice/internal/domain/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
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

func (bs *lotStorage) GetOne(ctx context.Context, num int, url string) (lot entity.Lot, err error) {

	filter := bson.M{"num": num, "docURL": url}
	result := bs.db.Collection(lotsCollection).FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return lot, apperror.ErrorNotFound
		}
		return lot, fmt.Errorf("failed to find user by num: %d and url: %s, error: %v", num, url, result.Err())
	}
	if err = result.Decode(&lot); err != nil {
		return lot, fmt.Errorf("failed to decode user (num: %d, url: %s) frob DB error: %v", num, url, err)
	}
	return lot, nil

}
func (bs *lotStorage) GetAll(ctx context.Context) ([]entity.LotView, error) {

	cursor, err := bs.db.Collection(lotsCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Закрываем курсор в конце
	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}()

	// Объявляем слайс для результатов
	var lots []entity.LotView

	// Читаем результаты в слайс
	if err = cursor.All(ctx, &lots); err != nil {
		return nil, err
	}

	return lots, nil
}
func (bs *lotStorage) Create(l entity.Lot) error {

	_, err := bs.db.Collection(lotsCollection).InsertOne(context.Background(), l)
	if err != nil {
		return err
	}
	return nil
}

func (bs *lotStorage) GetLastDateUpdate(ctx context.Context) time.Time {

	lastUpdateDate := time.Now().AddDate(-1, 0, 0)

	pipeline := mongo.Pipeline{
		{{Key: "$sort", Value: bson.D{{Key: "notice_date", Value: -1}}}},
		{{Key: "$limit", Value: 1}},
		{{Key: "$project", Value: bson.D{{Key: "notice_date", Value: 1}}}}, // Проекция только для PublishDate
	}

	cursor, err := bs.db.Collection(lotsCollection).Aggregate(ctx, pipeline)
	if err != nil {
		return time.Time{}
	}
	defer cursor.Close(ctx)

	var result struct {
		NoticeDate time.Time `bson:"notice_date"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return lastUpdateDate
		}
		return result.NoticeDate // Возвращаем найденную максимальную дату
	}

	// Возвращаем минимальную дату при отсутствии документов
	return lastUpdateDate

}
