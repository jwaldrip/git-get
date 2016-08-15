package main

import (
	"fmt"
	"os"

	"github.com/jwaldrip/tint"
)

func exitWithMsg(msgs ...interface{}) {
	fmt.Println(msgs...)
	os.Exit(1)
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, tint.Colorize(fmt.Sprintf("ERROR: %s", err), tint.Red))
	}
}

func exitIfErrWithMsg(err error, msgs ...interface{}) {
	if err != nil {
		exitWithMsg(msgs...)
	}
}
