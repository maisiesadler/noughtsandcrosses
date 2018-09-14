package game

import "testing"

func Test_placeCounter_validCounter_returnsNil(t *testing.T) {
	g := newGame()
	err := g.placeCounter("X", 0, 0)
	if err != nil {
		t.Error("Expected nil respose")
	}
}

func Test_placeCounter_invalidCounter_returnsError(t *testing.T) {
	g := newGame()
	err := g.placeCounter("invalid", 0, 0)
	if err == nil {
		t.Error("Counter was not valid, expected error")
	}
}

func Test_placeCounter_invalidRange_returnsError(t *testing.T) {
	g := newGame()
	err := g.placeCounter("X", 7, 20)
	if err == nil {
		t.Error("Counter was placed out of range, expected error")
	}
}

func Test_placeCounter_counterPlacedOnNonEmptySpot_returnsError(t *testing.T) {
	g := newGame()
	g.placeCounter("X", 0, 0)
	err := g.placeCounter("O", 0, 0)
	if err == nil {
		t.Error("Counter was placed on non empty spot, expected error")
	}
}

func Test_placeCounter_counterPlacedBeforeTurn_returnsError(t *testing.T) {
	g := newGame()
	err := g.placeCounter("O", 0, 0)
	if err == nil {
		t.Error("Was not counters turn, expected error")
	}
}

func Test_placeCounter_counterAttemptsTwoTurns_returnsError(t *testing.T) {
	g := newGame()
	g.placeCounter("X", 0, 0)
	err := g.placeCounter("X", 0, 1)
	if err == nil {
		t.Error("Counter places before rurn, expected error")
	}
}
