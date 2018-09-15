package server

type placeRequest struct {
	Counter string `json:"counter"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}
