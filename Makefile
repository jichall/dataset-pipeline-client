CC=go
FLAGS=
SRC=file.go json_utils.go sender.go console.go  main.go
EXEC=client.out

all:
	cd src && go run $(SRC) -dir=./files/

build:
	cd src && go build -o $(EXEC)

test:
	cd src && go test -v *.go
