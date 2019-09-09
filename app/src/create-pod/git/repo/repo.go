package repo

import (
  "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/plumbing"
)

func CloneRepo(repoName string) {
  _, err := git.PlainClone("./", false, &git.CloneOptions{
    URL: repoName,
    ReferenceName: plumbing.ReferenceName{
      ReferenceName: ""
    }
  })

  if err != nil {
    log.Fatal(err)
  }
}
