package models

type Event struct {
	VideoId   string `json:"video_id"`
	EventType string `json:"event_type"`
	Value     int64  `json:"value"`
	Region    string `json:"region"`
	Category  string `json:"category"`
}

type TopKItem struct {
	Rank     int     `json:"rank"`
	VideoId  string  `json:"video_id"`
	Score    float64 `json:"score"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
}
