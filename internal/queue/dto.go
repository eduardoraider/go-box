package queue

import "encoding/json"

type AppQueueDto struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	ID       int    `json:"id"`
}

func (q *AppQueueDto) Marshal() ([]byte, error) {
	return json.Marshal(q)
}

func (q *AppQueueDto) Unmarshal(data []byte) error {
	return json.Unmarshal(data, q)
}
