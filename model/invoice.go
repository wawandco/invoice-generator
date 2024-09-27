package model

import "time"

type Invoice struct {
	ID        int       `json:"id"`
	Data      []byte    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}
