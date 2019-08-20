# go-flash 

Writes the content of an environment variable to a named file.

## Using it

The `flash` program expects exactly two arguments. The first is the name of an environment variable. The second is the name of an output file. The environment variable must be set. The output file must not already exist. Non-existant paths will be created automatically.

## Licensing

See the [LICENSE](LICENSE) file.

## Building it

It is pure Go. Run `go build`. Or maybe some variation like:

    GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-s -w -extldflags "-static"'
