package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type Repo struct {
	CloneUrl string `json:"clone_url"`
}

type Head struct {
	Repo Repo   `json:"repo"`
  Ref  string `json:"ref"`
}

type PullRequest struct {
	Head Head `json:"head"`
}

type Payload struct {
	PullRequest PullRequest `json:"pull_request"`
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

func Payloadfunc(w http.ResponseWriter, r *http.Request) {
  var err error
  var payload Payload
  var payloadBytes []byte

  payloadBytes, err = ioutil.ReadAll(r.Body)
  if err != nil {
    fmt.Println(err)
  }

  err = json.Unmarshal(payloadBytes, &payload)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(payload.PullRequest.Head.Repo.CloneUrl)
  fmt.Println(payload.PullRequest.Head.Ref)
}
