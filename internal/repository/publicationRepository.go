package repository

import (
	"github.com/noskine/pilot_api/config"
	"github.com/noskine/pilot_api/internal/models"
	"gorm.io/gorm"
)

type PublicationRepository struct {
	Db *gorm.DB
}

func Connecting() *PublicationRepository {
	return &PublicationRepository{
		Db: config.NewConnection().Db,
	}
}

func (pr *PublicationRepository) Create(entity *models.Publications) error {
	pr.Db.Create(&entity)
	return nil
}

func (pr *PublicationRepository) FindAll() (*[]models.Publications, error) {
	var pub *[]models.Publications

	if result := pr.Db.Find(&pub); result.Error != nil {
		return nil, result.Error
	}

	return pub, nil
}
