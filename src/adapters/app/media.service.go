package app

import (
	"ukraine-picture/src/ports"
)

type MediaService struct {
	storage ports.MediaStoragePort
}

func NewMediaService(db ports.MediaStoragePort) *MediaService {
	return &MediaService{storage: db}
}

func (app *MediaService) Create(Media *ports.UpsertMedia) (uint, error) {
	// validate
	id, err := app.storage.Create(Media)
	if err != nil {
		// log error
		return 0, err
	}

	return id, nil
}

func (app *MediaService) Update(id uint, Media *ports.UpsertMedia) (uint, error) {
	// validate
	id, err := app.storage.Update(id, Media)
	if err != nil {
		// log error
		return 0, err
	}

	return id, nil
}

func (app *MediaService) Delete(id uint) error {
	err := app.storage.Delete(id)
	if err != nil {
		// log error
		return err
	}

	return nil
}

func (app *MediaService) Find(id uint) (*ports.Media, error) {
	Media, err := app.storage.Find(id)
	if err != nil {
		// log error
		return nil, err
	}

	return Media, nil
}

func (app *MediaService) Query(id uint) (*[]ports.Media, error) {
	stories, err := app.storage.Query(id)
	if err != nil {
		// log error
		return nil, err
	}

	return stories, nil
}
