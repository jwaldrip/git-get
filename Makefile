GIT2GO=github.com/libgit2/git2go

build: get-dependencies
	go build -o ./bin/git-get

install: get-dependencies
	go install

get-dependencies:
	go get -u github.com/jwaldrip/odin
	go get -u github.com/jwaldrip/tint
	go get -u github.com/bgentry/speakeasy
	go get -u -d $(GIT2GO)
	cd $(GOPATH)/src/$(GIT2GO) && git submodule update --init && make install
