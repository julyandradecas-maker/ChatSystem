package models

import "time"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Message struct {
	ID            string    `json:"id"`
	RoomID        string    `json:"room_id"`
	UserID        string    `json:"user_id"`
	Username      string    `json:"username"`
	Texto         string    `json:"texto"`
	FechaCreacion time.Time `json:"fechaCreacion"`
}

type Room struct {
	ID            string    `json:"id"`
	Nombre        string    `json:"nombre"`
	FechaCreacion time.Time `json:"fechaCreacion"`
}
