package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: 1, Title: "Bad", Artist: "Michael", Year: 2022},
	{ID: 2, Title: "I need a hero", Artist: "Sherk", Year: 2022},
	{ID: 3, Title: "Le pido a Dios", Artist: "Feid", Year: 2022},
	{ID: 4, Title: "ventilador sonido para dormir", Artist: "Mom", Year: 2022},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)

	if err != nil {
		c.Status(http.StatusNotAcceptable)
		return
	}

	for _, a := range albums {
		if a.ID == idn {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")

}
