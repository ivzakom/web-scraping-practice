package main

import (
	"context"
	mongodb "github.com/ivzakom/web-scraping-practice/internal/adapters/db/mongodb/lot"
	"github.com/ivzakom/web-scraping-practice/internal/config"
	"github.com/ivzakom/web-scraping-practice/internal/domain/service"
	mongo "github.com/ivzakom/web-scraping-practice/pkg/client/mongodb"
)

func main() {

	cfg := config.GetConfig()
	cfgMongo := cfg.MongoDB
	MongoDBCient, err := mongo.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}

	lotStorage := mongodb.NewLotStorage(MongoDBCient)
	lotService := service.NewLotService(lotStorage)

	//// Инициализация скреперов
	//platformAScraper := scraper.NewPlatformAScraper("https://platforma.com/lots")
	//
	//// Инициализация службы скрапинга с скреперами
	//scraperService := scraper.NewScraperService(lotService, []scraper.Scraper{
	//	platformAScraper,
	//	platformBScraper,
	//	// Добавьте другие скреперы здесь
	//})
	//
	//// Запуск процесса скрапинга
	//err := scraperService.ScrapAll()
	//if err != nil {
	//	log.Fatalf("Ошибка при скрапинге: %v", err)
	//}
	//
	//// Пример получения лотов
	//lots, err := lotService.GetAll(10, 0)
	//if err != nil {
	//	log.Fatalf("Ошибка при получении лотов: %v", err)
	//}
	//
	//for _, lot := range lots {
	//	log.Printf("Лот: %+v\n", lot)
	//}
}
