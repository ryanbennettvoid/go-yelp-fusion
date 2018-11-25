
package yelp

// Category represents a category object from the Yelp API
type Category struct {
  Alias string `json:"alias"`
  Title string `json:"title"`
}