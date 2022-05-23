package models

import (
	"time"
	"ukraine-picture/src/ports"

	"gorm.io/gorm"
)

type Story struct {
	Id        uint `gorm:"primaryKey"`
	Media     []Media
	TitleUA   string
	TitleEN   string
	TextUA    string
	TextEN    string
	Date      time.Time
	Place     string         `gorm:"size:255;index:idx_place"`
	Source    string         `gorm:"size:255"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func MapUpsertStory(src *ports.UpsertStory, media []Media) Story {
	story := Story{
		Media:   media,
		TitleUA: src.TitleUA,
		TitleEN: src.TitleEN,
		TextUA:  src.TextUA,
		TextEN:  src.TextEN,
		Source:  src.Source,
		Date:    time.Unix(src.Date, 0),
		Place:   src.Place,
	}

	return story
}

func MapStoryResponse(src *Story) ports.Story {
	mediaResponse := make([]ports.Media, len(src.Media))
	for id, story := range src.Media {
		mediaResponse[id] = MapMediaResponse(&story)
	}

	story := ports.Story{
		Id:        src.Id,
		Media:     mediaResponse,
		TitleUA:   src.TitleUA,
		TitleEN:   src.TitleEN,
		TextUA:    src.TextUA,
		TextEN:    src.TextEN,
		Source:    src.Source,
		Date:      src.Date.Unix(),
		Place:     src.Place,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		DeletedAt: src.DeletedAt.Time,
	}

	return story
}
