package models

type VideoInstance struct {
	*Model
	InstId       string
	Title        string
	Info         string
	Status       VideoStatus
	Actress      string
	CoverPicture string
}

type VideoStatus int

const (
	Finished VideoStatus = iota
	Continued
)

type EpisodesInstance struct {
	*Model
	InstId          string
	EpsId           int // 集数
	PlaylistAddress string
}
