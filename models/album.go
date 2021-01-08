package models

import (
	"gorm.io/gorm"
)

// Album struct to model database
type Album struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key"`
	SpotifyURL string `json:"spotifyURL"`
	AlbumTitle   string `json:"albumTitle"`
	AlbumArtists string `json:"albumArtists"`
	AlbumArt     string `json:"albumArt"`
}

// CreateAlbumInput struct to specify input for creating an album in the database
type CreateAlbumInput struct {
	gorm.Model
	SpotifyURL string `json:"spotifyURL"`
}

// UpdateAlbumInput struct to specify input for updating an album in the database
type UpdateAlbumInput struct {
	gorm.Model
	SpotifyURL string `json:"spotifyURL"`
}
