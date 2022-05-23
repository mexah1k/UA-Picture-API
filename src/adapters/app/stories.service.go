package app

import (
	"ukraine-picture/src/ports"
)

type StoryService struct {
	storage ports.StoriesStoragePort
}

func NewStoryService(db ports.StoriesStoragePort) *StoryService {
	return &StoryService{storage: db}
}

func (app *StoryService) Create(story *ports.UpsertStory) (uint, error) {
	// validate
	id, err := app.storage.Create(story)
	if err != nil {
		// log error
		return 0, err
	}

	return id, nil
}

func (app *StoryService) Update(id uint, story *ports.UpsertStory) (uint, error) {
	// validate
	id, err := app.storage.Update(id, story)
	if err != nil {
		// log error
		return 0, err
	}

	return id, nil
}

func (app *StoryService) Delete(id uint) error {
	err := app.storage.Delete(id)
	if err != nil {
		// log error
		return err
	}

	return nil
}

func (app *StoryService) Find(id uint) (*ports.Story, error) {
	story, err := app.storage.Find(id)
	if err != nil {
		// log error
		return nil, err
	}

	return story, nil
}

func (app *StoryService) Query(id uint) (*[]ports.Story, error) {
	stories, err := app.storage.Query(id)
	if err != nil {
		// log error
		return nil, err
	}

	return stories, nil
}
