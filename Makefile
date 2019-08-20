build:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-s -w -extldflags "-static"' -o bin/flash-darwin64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-s -w -extldflags "-static"' -o bin/flash-linux64
