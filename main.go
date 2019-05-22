package main

import (
  "fmt"
  "github.com/google/go-github/v25/github"
  "context"
)

var (

)

const (

)

func main() {
  client := github.NewClient(nil)
  opt := &github.RepositoryListByOrgOptions{Type: "public"}
  repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
  if err != nil {
    panic(err)
  }
  for _, repo := range repos {
    fmt.Println(repo)
  }
}
