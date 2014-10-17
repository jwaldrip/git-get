# Git Get
Git clone that clones into a common path. It was inspired by `go get` and follows a similar folder structure.

## Installation
```
$ brew install git-get
```

### From Source

```
$ git checkout git@github.com:jwaldrip/git-get.git
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
