package main

import (
	"flag"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"

	"github.com/tarunngusain08/culturehub/config"
	"github.com/tarunngusain08/culturehub/http/rest"
	"github.com/tarunngusain08/culturehub/pkg/db"
	"github.com/tarunngusain08/culturehub/pkg/db/postgres"
	"github.com/tarunngusain08/culturehub/pkg/log"
	"github.com/tarunngusain08/culturehub/pkg/models"
)

var resetDB = flag.Bool("m", false, "migrate database")

func main() {
	logger := log.New("main")
	flag.Parse()

	if e := config.Startup(); e != nil {
		logger.Fatal("config startup error", zap.Error(e))
	}

	if *resetDB {
		e := db.ResetDB()
		if e != nil {
			logger.Fatal("db reset error", zap.Error(e))
		}
	}

	db, err := db.NewDB(postgres.ConnectionStrWithDB())
	if err != nil {
		logger.Fatal("db connection error", zap.Error(err))
	}

	dao := models.DAO(db)
	if *resetDB {
		e := dao.Migrate()
		if e != nil {
			logger.Fatal("db migration error", zap.Error(e))
		}
		os.Exit(0)
	}

	if err := rest.Serve(dao); err != nil {
		logger.Fatal("rest server error", zap.Error(err))
	}
}
