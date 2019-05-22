package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v25/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"os"
	"sync"
)

var (
	githubAccessToken string

	prodLogger *zap.Logger
	logger     zap.SugaredLogger
	once       sync.Once
)

const ()

func init() {
	var set bool
	once.Do(func() {
		prodLogger, _ = zap.NewDevelopment()
		logger = *prodLogger.Sugar()
		defer logger.Sync()
		githubAccessToken, set = os.LookupEnv("GITHUB_ACCESS_TOKEN")
		if !set {
			logger.Fatalw("Token environment variable not set",
				"envVarToSet", "GITHUB_ACCESS_TOKEN",
			)
		}
	})
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repos, _, err := client.Repositories.List(context.Background(), "", nil)
	if err != nil {
		panic(err)
	}
	for _, repo := range repos {
		fmt.Println(repo)
	}
}
