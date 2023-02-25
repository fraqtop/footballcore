package team

import "testing"

func TestValid(t *testing.T) {
	if entity := New(1, "qwerty", "qwe"); !entity.IsValid() {
		t.Fail()
	}
}

func TestInvalid(t *testing.T) {
	empty, partlyFilled := Team{}, New(1, "", "qwe")
	if empty.IsValid() || partlyFilled.IsValid() {
		t.Fail()
	}
}
