package usecases

import (
	"time"

	"github.com/noskine/pilot_api/internal/models"
	"github.com/noskine/pilot_api/internal/repository"
	"github.com/noskine/pilot_api/pkg/dto"
)

type PublicationUsecase struct {
	Entity *models.Publications
}

func NewPublicationsUseCase() *PublicationUsecase {
	return &PublicationUsecase{
		Entity: new(models.Publications),
	}
}

func (pu *PublicationUsecase) CreatePublicationsUseCase(dtos *dto.DTORequestPepplos) (*[]models.Publications, error) {
	connect := repository.Connecting()

	pu.Entity = &models.Publications{
		Title:     dtos.Title,
		Text:      dtos.Text,
		Author:    dtos.Author,
		Create_at: time.Now(),
	}

	if err := connect.Create(pu.Entity); err != nil {
		return nil, err
	}

	result, err := connect.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}
