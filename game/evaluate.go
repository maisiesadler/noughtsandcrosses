package game

import (
	"math"
)

var winningCombinations = [][]int{
	[]int{0, 1, 2},
	[]int{3, 4, 5},
	[]int{6, 7, 8},
	[]int{0, 3, 6},
	[]int{1, 4, 7},
	[]int{2, 5, 8},
	[]int{0, 4, 8},
	[]int{2, 4, 6},
}

func (g *game) evaluateStatus() {
	if someoneHasWon, winner := g.checkForWinner(); someoneHasWon {
		if winner == "X" {
			g.status = resultX
		} else if winner == "O" {
			g.status = resultO
		}
	}
}

func (g *game) checkForWinner() (bool, string) {
	for _, combination := range winningCombinations {
		if someoneHasWon, winner := g.combinationIsWinning(combination); someoneHasWon {
			return true, winner
		}
	}
	return false, ""
}

func (g *game) combinationIsWinning(combination []int) (bool, string) {

	var combinationType string

	for _, coord := range combination {
		x := int(math.Floor(float64(coord) / 3))
		y := int(math.Mod(float64(coord), 3))
		cell := g.board[x][y]
		if cell == emptyCell {
			return false, ""
		}
		if combinationType == "" {
			combinationType = cell
		}
		if cell != combinationType {
			return false, ""
		}
	}

	return true, combinationType
}
