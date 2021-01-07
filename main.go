package main

import (
	"net/http"

	"music_club/controllers"
	"music_club/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load(".env")

	if err != nil {
		panic(".env file could not be loaded")
	}

	models.ConnectDatabase()

	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/albums", controllers.GetAllAlbums)
	r.POST("/albums", controllers.AddAlbum)
	r.PATCH("/albums/:id", controllers.UpdateAlbum)
	r.DELETE("/albums/:id", controllers.DeleteAlbum)

	r.Run()
}
