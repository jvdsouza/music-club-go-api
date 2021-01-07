package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key"`
	SpotifyURL string `json:"spotifyURL"`
}

type CreateAlbumInput struct {
	gorm.Model
	SpotifyURL string `json:"spotifyURL"`
}

type UpdateAlbumInput struct {
	gorm.Model
	SpotifyURL string `json:"spotifyURL"`
}
