package controllers

import "ukraine-picture/src/ports"

// stories
type UpsertStoryRequest struct {
	MediaIds []uint `json:"mediaIds" binding:"required"`
	TitleUA  string `json:"titleUA"`
	TitleEN  string `json:"titleEN"`
	TextUA   string `json:"textUA"`
	TextEN   string `json:"textEN"`
	Source   string `json:"source"`
	Date     int64  `json:"date"`
	Place    string `json:"place"`
}

type StoryResponse struct {
	Id        uint     `json:"id"`
	MediaUrls []string `json:"media"`
	TitleUA   string   `json:"titleUA"`
	TitleEN   string   `json:"titleEN"`
	TextUA    string   `json:"textUA"`
	TextEN    string   `json:"textEN"`
	Source    string   `json:"source"`
	Date      int64    `json:"date"`
	Place     string   `json:"place"`
}

func MapUpsertStoryRequest(src *UpsertStoryRequest) ports.UpsertStory {
	story := ports.UpsertStory{
		Media:   src.MediaIds,
		TitleUA: src.TitleUA,
		TitleEN: src.TitleEN,
		TextUA:  src.TextUA,
		TextEN:  src.TextEN,
		Source:  src.Source,
		Date:    src.Date,
		Place:   src.Place,
	}

	return story
}

func MapStoryResponse(src *ports.Story) StoryResponse {
	var urls []string
	for _, curMedia := range src.Media {
		urls = append(urls, curMedia.Url)
	}

	story := StoryResponse{
		Id:        src.Id,
		MediaUrls: urls,
		TitleUA:   src.TitleUA,
		TitleEN:   src.TitleEN,
		TextUA:    src.TextUA,
		TextEN:    src.TextEN,
		Source:    src.Source,
		Date:      src.Date,
		Place:     src.Place,
	}

	return story
}

// media
type UpsertMediaRequest struct {
	Url  string `json:"url" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type MediaResponse struct {
	Id   uint   `json:"id"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

func MapUpsertMediaRequest(src *UpsertMediaRequest) ports.UpsertMedia {
	media := ports.UpsertMedia{
		Url:  src.Url,
		Type: src.Type,
	}

	return media
}

func MapMediaResponse(src *ports.Media) MediaResponse {
	media := MediaResponse{
		Id:   src.Id,
		Url:  src.Url,
		Type: src.Type,
	}

	return media
}
