package main

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found or error reading it: %v", err)
	}

	viper.SetDefault("ADDR", ":8080")
}
func main() {

	cfg := config{
		addr: viper.GetString("ADDR"),
	}
	app := &aplication{
		config: cfg,
	}
	mux := app.mount()
	app.run(mux)
}
