
.PHONY: build
build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/binary cmd/main.go

.PHONY: install
install:
	cp -rf bin/binary $$GOPATH/bin/binary

.PHONY: test
test:
	bin/binary -dir ./test -out ./output -pkg test -max 300

.PHONY: clean
clean:
	rm -rf output/*
