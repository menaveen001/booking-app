package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{

	{ID: "1", Title: "blue train", Artist: "john coltrane", Price: 56.99},
	{ID: "2", Title: "jeru", Artist: "gerry mulligan", Price: 17.99},
	{ID: "3", Title: "sarah vaughan and clifford brown", Artist: "sarah vaughan", Price: 39.99},
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	//router.GET("/albums/:ID", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.GET("/albums/:a/:b", addition)
	router.Run("localhost:8080")

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("ID")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func addition(c *gin.Context) {

	a, err := strconv.Atoi(c.Param("a"))
	if err != nil {

		log.Println(err)
	}

	b, err := strconv.Atoi(c.Param("b"))

	if err != nil {

		log.Println(err)
	}
	//a := c.Param("strconv.Atoi(a)")
	//b := c.Param("strconv.Atoi(b)")
	sum := a + b
	fmt.Println(sum)
	c.JSON(http.StatusOK, sum)
}
