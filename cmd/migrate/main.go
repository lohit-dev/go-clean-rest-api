package migrate

import (
	"log"

	"github.com/lohit-dev/go-clean-rest-api/config"
	"github.com/lohit-dev/go-clean-rest-api/internal/store"
)

func main() {
	cfg, _ := config.New()
	db, _ := store.New(cfg.DatabaseURL)

	db.DB().AutoMigrate(
	// all models
	)

	log.Println("migrations complete")
}
