gen-coverage:
	go test -coverprofile coverprofile.out ./...
	go tool cover -func=coverprofile.out

gen-html-coverage:
	go test -coverprofile coverprofile.out ./...
	go tool cover -html=coverprofile.out

run-tests:
	go test ./...

build:
	go build
