package yelp

// Region represents a region object from the Yelp API
type Region struct {
  Center Coordinate `json:"center"`
}
