package test

import (
  "testing"
  "net/url"
  "net/http"
  "net/http/httptest"
  "myapp/server"
  "test"
)

var mymux = &server.MyMx{}

func TestLoginOK(t *testing.T) {

  req := loginRequest("marco", "pass")
  response := send(req)

  test.AssertEquals(200 ,response.Code,t)

}

func TestLoginWithWrongCredential(t *testing.T) {

  req := loginRequest("pippo", "pass")
  response := send(req)

  test.AssertEquals(401 ,response.Code,t)

}

func send(request *http.Request) *httptest.ResponseRecorder {

  rw := httptest.NewRecorder()
  mymux.ServeHTTP(rw, request)

  return rw

}

func loginRequest(name string, password string) *http.Request {

  req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/login", nil)

  req.Form = url.Values{}

  req.Form.Set("nome", name)
  req.Form.Set("password", password)

  return req
}
