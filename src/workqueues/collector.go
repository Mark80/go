package workqueues

import (
  "net/http"
  "time"
  "fmt"
)


var WorkQueue = make(chan WorkRequest , 100)

type HandlerWork struct{}

func (handler  *HandlerWork) ServeHTTP(responseWriter  http.ResponseWriter , request *http.Request ) {

  if request.Method != "POST" {
	responseWriter.Header().Set("Allow" , "POST")
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
   	return
  }

   delay ,err := time.ParseDuration(request.FormValue("Delay"))

  if err != nil {
	fmt.Println(err.Error())
	responseWriter.WriteHeader(http.StatusBadRequest)
	return
  }

  var name = request.FormValue("Name")
  var workRequest = WorkRequest{Name:name, Delay:delay}
  WorkQueue <- workRequest

  responseWriter.WriteHeader(http.StatusCreated)

}

