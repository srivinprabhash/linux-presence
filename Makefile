default: run

run:
	go run main.go

build:
	rm -rf main
	go build -o main -ldflags "-s -w" main.go