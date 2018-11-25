
package yelp

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kr/pretty"
)

func TestYelp( t *testing.T ) {
  assert := assert.New( t )
  configPath := "../config.json"
  authOptions, err := AuthOptionsFromJsonFile( configPath )
  assert.NoError( err )
  client := NewClient( authOptions )
  assert.NotNil( client )
  results, err := client.Search( SearchOptions {
    Term: "bars",
    Location: "los angeles",
    Limit: 50,
  } )
  assert.NoError( err )
  assert.NotNil( results )
  assert.True( len( results.Businesses ) > 0 )
  assert.True( len( results.Businesses ) <= 50 )
  pretty.Println( results )
}