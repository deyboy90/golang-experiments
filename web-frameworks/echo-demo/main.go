package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Logger.Fatal(router.Start(":8080"))
}

// getAlbums responds with the list of all albums as JSON.
// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more.
// Despite the similar name, this is different from Go’s built-in context package.
func getAlbums(e echo.Context) error {
	return e.JSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(e echo.Context) error {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := e.Bind(&newAlbum); err != nil {
		return err
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	return e.JSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(e echo.Context) error {
	id := e.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			return e.JSON(http.StatusOK, a)
		}
	}

	return e.JSON(http.StatusNotFound, map[string]interface{}{"message": "album not found"})
}
