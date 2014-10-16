package main

import (
	"fmt"
	"net/url"

	git "github.com/libgit2/git2go"
)

var i = &info{}

func getAuthCallback(u *url.URL) git.CredentialsCallback {
	return func(url, usernameFromURL string, allowedTypes git.CredType) (int, *git.Cred) {
		i, cred := git.NewCredDefault()

		if allowedTypes&git.CredTypeUserpassPlaintext != 0 {
			i, cred = git.NewCredUserpassPlaintext(getInput("username"), getInput("password"))
			return i, &cred
		}
		if allowedTypes&git.CredTypeSshKey != 0 {
			i, cred = git.NewCredSshKeyFromAgent(u.User.Username())
			return i, &cred
		}
		if allowedTypes&git.CredTypeSshCustom != 0 {
			exitWithMsg("custom ssh not implemented")
		}

		if allowedTypes&git.CredTypeDefault == 0 {
			exitWithMsg("invalid cred type")
		}

		return i, &cred
	}
}

func remoteSidebandCallback(str string) int {
	fmt.Printf("\rremote: %v", str)
	return 0
}

func transferProgressCb(stats git.TransferProgress) int {
	i.stats = stats
	return i.update()
}

func calcPercent(partial, total uint) uint {
	percent := (float64(partial) / float64(total)) * 100
	return uint(percent)
}
