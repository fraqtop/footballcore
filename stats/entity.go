package stats

import (
	"github.com/fraqtop/footballcore/competition"
	"github.com/fraqtop/footballcore/team"
)

type Stats interface {
	Team() team.Team
	Competition() competition.Competition
	Season() string
	Games() int
	Points() int
	Wins() int
	Draws() int
	Losses() int
	Scored() int
	Passed() int
}

type stats struct {
	team        team.Team
	competition competition.Competition
	season      string
	games       int
	points      int
	wins        int
	draws       int
	losses      int
	scored      int
	passed      int
}

func (s stats) Wins() int {
	return s.wins
}

func (s stats) Draws() int {
	return s.draws
}

func (s stats) Losses() int {
	return s.losses
}

func (s stats) Team() team.Team {
	return s.team
}

func (s stats) Competition() competition.Competition {
	return s.competition
}

func (s stats) Season() string {
	return s.season
}

func (s stats) Points() int {
	return s.points
}

func (s stats) Scored() int {
	return s.scored
}

func (s stats) Passed() int {
	return s.passed
}

func (s stats) Games() int {
	return s.games
}

func New(
	team team.Team,
	competition competition.Competition,
	season string,
	games,
	points,
	wins,
	draws,
	losses,
	scored,
	passed int) Stats {
	return &stats{
		team:        team,
		competition: competition,
		season:      season,
		games:       games,
		points:      points,
		wins:        wins,
		draws:       draws,
		losses:      losses,
		scored:      scored,
		passed:      passed,
	}
}
