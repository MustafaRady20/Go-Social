package main

import (
	"fmt"
	"log"

	"github.com/MustafaRady20/Go-Social/internal/db"
	"github.com/MustafaRady20/Go-Social/internal/store"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found or error reading it: %v", err)
	}

	viper.SetDefault("ADDR", ":8080")
	viper.SetDefault("DB_ADDR", "postgres://postgres:postgres@localhost/social?sslmode=disable")
	viper.SetDefault("MAX_IDLE_CONNS", 30)
	viper.SetDefault("MAX_OPEN_CONNS", 30)
	viper.SetDefault("MAX_IDLE_TIME", "30m")

}
func main() {

	cfg := config{
		addr: viper.GetString("ADDR"),
		db: dbConfig{
			addr:         viper.GetString("DB_ADDR"),
			maxOpenConns: viper.GetInt("MAX_OPEN_CONNS"),
			maxIdleConns: viper.GetInt("MAX_IDLE_CONNS"),
			maxIdleTime:  viper.GetString("MAX_IDLE_TIME"),
		},
	}
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic()
	}
	defer db.Close()
	fmt.Println("database connection pool established")
	store := store.NewStorage(db)

	app := &aplication{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
