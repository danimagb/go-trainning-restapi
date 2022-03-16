package persistence

import (
	"github.com/danimagb/go-training-restapi/domain/entities"
	"github.com/jinzhu/gorm"
)

type AlbumRepository struct{
	db * gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{db}
}


func (r *AlbumRepository) GetAll() (*[]entities.Album, error) {
	var err error
	albums := []entities.Album{}

	err = r.db.Debug().Model(&entities.Album{}).Limit(100).Find(&albums).Error
	if err != nil {
		return &[]entities.Album{}, err
	}

	return &albums, nil
}
