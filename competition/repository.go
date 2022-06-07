package competition

type ReadRepository interface {
	All() []Competition
}

type WriteRepository interface {
	Save(competition Competition) error
}
