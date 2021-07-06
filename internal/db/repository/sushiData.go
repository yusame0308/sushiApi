package repository

type SushiData struct {
	*BaseRepository
}

func NewSushiData(baseRepository *BaseRepository) SushiDataInterface {
	return &SushiData{BaseRepository: baseRepository}
}
