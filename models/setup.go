package models

import (
	"context"
	"fmt"
	"os"

	"github.com/zmb3/spotify"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"golang.org/x/oauth2/clientcredentials"
)

// DB is the variable storing the gorm ORM functions
var DB *gorm.DB

// SpotifyClient is the variable storing the Spotify Client functions from zmb3/spotify package
var SpotifyClient spotify.Client

// ConnectDatabase Connect to postgresql database with gorm package
func ConnectDatabase() {

	databaseHost := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	databasePort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", databaseHost, username, password, databaseName, databasePort)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Album{})

	DB = database
}

// ConnectSpotifyAPI Connect to the spotify api and retrieve an authenticated client to use for requests to the api using clientcredentials and zmb3/spotify packages
func ConnectSpotifyAPI() {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())

	if err != nil {
		panic("Failed to get Spotify token, error: " + fmt.Sprintf("%v", err))
	}

	SpotifyClient = spotify.Authenticator{}.NewClient(token)
}
