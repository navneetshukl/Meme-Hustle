package memes

import "time"

type Memes struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	ImageURL  string     `json:"image_url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
