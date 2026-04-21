package handlers

import (
	"github.com/avalokitasharma/topk-youtube-system/internal/kafka"
	"github.com/avalokitasharma/topk-youtube-system/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IngestHandler struct {
	producer *kafka.Producer
}

func NewIngestHandler(p *kafka.Producer) *IngestHandler {
	return &IngestHandler{
		producer: p,
	}
}

func(h *IngestHandler) Handle(c *gin.Context)) {
	var e models.Event
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if e.Timestamp == "" {
		e.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}
	if err := h.producer.Publish(e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "event_accepted","video_id":e.VideoId})
}