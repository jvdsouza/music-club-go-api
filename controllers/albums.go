package controllers

import (
	"errors"
	"music_club/helpers"
	"music_club/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
)

// GET /albums
// Get all albums
func GetAllAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

// POST /albums
// Create new album
func AddAlbum(c *gin.Context) {
	var input models.CreateAlbumInput
	// https://open.spotify.com/album/0EMbzFBRoIt0fmTsowZ8Zv?si=hX3wFj0-TTOPySlT-909jQ
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	spotifyAlbum, err := getSpotifyAlbumData(input.SpotifyURL)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Create album
	album := models.Album{
		SpotifyURL:   input.SpotifyURL,
		AlbumTitle:   spotifyAlbum.Name,
		AlbumArtists: spotifyAlbum.Artists[0].Name,
		AlbumArt:     spotifyAlbum.Images[0].URL}

	models.DB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// PATCH /albums/:id
// Update an album
func UpdateAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.First(&album, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found"})
		return
	}

	var input models.UpdateAlbumInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	spotifyAlbum, err := getSpotifyAlbumData(input.SpotifyURL)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedAlbum := models.Album{
		SpotifyURL:   input.SpotifyURL,
		AlbumTitle:   spotifyAlbum.Name,
		AlbumArtists: spotifyAlbum.Artists[0].Name,
		AlbumArt:     spotifyAlbum.Images[0].URL}

	models.DB.Model(&album).Updates(updatedAlbum)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// DELETE /albums
// Delete an album
func DeleteAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.First(&album, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found"})
		return
	}

	models.DB.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func getSpotifyAlbumData(spotifyURL string) (*spotify.FullAlbum, error) {
	// Check if input is a spotify album url
	isSpotifyAlbumURL, err := regexp.MatchString(`open\.spotify\.com\/album\/`, spotifyURL)

	if err != nil {
		return nil, errors.New(err.Error())
	} else if !isSpotifyAlbumURL {
		return nil, errors.New("URL is not of \"open.spotify.com/album\" format")
	}

	// Grab the album ID
	albumID, err := helpers.GetAlbumIDFromURL(spotifyURL)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	spotifyAlbum, err := models.SpotifyClient.GetAlbum(spotify.ID(albumID))

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return spotifyAlbum, nil
}
