build: get-dependencies
	go build -o ./bin/git-get

install: get-dependencies
	go install

get-dependencies:
	dep ensure -v
