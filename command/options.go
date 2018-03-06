package command

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

func cloneOptionsForURL(u string) *git.CloneOptions {
	return &git.CloneOptions{
		URL: u,
		Progress: os.Stdout,
	}
}
