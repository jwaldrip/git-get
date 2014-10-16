GIT2GO=github.com/libgit2/git2go

build: get-dependencies
	go build .

install: get-dependencies
	go install

get-dependencies:
	go get $(GIT2GO)
	sh -c "cd $(GOPATH)/src/$(GIT2GO) && git submodule update --init"
	make -C $(GOPATH)/src/$(GIT2GO) install
