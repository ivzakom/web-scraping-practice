package main

import (
	"context"
	"fmt"
	mongodb "github.com/ivzakom/web-scraping-practice/internal/adapters/db/mongodb/lot"
	gurievskGovScraper "github.com/ivzakom/web-scraping-practice/internal/adapters/scraper/gurievskGovScraper/lot"
	"github.com/ivzakom/web-scraping-practice/internal/config"
	v1 "github.com/ivzakom/web-scraping-practice/internal/controller/http/v1"
	"github.com/ivzakom/web-scraping-practice/internal/domain/service"
	lot_usecase "github.com/ivzakom/web-scraping-practice/internal/domain/usecase/lot"
	mongo "github.com/ivzakom/web-scraping-practice/pkg/client/mongodb"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {

	cfg := config.GetConfig()
	cfgMongo := cfg.MongoDB
	MongoDBCient, err := mongo.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}

	lotStorage := mongodb.NewLotStorage(MongoDBCient)
	lotScraper := gurievskGovScraper.NewGurievskGovScraper()
	lotService := service.NewLotService(lotStorage, lotScraper)
	lotUseCase := lot_usecase.NewLotUseCase(lotService)
	lotHandler := v1.NewLotHandler(lotUseCase)

	router := httprouter.New()
	lotHandler.Register(router)

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {

	//logger := logging.GetLogger()
	//logger.Info("start server")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appdir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			//logger.Fatal(err)
		}
		//logger.Info("create socket ")
		socketPath := path.Join(appdir, "app.sock")
		//logger.Debugf("socket path: %s", socketPath)

		//logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)

	} else {
		//logger.Info("listen unix")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		//logger.Info(fmt.Sprintf("start is lissening %s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	}

	if listenErr != nil {
		panic(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.Serve(listener)

}
