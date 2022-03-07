all: dev

build: cmd/main.go
	go build -o one cmd/main.go

dev: cmd/main.go
	go run cmd/main.go

clean:
	rm -f one

.PHONY: build clean
