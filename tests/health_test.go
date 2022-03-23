

package test

import (
	"fmt"
	"testing"
  "github.com/go-resty/resty/v2"
  "github.com/stretchr/testify/assert"
)

func TestHelathEndpoint(t *testing.T){
  fmt.Println("running e2e test for health check endpoint")

  client := resty.New()
  resp, err := client.R().Get("http://localhost:4000/api/health")

  if err != nil{
    t.Fail()
  }

  assert.Equal(t, 200 , resp.StatusCode())
}
