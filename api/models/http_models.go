package models

// TODO: Name them better
// Receipt Process endpoint Response struct
type ProcessResponse struct {
	Id string `json:"id"`
}

// Receipt Points calculation endpoint Response struct
type PointsResponse struct {
	Points string `json:"points"`
}
