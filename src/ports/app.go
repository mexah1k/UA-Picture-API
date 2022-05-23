package ports

type StoriesPort interface {
	Create(story *UpsertStory) (uint, error)
	Update(id uint, story *UpsertStory) (uint, error)
	Delete(id uint) error
	Find(id uint) (*Story, error)
	Query(id uint) (*[]Story, error)
}

type MediaPort interface {
	Create(media *UpsertMedia) (uint, error)
	Update(id uint, story *UpsertMedia) (uint, error)
	Delete(id uint) error
	Find(id uint) (*Media, error)
	Query(id uint) (*[]Media, error)
}
