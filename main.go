package main

import (
	"fmt"

	"github.com/jwaldrip/odin/cli"
	git "github.com/libgit2/git2go"
)

var app = cli.New(version, "clone a repo into a common path", func(c cli.Command) {
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
