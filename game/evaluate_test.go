package game

import "testing"

func Test_combinationIsWinning_newGame_combinationIsNotWinning(t *testing.T) {
	g := newGame()
	isWinning, _ := g.combinationIsWinning([]int{1, 2, 3})
	if isWinning {
		t.Error("new game has winning combination!")
	}
}

func Test_combinationIsWinning_countersPlaced_combinationIsNotWinning(t *testing.T) {
	g := newGame()
	g.placeCounter("X", 0, 0)
	g.placeCounter("O", 1, 0)
	g.placeCounter("X", 1, 0)
	isWinning, _ := g.combinationIsWinning([]int{1, 2, 3})
	if isWinning {
		t.Error("expected no winning combinations, but found some!")
	}
}

func Test_combinationIsWinning_countersPlaced_combinationIsWinning(t *testing.T) {
	g := newGame()
	g.placeCounter("X", 0, 0)
	g.placeCounter("O", 1, 0)
	g.placeCounter("X", 0, 1)
	g.placeCounter("O", 1, 1)
	g.placeCounter("X", 0, 2)
	isWinning, _ := g.combinationIsWinning([]int{0, 1, 2})
	if !isWinning {
		t.Error("expected winning combination but founfd none!")
	}
}
