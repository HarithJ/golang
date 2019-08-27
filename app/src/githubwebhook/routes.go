package main

import (
  "net/http"

  "github.com/gorilla/mux"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  Handlerfunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  for _, route := range routes {
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(route.Handlerfunc)
  }

  return router
}

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "githubwebhook",
    "POST",
    "/payload",
    Payloadfunc,
  },
}
