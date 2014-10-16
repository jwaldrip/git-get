package main

import (
	"fmt"

	git "github.com/libgit2/git2go"

	"github.com/jwaldrip/odin/cli"
)

var app = cli.New("0.0.1", "clone a repo like go get", func(c cli.Command) {
	if len(c.Args()) > 0 {
		exitWithMsg("too many arguments")
	}
	u := parseURL(c.Param("url").String())
	p := parsePath(u)
	fmt.Printf("Cloning into '%s'...\n", p)
	_, err := git.Clone(u.String(), p, cloneOptionsForURL(u))
	exitIfErr(err)
})

func init() {
	app.DefineParams("url")
}

func main() {
	app.Start()
}
