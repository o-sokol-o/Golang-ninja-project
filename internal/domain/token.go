package domain

import "time"

type Token struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
}

type TokenResponse struct {
	Token string `json:"accessToken"`
}
