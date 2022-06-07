package stats

import "github.com/fraqtop/footballcore/competition"

type ReadRepository interface {
	ByCompetition(competition competition.Competition) []Stats
}

type WriteRepository interface {
	BatchUpdate([]Stats) error
}
