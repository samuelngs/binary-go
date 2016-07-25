
all:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/binary cmd/main.go
