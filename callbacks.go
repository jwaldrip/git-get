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
	return func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
		if u.Scheme != "ssh" && valid == false {
			exitWithMsg("host key check failed for:", hostname)
		}
		return 0
	}
}

func buildCredentialsCallback(u *url.URL) git.CredentialsCallback {
	return func(url, usernameFromURL string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
		i, cred := git.NewCredDefault()

		if allowedTypes&git.CredTypeUserpassPlaintext != 0 {
			i, cred = git.NewCredUserpassPlaintext(getInput("username"), getMaskedInput("password"))
			return git.ErrorCode(i), &cred
		}
		if allowedTypes&git.CredTypeSshKey != 0 {
			i, cred = git.NewCredSshKeyFromAgent(u.User.Username())
			return git.ErrorCode(i), &cred
		}
		if allowedTypes&git.CredTypeSshCustom != 0 {
			exitWithMsg("custom ssh not implemented")
		}

		if allowedTypes&git.CredTypeDefault == 0 {
			exitWithMsg("invalid cred type")
		}

		return git.ErrorCode(i), &cred
	}
}

func sidebandProgressCallback(str string) git.ErrorCode {
	fmt.Printf("\rremote: %v", str)
	return 0
}

func transferProgressCallback(stats git.TransferProgress) git.ErrorCode {
	i.stats = stats
	return git.ErrorCode(i.update())
}
