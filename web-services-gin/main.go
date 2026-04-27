package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Taj Mahal", Artist: "Kara", Price: 9.98},
	{ID: "2", Title: "Blue Train", Artist: "John Coltrane", Price: 8.99},
	{ID: "3", Title: "Jeru", Artist: "Gerry Mulligan", Price: 7.99},
	{ID: "4", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 8.99},
}

func getAlbums(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, albums)
}

func createAlbums(c *gin.Context)  {
    var newAlbum album

	 // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context) {
     id := c.Param("id")

	 for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	 }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.Run()
}
