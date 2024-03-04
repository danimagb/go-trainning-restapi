package main
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
