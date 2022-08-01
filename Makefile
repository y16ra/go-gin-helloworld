all: clean test build package
build:
	GOARCH=amd64 GOOS=linux go build -o main
test:
	go test
clean:
	rm -f main
package:
	zip main.zip main