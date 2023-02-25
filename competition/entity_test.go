package competition

import (
	"testing"
)

func TestValid(t *testing.T) {
	if entity := New(1, "qwe"); !entity.IsValid() {
		t.Fail()
	}
}

func TestInvalid(t *testing.T) {
	entity := Competition{}
	if entity.IsValid() {
		t.Fail()
	}
}
