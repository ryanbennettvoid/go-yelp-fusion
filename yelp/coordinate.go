
package yelp

// Coordinate represents a coordinate object from the Yelp API
type Coordinate struct {
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
}