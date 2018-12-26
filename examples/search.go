package main

import (
  "fmt"

  "github.com/ryanbennettvoid/go-yelp-fusion/yelp"
)

func main() {
  authOptions, err := yelp.AuthOptionsFromJsonFile("../config.json")
  if err != nil {
    panic(err)
  }
  client := yelp.NewClient(authOptions)
  client.Debug = true
  results, err := client.Search(yelp.SearchOptions{
    Term:     "bar",
    Location: "los angeles",
    Limit:    20,
  })
  if err != nil {
    panic(err)
  }
  for i, biz := range results.Businesses {
    fmt.Printf("%d.) %s\n", i+1, biz.Name)
  }
}
