// Package yelp provides a client for the Yelp Fusion (v3) REST API
package yelp

// Client provides methods to perform requests on the Yelp API
type Client struct {
  AuthOptions AuthOptions
  Debug       bool
}

// NewClient creates a new Client instance
func NewClient(authOptions AuthOptions) Client {
  return Client{
    AuthOptions: authOptions,
  }
}
