package main

import (
  "os"
  "log"

  "gopkg.in/src-d/go-git.v4"
)

func main() {
  _, err := git.PlainClone("/go/git/temp", false, &git.CloneOptions{
    URL: "https://github.com/harithj/golang/docker",
    Progress: os.Stdout,
  })

  if err != nil {
    log.Fatal(err)
  }
}
