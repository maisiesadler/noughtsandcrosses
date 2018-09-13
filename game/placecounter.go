package game

import "errors"

func (g *game) placeCounter(counter string, positionX int, positionY int) error {

	if err := g.validateCounter(counter); err != nil {
		return err
	}

	if err := g.validateRange(positionX, positionY); err != nil {
		return err
	}

	if cell := g.board[positionX][positionY]; cell != emptyCell {
		return errors.New("cell already taken")
	}

	g.board[positionX][positionY] = counter
	g.switchTurn()

	return nil
}

func (g *game) validateCounter(counter string) error {
	if counter != "X" && counter != "O" {
		return errors.New("counter not valid")
	}

	if counter == "X" && g.status != turnX {
		return errors.New("It is not X's turn")
	}

	if counter == "O" && g.status != turnO {
		return errors.New("It is not O's turn")
	}

	return nil
}

func (g *game) switchTurn() {
	if g.status == turnX {
		g.status = turnO
	} else {
		g.status = turnX
	}
}

func (g *game) validateRange(positionX int, positionY int) error {
	if positionX >= 3 {
		return errors.New("x out of range")
	}
	if positionY >= 3 {
		return errors.New("y out of range")
	}
	return nil
}
