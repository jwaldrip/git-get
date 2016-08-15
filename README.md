# Git Get
Git clone that clones into a common path. It was inspired by `go get` and follows a similar folder structure.

## Installation

### Install with Homebrew (Mac)

```
$ brew install jwaldrip/utils/git-get
```

### Manual Install

#### Prerequisites

* Go
* GCC
* OpenSSL
* LibSSL2
* cmake
* pkg-config

```
$ git clone git@github.com:jwaldrip/git-get.git
$ cd git-get
$ make install
```

## Setup
You must have either a set a `$GITPATH` or `$GOPATH` variable.

_**NOTE:** using a `$GOPATH` will also append a /src to the path._

**bash:**
```
$ echo "export GITPATH=$HOME/dev" > .bashprofile
```

**zsh:**
```
$ echo "export GITPATH=$HOME/dev" > .zshrc
```

## Usage

```
$ git get [repo-url]
```

### Auto host
By default, if no host is specified, git-get will lookup repos by the provided
path on github. You can use the following flags to override this:

* `--bitbucket`
* `--github`
* `--host {hostname}`
* `-h {hostname}`

### Force SSH
By default git-get will either use the scheme passed, or default to https:, to
use SSH use the `--ssh` or `-S` flag.

## Example

```
$ git get git@github.com:jwaldrip/git-get
Cloning into '/Users/jwaldrip/dev/github.com/jwaldrip/git-get'...
remote: Counting objects: 13, done.
remote: s: 100% (11/11), done.11)   /11)
Total 13 (delta 2), reused 13 (delta 2)
Recieving objects: 92% (12/13) objects...
Recieving deltas: 100% (0/0) objects...
Indexing objects: 100% (13/13) objects...
```

## Contributing

See [CONTRIBUTING](https://github.com/jwaldrip/git-get/blob/master/CONTRIBUTING.md) for details on how to contribute.
