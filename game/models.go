package game

// Game is the current in play game
type game struct {
	status gameStatus
	board  *[3][3]string
}
type gameStatus string

const (
	turnX      gameStatus = "X Turn"
	turnO      gameStatus = "O Turn"
	resultX    gameStatus = "X Has Won"
	resultO    gameStatus = "O Has Won"
	resultDraw gameStatus = "Draw!"
)

type coordinates struct {
	x int
	y int
}
