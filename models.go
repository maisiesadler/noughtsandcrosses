package main

type joinGameRequest struct {
	GameName   string `json:"gameName"`
	PlayerName string `json:"playerName"`
}

type placeRequest struct {
	GameName string `json:"gameName"`
	Counter  string `json:"counter"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}
