package ports

import "time"

// stories
type Story struct {
	Id        uint
	Media     []Media
	TitleUA   string
	TitleEN   string
	TextUA    string
	TextEN    string
	Date      int64
	Place     string
	Source    string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}

type UpsertStory struct {
	Media   []uint
	TitleUA string
	TitleEN string
	TextUA  string
	TextEN  string
	Source  string
	Date    int64
	Place   string
}

// media
type Media struct {
	Id        uint
	Url       string
	Type      string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}

type UpsertMedia struct {
	Url  string
	Type string
}
