package game

func (g *game) string() string {
	return string(g.status) + "\n" + g.viewBoard()
}

func (g *game) viewBoard() string {
	s := ""
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			cell := " " + g.board[x][y] + " "
			s += cell
			if x != 2 {
				s += "|"
			}
		}
		if y != 2 {
			s += "\n---+---+---\n"
		}
	}
	return s
}
