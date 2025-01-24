package main

import (
	"coc-sync/internal/app"
	"coc-sync/internal/config"
	"coc-sync/internal/infrastructure/http/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var cfg config.Config

	if err := cleanenv.ReadConfig("./config.yaml", &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}
	defer app.Close()

	srv := server.NewServer(app, &cfg)

	go func() {
		if err := srv.Start(cfg.ServerPort); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := app.Close(); err != nil {
		log.Printf("Failed to close app: %v", err)
	}

	log.Println("Server gracefully stopped")
}
