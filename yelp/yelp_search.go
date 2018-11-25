
package yelp

import (
  "fmt"
)

// SearchOptions defines parameters for the Search method
type SearchOptions struct {
  Term string
  Location string
  Limit int
}

// SearchResponse represents the response returned from the Yelp API
type SearchResponse struct {
  Total int `json:"total"`
  Businesses []Business `json:"businesses"`
  Region Region `json:"region"`
}

// Search sends a business search request to the Yelp API
//
// See https://www.yelp.com/developers/documentation/v3/business_search
func ( client *Client ) Search( request SearchOptions ) ( SearchResponse, error ) {
  method := "GET"
  endpoint := fmt.Sprintf( 
    "/businesses/search?term=%s&location=%s&limit=%d", 
    request.Term, 
    request.Location, 
    request.Limit,
  )
  params := make( map[string]interface{} )
  response := SearchResponse{}
  err := client.request( method, endpoint, params, &response )
  if err != nil {
    return response, err
  }
  return response, nil
}