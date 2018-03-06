package gitget

import (
	"fmt"
	"os"

	"github.com/jwaldrip/odin/cli"
	git "gopkg.in/src-d/go-git.v4"
)

var app = cli.New(version, "clone a repo into a common path", get)

func init() {
	app.DefineStringFlag("host", "github.com", "define the host for path based queries.")
	app.DefineBoolFlag("force", false, "overwrite existing directory.")
	app.DefineBoolFlag("ssh", false, "use ssh for the connection.")
	app.DefineBoolFlag("bitbucket", false, "lookup with bitbucket.")
	app.DefineBoolFlag("github", false, "lookup with github.")
	app.AliasFlag('f', "force")
	app.AliasFlag('H', "host")
	app.AliasFlag('S', "ssh")
	app.DefineParams("url")
}

func Start() {
	app.Start()
}

func get(c cli.Command) {
	if len(c.Args()) > 0 {
		exitWithMsg("too many arguments")
	}

	// Vars
	var host string
	if (c.Flag("bitbucket").Get() == true) {
		host = "bitbucket.org"
	} else if (c.Flag("github").Get() == true){
		host = "github.com"
	} else {
		host = c.Flag("host").String()
	}
	repoURL := parseURL(c.Param("url").String(), host, c.Flag("ssh").Get() == true)
	clonePath := parsePath(repoURL)
	cloneOpts := cloneOptionsForURL(repoURL.String())

	// Remove if --force
	if c.Flag("force").Get() == true {
		os.RemoveAll(clonePath)
	}

	// Clone
	fmt.Printf("Cloning '%s' into '%s'...\n", repoURL, clonePath)
	_, err := git.PlainClone(clonePath, false, cloneOpts)
	exitIfErr(err)
	fmt.Println("\nDone!")
}
