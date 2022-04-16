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

type Result struct {
	result_sub   int
	result_mult  int
	result_divid int
}

var albums = []album{

	{ID: "1", Title: "blue train", Artist: "john coltrane", Price: 56.99},
	{ID: "2", Title: "jeru", Artist: "gerry mulligan", Price: 17.99},
	{ID: "3", Title: "sarah vaughan and clifford brown", Artist: "sarah vaughan", Price: 39.99},
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.GET("add/:a/:b", addition)
	router.POST("/sub", substraction)
	router.POST("/multiply", multiplication)
	router.POST("/divide", devision)

	router.POST("/result/all", substraction, multiplication, devision)
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

// addition of two numbers

func addition(c *gin.Context) {

	a, err := strconv.Atoi(c.Param("a"))
	if err != nil {

		log.Println(err)
	}

	b, err := strconv.Atoi(c.Param("b"))

	if err != nil {

		log.Println(err)
	}
	sum := a + b
	fmt.Println(sum)
	c.JSON(http.StatusOK, sum)
}

//substraction of two numbers

func substraction(c *gin.Context) {

	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {

		log.Println(err)
	}
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {

		log.Println(err)
	}

	sub := a - b

	c.JSON(http.StatusOK, sub)

}

// multiplication of two numbers

func multiplication(c *gin.Context) {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		log.Println(err)
	}
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		log.Println(err)
	}
	multi := a * b
	c.JSON(http.StatusOK, multi)

}

// division of two numbers

func devision(c *gin.Context) {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		log.Println(err)
	}
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		log.Println(err)
	}
	divid := float64(a) / float64(b)
	c.JSON(http.StatusOK, divid)

}
