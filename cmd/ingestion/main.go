package main

import (
	"log"

	"github.com/avalokitasharma/topk-youtube-system/internal/config"
	"github.com/avalokitasharma/topk-youtube-system/internal/handlers"
	"github.com/avalokitasharma/topk-youtube-system/internal/kafka"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	producer := kafka.NewProducer(cfg.KafkaBrokers)

	r := gin.Default()
	r.POST("v1/ingest", handlers.NewIngestHandler(producer).Handle)
	log.Printf("Ingestion service listening on :%d", cfg.Port)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
