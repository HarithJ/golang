package main

import (
  "os"

  "gopkg.in/src-d/go-git.v4"
)

func main() {
  git.PlainClone("/go", false, &git.CloneOptions{
    URL: "https://github.com/harithj/golang",
    Progress: os.Stdout,
  })
}
