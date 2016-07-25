
.PHONY: build
build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/binary cmd/main.go

.PHONY: test
test:
	bin/binary -dir ./test -out ./output
