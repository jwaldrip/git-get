package main

import (
	"bufio"
	"fmt"
	"os"

	git "github.com/libgit2/git2go"

	"github.com/jwaldrip/odin/cli"
)

var app = cli.New("0.0.1", "clone a repo like go get", clone)

func init() {
	app.DefineParams("url")
}

func main() {
	app.Start()
}

func clone(c cli.Command) {
	if len(c.Args()) > 0 {
		exitWithMsg("too many arguments")
	}
	u := parseURL(c.Param("url").String())
	p := parsePath(u)
	remoteCallbacks := &git.RemoteCallbacks{
		SidebandProgressCallback: remoteSidebandCallback,
		TransferProgressCallback: transferProgressCb,
		CredentialsCallback:      getAuthCallback(u),
	}
	cloneOpts := &git.CloneOptions{
		Bare:            false,
		RemoteCallbacks: remoteCallbacks,
	}
	fmt.Printf("Cloning into '%s'...\n", p)
	_, err := git.Clone(u.String(), p, cloneOpts)
	exitIfErr(err)
}

func getInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", question)
	text, _ := reader.ReadString('\n')
	return text
}
