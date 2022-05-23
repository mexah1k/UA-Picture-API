package database

import (
	"errors"
	"ukraine-picture/src/adapters/frameworks/database/models"
	"ukraine-picture/src/common"
	"ukraine-picture/src/ports"

	"gorm.io/gorm"
)

type StoriesStorage struct {
	conn *DbConnector
}

func NewStoriesStorage(connection *DbConnector) *StoriesStorage {
	return &StoriesStorage{conn: connection}
}

func (storage *StoriesStorage) findStoryMedia(ids []uint) ([]models.Media, error) {
	var media []models.Media
	if mediaRes := storage.conn.db.Where("id IN ?", ids).Find(&media); mediaRes.Error != nil {
		if errors.Is(mediaRes.Error, gorm.ErrRecordNotFound) {
			return nil, common.ItemNotFoundError
		}

		return nil, common.UnexpectedError
	}

	return media, nil
}

func (storage *StoriesStorage) Create(request *ports.UpsertStory) (uint, error) {
	media, err := storage.findStoryMedia(request.Media)
	if err != nil {
		return 0, err
	}

	story := models.MapUpsertStory(request, media)
	result := storage.conn.db.Create(&story)

	return story.Id, result.Error
}

func (storage *StoriesStorage) Update(id uint, request *ports.UpsertStory) (uint, error) {
	media, err := storage.findStoryMedia(request.Media)
	if err != nil {
		return 0, err
	}

	story, err := storage.Find(id)
	if err != nil {
		return 0, err
	}

	updateStory := models.MapUpsertStory(request, media)
	result := storage.conn.db.Model(&story).Updates(updateStory)

	return story.Id, result.Error
}

func (storage *StoriesStorage) Delete(id uint) error {
	res := storage.conn.db.Delete(&models.Story{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (storage *StoriesStorage) Find(id uint) (*ports.Story, error) {
	var story models.Story
	err := storage.conn.db.First(&story, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ItemNotFoundError
		}

		return nil, common.UnexpectedError
	}

	response := models.MapStoryResponse(&story)

	return &response, nil
}

func (storage *StoriesStorage) Query(id uint) (*[]ports.Story, error) {
	var stories []models.Story
	result := storage.conn.db.Where("id > ?", id).Limit(PAGE_SIZE).Find(&stories)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, common.ItemNotFoundError
		}

		return nil, common.UnexpectedError
	}

	response := make([]ports.Story, len(stories))
	for id, story := range stories {
		response[id] = models.MapStoryResponse(&story)
	}

	return &response, result.Error
}
