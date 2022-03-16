package services

import (
	//"github.com/danimagb/go-training-restapi/application/dtos"
	"github.com/danimagb/go-training-restapi/domain/entities"
	"github.com/danimagb/go-training-restapi/domain/repositories"
)



type AlbumServiceInterface interface {
	GetAll() (*[]entities.Album, error)
}

type albumService struct{
	AlbumRepository repositories.AlbumRepository
}

func NewAlbumService(ar repositories.AlbumRepository) AlbumServiceInterface  {
	return &albumService{ AlbumRepository: ar}
}


func (as *albumService) GetAll() (*[]entities.Album, error) {
	return as.AlbumRepository.GetAll()
}