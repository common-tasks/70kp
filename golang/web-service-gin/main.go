package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Jharkhand top ten", Artist: "hero khalkho", Price: 200},
	{ID: "2", Title: "iskool ke tame pe", Artist: "jheel rani", Price: 20},
	{ID: "3", Title: "jhar bros", Artist: "pan kumar", Price: 20.25},
}

func main() {
	fmt.Println("rest apis")
	handleRestCalls()

}

func handleRestCalls() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/ping", handlePing)
	router.GET("/albums/:id",getAlbumByID)
	router.Run("localhost:8080")
}
func handlePing(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func getAlbums(c *gin.Context) {
	fmt.Println("getAlbums called")
	c.IndentedJSON(http.StatusOK, albums)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range albums {
		if(v.ID==id){
			c.IndentedJSON(http.StatusOK,v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound,gin.H{
		"message":"album not found",
	})
}
func postAlbums(c *gin.Context) {
	fmt.Println("post album request received")
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
