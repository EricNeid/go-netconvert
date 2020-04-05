all: build build-windows

build:
	cd cmd/netconvert && go build

build-windows:
	cd cmd/netconvert && GOOS=windows GOARCH=amd64 go build

test:
	go test ./...
