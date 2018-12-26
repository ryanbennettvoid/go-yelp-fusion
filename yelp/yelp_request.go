package yelp

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
)

func (client *Client) request(method string, endpoint string, params map[string]interface{}, response interface{}) error {
  url := fmt.Sprintf("https://api.yelp.com/v3%s", endpoint)
  if client.Debug {
    fmt.Printf("%s %s %+v\n", method, url, params)
  }
  httpClient := &http.Client{}
  paramsAsBytes, err := json.Marshal(params)
  if err != nil {
    return err
  }
  req, err := http.NewRequest(method, url, bytes.NewBuffer(paramsAsBytes))
  if err != nil {
    return err
  }
  req.ContentLength = int64(len(paramsAsBytes))
  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.AuthOptions.ApiKey))
  res, err := httpClient.Do(req)
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.StatusCode != http.StatusOK {
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
      return err
    }
    return errors.New(string(data))
  }
  json.NewDecoder(res.Body).Decode(response)
  return nil
}
