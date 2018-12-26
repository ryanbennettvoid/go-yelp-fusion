package yelp

import (
  "encoding/json"
  "io/ioutil"
)

// AuthOptions contains authentication options for the Yelp API
type AuthOptions struct {
  ApiKey string `json:"api_key"`
}

// AuthOptionsFromJsonFile returns an AuthOptions object populated
// from a JSON file
func AuthOptionsFromJsonFile(configPath string) (AuthOptions, error) {
  authOptions := AuthOptions{}
  bytes, err := ioutil.ReadFile(configPath)
  if err != nil {
    return authOptions, err
  }
  err = json.Unmarshal(bytes, &authOptions)
  if err != nil {
    return authOptions, err
  }
  return authOptions, nil
}
