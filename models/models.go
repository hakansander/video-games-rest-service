package models

import "time"

type Game struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      int       `json:"rating"`
	Genre       []Genre   `json:"-"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GameGenre struct {
	ID        int       `json:"id"`
	GameID    int       `json:"game_id"`
	GenreID   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
