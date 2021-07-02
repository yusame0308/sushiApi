package repository

type SushiData struct {
	*BaseRepository
}

func NewSushiData(baseRepository *BaseRepository) *SushiData {
	return &SushiData{BaseRepository: baseRepository}
}
