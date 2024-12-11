package models

import "time"

type Token struct {
	AccessToken         string    `json:"access_token" db:"access_token"`
	RefreshToken        string    `json:"refresh_token" db:"refresh_token"`
	AccessTokenExpires  time.Time `json:"access_token_expiration" db:"access_token_expiration"`
	RefreshTokenExpires time.Time `json:"refresh_token_expiration" db:"refresh_token_expiration"`
}
