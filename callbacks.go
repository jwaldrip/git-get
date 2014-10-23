package main

import (
	"fmt"
	"net/url"

	git "github.com/libgit2/git2go"
)

var i = &info{}

func calcPercent(partial, total uint) uint {
	percent := (float64(partial) / float64(total)) * 100
	return uint(percent)
}

func buildCertificateCheckCallback(u *url.URL) git.CertificateCheckCallback {
	return func(cert *git.Certificate, valid bool, hostname string) int {
		if u.Scheme != "ssh" && valid == false {
			exitWithMsg("host key check failed for:", hostname)
		}
		return 0
	}
}

func buildCredentialsCallback(u *url.URL) git.CredentialsCallback {
	return func(url, usernameFromURL string, allowedTypes git.CredType) (int, *git.Cred) {
		i, cred := git.NewCredDefault()

		fmt.Println("run auth callback:", allowedTypes)

		if allowedTypes&git.CredTypeUserpassPlaintext != 0 {
			i, cred = git.NewCredUserpassPlaintext(getInput("username"), getMaskedInput("password"))
			return i, &cred
		}
		if allowedTypes&git.CredTypeSshKey != 0 {
			fmt.Println("ssh key")
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

func sidebandProgressCallback(str string) int {
	fmt.Printf("\rremote: %v", str)
	return 0
}

func transferProgressCallback(stats git.TransferProgress) int {
	i.stats = stats
	return i.update()
}
