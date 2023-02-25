package stats

import (
	"github.com/fraqtop/footballcore/competition"
	"github.com/fraqtop/footballcore/team"
	"testing"
)

func TestValid(t *testing.T) {
	entity := New(
		team.New(1, "qwe", "qwerty"),
		competition.New(1, "qwe"),
		"2021-2022",
		5,
		5,
		0,
		5,
		0,
		0,
		0,
	)
	if !entity.IsValid() {
		t.Fail()
	}
}

func TestInvalid(t *testing.T) {
	invalidEntities := []Stats{
		{},
		New(
			team.Team{},
			competition.New(1, "qwe"),
			"2021-2022",
			5,
			5,
			0,
			5,
			0,
			0,
			0,
		),
		New(
			team.New(1, "qwe", "qwerty"),
			competition.Competition{},
			"2021-2022",
			5,
			5,
			0,
			5,
			0,
			0,
			0,
		),
		New(
			team.New(1, "qwe", "qwerty"),
			competition.New(1, "qwerty"),
			"",
			5,
			5,
			0,
			5,
			0,
			5,
			0,
		),
		New(
			team.New(1, "qwe", "qwerty"),
			competition.New(1, "qwerty"),
			"2021-2022",
			-1,
			5,
			-1,
			-1,
			-1,
			-1,
			-1,
		),
	}

	for _, invalidEntity := range invalidEntities {
		if invalidEntity.IsValid() {
			t.Fail()
		}
	}
}
