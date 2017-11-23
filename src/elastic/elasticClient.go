package elastic

import (
  "net/http"
  "encoding/json"
  "fmt"
)

type Hits struct {
  Total int `json:"total"`
}

type Result struct {
  Hits  Hits `json:"hits"`
}

type ElasticClient struct {
  Host string
  Port int
}

func New(host string , port int) ElasticClient {
  return ElasticClient{host ,port}
}


func (elastic ElasticClient)GetAllDocument(_index string , _type string) Result {
  var client = http.Client{}
  var query = fmt.Sprintf("http://%s:%d/%s/%s/_search" ,elastic.Host, elastic.Port ,_index,_type)
  var response ,_ =  client.Get(query)
  defer response.Body.Close()
  return decodeResult(response)

}

func decodeResult(response *http.Response) Result {
  decoder  := json.NewDecoder(response.Body)
  var result = &Result{}
  decoder.Decode(result)

  return *result

}
