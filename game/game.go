package game

var emptyCell = " "

func newGame() *game {
	g := &game{status: turnX, board: &[3][3]string{}}
	g.resetBoard()
	return g
}

var currentGame *game

func getCurrentGame() *game {
	if currentGame == nil {
		currentGame = newGame()
	}
	return currentGame
}

// PlaceCounter allows clients to place a counter on the board
func PlaceCounter(counter string, positionX int, positionY int) error {
	g := getCurrentGame()
	err := g.placeCounter(counter, positionX, positionY)
	if err != nil {
		return err
	}
	g.evaluateStatus()
	return nil
}

// String prints current game as a string
func String() string {
	g := getCurrentGame()
	return g.string()
}

// Reset resets the board
func Reset() {
	g := getCurrentGame()
	g.resetBoard()
}

func (g *game) resetBoard() {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			g.board[x][y] = emptyCell
		}
	}
	g.status = turnX
}
