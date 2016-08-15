build: get-dependencies
	go build -o ./bin/git-get

install: get-dependencies
	go install

get-dependencies:
	cd vendor/github.com/libgit2/git2go && git submodule update --init && make install
