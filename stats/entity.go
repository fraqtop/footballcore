package stats

import (
	"github.com/fraqtop/footballcore/competition"
	"github.com/fraqtop/footballcore/team"
)

type core struct {
	Team        team.Team
	Competition competition.Competition
	Season      string
	Games       int
	Points      int
	Wins        int
	Draws       int
	Losses      int
	Scored      int
	Passed      int
}

type Stats struct {
	core
}

func (s Stats) Wins() int {
	return s.core.Wins
}

func (s Stats) Draws() int {
	return s.core.Draws
}

func (s Stats) Losses() int {
	return s.core.Losses
}

func (s Stats) Team() team.Team {
	return s.core.Team
}

func (s Stats) Competition() competition.Competition {
	return s.core.Competition
}

func (s Stats) Season() string {
	return s.core.Season
}

func (s Stats) Points() int {
	return s.core.Points
}

func (s Stats) Scored() int {
	return s.core.Scored
}

func (s Stats) Passed() int {
	return s.core.Passed
}

func (s Stats) Games() int {
	return s.core.Games
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
	return Stats{
		core{
			Team:        team,
			Competition: competition,
			Season:      season,
			Games:       games,
			Points:      points,
			Wins:        wins,
			Draws:       draws,
			Losses:      losses,
			Scored:      scored,
			Passed:      passed,
		},
	}
}
