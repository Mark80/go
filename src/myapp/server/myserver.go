package server

import (
  "net/http"
  "strings"
)

type MyMx struct{}

func (m *MyMx)  ServeHTTP(w http.ResponseWriter, r *http.Request) {

  switch r.URL.Path {

  case "/login" :
	Login(w, r)

  default :
	w.WriteHeader(http.StatusNotFound)

  }


}

func Login(w http.ResponseWriter, req *http.Request) {

  method := req.Method

  req.ParseForm()

  if method == "POST" {

	var nome = req.Form["nome"][0]
	var password = req.Form["password"][0]

	if "MARCO" == strings.ToUpper(nome) && password == "pass" {
	  w.WriteHeader(http.StatusOK)
	} else {
	  w.WriteHeader(http.StatusUnauthorized)

	}
  } else {
	w.WriteHeader(http.StatusMethodNotAllowed)

  }

}

