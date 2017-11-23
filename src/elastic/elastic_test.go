package elastic

import "test"
import (
  "testing"
  "net/http"
)



func TestAllSearchGET(t *testing.T) {
  var response ,_ = http.Get("http://localhost:9200/users/user/_search")
  defer response.Body.Close()

  test.AssertEquals(200 , response.StatusCode ,t)
}

func TestAllSearchGETNumberDocuments(t *testing.T) {

  test.AssertEquals(3 , New("localhost" , 9200).GetAllDocument("users" , "user").Hits.Total ,t)
}




