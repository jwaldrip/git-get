package main

import (
	"net/url"

	git "github.com/libgit2/git2go"
)

func cloneOptionsForURL(u *url.URL) *git.CloneOptions {
	remoteCallbacks := &git.RemoteCallbacks{
		SidebandProgressCallback: sidebandProgressCallback,
		TransferProgressCallback: transferProgressCallback,
		CredentialsCallback:      buildCredentialsCallback(u),
		CertificateCheckCallback: buildCertificateCheckCallback(u),
	}
	return &git.CloneOptions{
		Bare:            false,
		RemoteCallbacks: remoteCallbacks,
	}
}
