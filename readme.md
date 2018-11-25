
# Yelp Fusion (v3) Golang Client

A simple Yelp client written in Go.

Import this:
```
"github.com/ryanbennettvoid/go-yelp-fusion/yelp"
```

Have a config JSON file like this:
```
{
    "api_key": "API_KEY_GOES_HERE"
}
```

Example:
```
  authOptions, _ := yelp.AuthOptionsFromJsonFile( "../config.json" )
  client := yelp.NewClient( authOptions )
  results, _ := client.Search( yelp.SearchOptions {
    Term: "bar",
    Location: "los angeles",
    Limit: 20,
  } )
  for i, biz := range results.Businesses {
    fmt.Printf( "%d.) %s\n", i+1, biz.Name )
  }
```