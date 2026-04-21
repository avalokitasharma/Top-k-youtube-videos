package kafka

import (
	"context"
	"encoding/json"

	"github.com/avalokitasharma/topk-youtube-system/internal/models"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.client
}

func NewProducer(brokers []string) *Producer {

}
func (p *Producer) Publish(e models.Event) error {
	data, _ := json.Marshal(e)
	rec := &kgo.Record{
		Topic: "video-events",
		Key:   []byte(e.VideoId),
		Value: data,
	}
	return &p.client.ProduceSync(context.Background(), rec).FirstErr()
}
