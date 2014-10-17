GIT2GO=github.com/libgit2/git2go

build: get-dependencies
	go build -o ./bin/git-get

build-all: get-dependencies
	sh build-all.sh

install: get-dependencies
	go install

get-dependencies:
	go get $(GIT2GO) > /dev/null || true
	sh -c "cd $(GOPATH)/src/$(GIT2GO) && git submodule update --init"
	make -C $(GOPATH)/src/$(GIT2GO) install
