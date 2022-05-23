package models

import (
	"database/sql"
	"time"
	"ukraine-picture/src/ports"

	"gorm.io/gorm"
)

type Media struct {
	Id        uint `gorm:"primaryKey"`
	Url       string
	Type      string `gorm:"size:50"`
	StoryId   sql.NullInt32
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func MapUpsertMedia(src *ports.UpsertMedia) Media {
	media := Media{
		Url:  src.Url,
		Type: src.Type,
	}

	return media
}

func MapMediaResponse(src *Media) ports.Media {
	media := ports.Media{
		Id:        src.Id,
		Url:       src.Url,
		Type:      src.Type,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		DeletedAt: src.DeletedAt.Time,
	}

	return media
}
