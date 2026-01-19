package domain

import "time"

type ProfileResponse struct {
	Username  string
	Email     string
	CreatedAt time.Time
}