package main

import (
	"net/url"
	"os"
	"path"
	"strings"
)

func parseURL(str string) *url.URL {
	u, err := url.Parse(str)
	exitIfErrWithMsg(err, "invalid url")
	if u.Scheme == "" {
		u = parseURL("ssh://" + str)
		parts := strings.Split(u.Host, ":")
		u.Host = parts[0]
		u.Path = path.Join(parts[1], u.Path)
	}
	return u
}

func parsePath(u *url.URL) string {
	p := u.Path
	if p[len(p)-4:] == ".git" {
		p = string(p[:len(p)-4])
	}
	return path.Join(os.Getenv("GOPATH"), "src", u.Host, p)
}
