package models

type BaseFileInstance struct {
	FileAddress string
	FileInfo    string
}

type VideoInstance struct {
	*Model
	*BaseFileInstance
	InstId       string
	Title        string
	Info         string
	Status       VideoStatus
	Actress      string
	CoverPicture string // ImgInstance
}

type VideoStatus int

const (
	Finished VideoStatus = iota
	Continued
)

// VideoInstance -> Muti EpisodesInstance
type EpisodesInstance struct {
	*Model
	*BaseFileInstance
	InstId          string
	EpsId           int // 集数
	PlaylistAddress string
}

type ImgInstances struct {
	*BaseFileInstance
	Name   string
	InstId string
}
