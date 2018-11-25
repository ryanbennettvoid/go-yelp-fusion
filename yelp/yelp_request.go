
package yelp

import (
  "errors"
  "fmt"
  "bytes"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

// ErrorResponse represents an error object from the Yelp API
type ErrorResponse struct {
  Error struct {
    Code string `json:"code"`
    Description string `json:"description"`
  } `json:"error"`
}

func ( client *Client ) request( method string, endpoint string, params map[string]interface{}, response interface{} ) error {
  url := fmt.Sprintf( "https://api.yelp.com/v3%s", endpoint )
  httpClient := &http.Client {}
  paramsAsBytes, err := json.Marshal( params )
  if err != nil {
    return err
  }
  req, err := http.NewRequest( method, url, bytes.NewBuffer( paramsAsBytes ) )
  if err != nil {
    return err
  }
  req.ContentLength = int64( len( paramsAsBytes ) )
  req.Header.Set( "Authorization", "Bearer " + client.AuthOptions.ApiKey )
  res, err := httpClient.Do( req )
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.StatusCode != http.StatusOK {
    bytes, err := ioutil.ReadAll( res.Body )
    if err != nil {
      return err
    }
    errorResponse := ErrorResponse{}
    err = json.Unmarshal( bytes, &errorResponse )
    if err != nil {
      return errors.New( string( bytes ) )
    }
    return errors.New( errorResponse.Error.Description )
  }
  json.NewDecoder( res.Body ).Decode( response )
  return nil
}