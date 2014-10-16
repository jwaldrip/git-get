package main

import (
	"fmt"
	"os"
)

func exitWithMsg(msg string) {
	println(msg)
	os.Exit(1)
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func exitIfErrWithMsg(err error, msg string) {
	if err != nil {
		exitWithMsg(msg)
	}
}
