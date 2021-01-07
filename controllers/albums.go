package controllers

import (
	"music_club/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create album
	album := models.Album{SpotifyURL: input.SpotifyURL}
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

	models.DB.Model(&album).Update("SpotifyURL", input)

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
