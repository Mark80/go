package remotecommand

import "net/http"



type AsyncServer interface{
  ServeAsyncHTTP(http.ResponseWriter, *http.Request, chan []byte)
}



