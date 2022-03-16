package repositories

import "github.com/danimagb/go-training-restapi/domain/entities"

type AlbumRepository interface {
	//Save(*entities.Album) (*entities.Food, map[string]string)
	//Get(uint64) (*entities.Album, error)
	GetAll() (*[]entities.Album, error)
	//Update(*entities.Album) (*entities.Album, map[string]string)
	//Delete(uint64) error
}