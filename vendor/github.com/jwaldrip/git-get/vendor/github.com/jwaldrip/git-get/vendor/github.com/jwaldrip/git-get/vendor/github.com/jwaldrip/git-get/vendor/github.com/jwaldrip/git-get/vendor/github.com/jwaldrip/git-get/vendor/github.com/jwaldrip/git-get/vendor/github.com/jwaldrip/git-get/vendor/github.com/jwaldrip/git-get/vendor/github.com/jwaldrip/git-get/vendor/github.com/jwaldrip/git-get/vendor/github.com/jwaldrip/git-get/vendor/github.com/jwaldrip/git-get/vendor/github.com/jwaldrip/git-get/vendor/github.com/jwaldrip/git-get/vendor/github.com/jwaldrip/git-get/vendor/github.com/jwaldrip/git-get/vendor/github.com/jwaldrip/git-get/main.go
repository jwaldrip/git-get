package main

import (
	"fmt"
	"os"

	"github.com/jwaldrip/odin/cli"
	git "github.com/libgit2/git2go"
)

var app = cli.New(version, "clone a repo into a common path", get)

func init() {
	app.DefineBoolFlag("force", false, "overwrite existing directory")
	app.AliasFlag('f', "force")
	app.DefineParams("url")
}

func main() {
	app.Start()
}

func get(c cli.Command) {
	if len(c.Args()) > 0 {
		exitWithMsg("too many arguments")
	}

	// Vars
	repoURL := parseURL(c.Param("url").String())
	clonePath := parsePath(repoURL)
	cloneOpts := cloneOptionsForURL(repoURL)

	// Remove if --force
	if c.Flag("force").Get() == true {
		os.RemoveAll(clonePath)
	}

	// Clone
	fmt.Printf("Cloning into '%s'...\n", clonePath)
	_, err := git.Clone(repoURL.String(), clonePath, cloneOpts)
	exitIfErr(err)
	fmt.Println("\nDone!")
}