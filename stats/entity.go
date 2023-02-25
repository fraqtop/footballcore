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

func (this Stats) Wins() int {
	return this.core.Wins
}

func (this Stats) Draws() int {
	return this.core.Draws
}

func (this Stats) Losses() int {
	return this.core.Losses
}

func (this Stats) Team() team.Team {
	return this.core.Team
}

func (this Stats) Competition() competition.Competition {
	return this.core.Competition
}

func (this Stats) Season() string {
	return this.core.Season
}

func (this Stats) Points() int {
	return this.core.Points
}

func (this Stats) Scored() int {
	return this.core.Scored
}

func (this Stats) Passed() int {
	return this.core.Passed
}

func (this Stats) Games() int {
	return this.core.Games
}

func (this Stats) IsValid() bool {
	return this.core.Team.IsValid() &&
		this.core.Competition.IsValid() &&
		len(this.core.Season) > 0 &&
		this.core.Games >= 0 &&
		this.core.Passed >= 0 &&
		this.core.Scored >= 0 &&
		this.core.Losses >= 0 &&
		this.core.Draws >= 0 &&
		this.core.Wins >= 0 &&
		this.core.Points >= 0
}
