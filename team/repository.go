package team

import "github.com/fraqtop/footballcore/competition"

type WriteRepository interface {
	BatchUpdate([]Team) error
}

type ReadRepository interface {
	ByCompetition(competition competition.Competition) []Team
}
