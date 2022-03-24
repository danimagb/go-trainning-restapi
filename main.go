package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/albums", getAlbums)
// 	router.GET("/albums/:id", getAlbumById)
// 	router.POST("/albums", postAlbums)

// 	router.Run("localhost:8080")
// }

// func getAlbums(c *gin.Context) {
// 	c.JSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	albums = append(albums, newAlbum)
// 	c.JSON(http.StatusCreated, newAlbum)
// }

// func getAlbumById(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, album := range albums {
// 		if album.ID == id {
// 			c.JSON(http.StatusOK, album)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }


import (
	"github.com/danimagb/go-training-restapi/presentation/controllers"
	"github.com/danimagb/go-training-restapi/infrastructure/persistence"
	"github.com/danimagb/go-training-restapi/application/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	router := gin.Default()

	//Get Env variables
	serverPort := os.Getenv("SERVER_PORT")
	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//Setup AlbumController
	repos, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer repos.Close()
	//repos.Automigrate()

	albumService := services.NewAlbumService(repos.Album);
	albumController := controllers.NewAlbumController(&albumService)
	albumController.Route(router)

	router.Run(serverPort)
}
