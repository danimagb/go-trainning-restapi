package controllers

import (
	"net/http"
	"github.com/danimagb/go-training-restapi/application/services"
	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	AlbumService services.AlbumServiceInterface
}

func NewAlbumController(ac *services.AlbumServiceInterface) AlbumController {
	return AlbumController{ AlbumService: *ac }
}

func (controller *AlbumController) Route(router *gin.Engine){
	router.GET("/albums", controller.getAlbums)
}


func (controller *AlbumController) getAlbums(c *gin.Context) {
	// album := models.Album{}

	// albums, err := album.FindAllAlbums(server.DB)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	albums, err := controller.AlbumService.GetAll();
	if(err != nil){
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, albums)
}
