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
	{ID: "2", Title: "iskool ke tame pe", Artist: "jheel rani", Price: 200},
	{ID: "1", Title: "jhar bros", Artist: "pan kumar", Price: 200},
}

func main() {
	fmt.Println("rest apis")
	handleRestCalls()

}

func handleRestCalls() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	fmt.Println("getAlbums called")
	c.IndentedJSON(http.StatusOK, albums)
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
