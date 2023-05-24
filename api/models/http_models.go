package models

// TODO: Name them better
// Receipt process endpoint's Response struct
type ProcessResponse struct {
	Id string `json:"id"`
}

// Receipt points endpoint's Response struct
type PointsResponse struct {
	Points string `json:"points"`
}
