package main

// var games = make(map[string]*game)

// func createOrGetSingleGame() *game {
// 	if _, exists := games[""]; !exists {
// 		games[""] = newGame()
// 	}
// 	return games[""]
// }

// // func createGame() bool {
// // 	game := newGame()
// // 	games[""] = game
// // 	return true
// // }

// // func getGame(gameName string) *game {
// // 	if game, exists := games[gameName]; exists {
// // 		return game
// // 	}
// // 	return nil
// // }

// // func joinGame(request joinGameRequest) bool {
// // 	if game, exists := games[request.GameName]; exists {
// // 		game.player2 = request.PlayerName
// // 		game.status = player1sTurn
// // 		return true
// // 	}
// // 	return false
// // }

// func viewgames() string {
// 	s := ""
// 	for key, value := range games {
// 		s += key + "-" + value.player1 + " vs " + value.player2 + ": " + string(value.status) + "\n"
// 		s += value.viewBoard()
// 	}
// 	return s
// }
