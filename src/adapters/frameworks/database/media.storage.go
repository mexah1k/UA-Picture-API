package database

import (
	"errors"
	"ukraine-picture/src/adapters/frameworks/database/models"
	"ukraine-picture/src/common"
	"ukraine-picture/src/ports"

	"gorm.io/gorm"
)

type MediaStorage struct {
	conn *DbConnector
}

func NewMediaStorage(connection *DbConnector) *MediaStorage {
	return &MediaStorage{conn: connection}
}

func (storage *MediaStorage) Create(request *ports.UpsertMedia) (uint, error) {
	media := models.MapUpsertMedia(request)
	result := storage.conn.db.Create(&media)

	return media.Id, result.Error
}

func (storage *MediaStorage) Update(id uint, request *ports.UpsertMedia) (uint, error) {
	media, err := storage.Find(id)
	if err != nil {
		return 0, err
	}

	updateMedia := models.MapUpsertMedia(request)
	result := storage.conn.db.Model(&media).Updates(updateMedia)

	return media.Id, result.Error
}

func (storage *MediaStorage) Delete(id uint) error {
	res := storage.conn.db.Delete(&models.Media{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (storage *MediaStorage) Find(id uint) (*ports.Media, error) {
	var media models.Media
	err := storage.conn.db.First(&media, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ItemNotFoundError
		}

		return nil, common.UnexpectedError
	}

	response := models.MapMediaResponse(&media)

	return &response, nil
}

func (storage *MediaStorage) Query(id uint) (*[]ports.Media, error) {
	var Media []models.Media
	result := storage.conn.db.Where("id > ?", id).Limit(PAGE_SIZE).Find(&Media)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, common.ItemNotFoundError
		}

		return nil, common.UnexpectedError
	}

	response := make([]ports.Media, len(Media))
	for id, media := range Media {
		response[id] = models.MapMediaResponse(&media)
	}

	return &response, result.Error
}
