compile:
	go build -o boil main.go

run:
	go run main.go

build:
	echo "building binary for linux amd64"
	GOOS=linux GOARCH=amd64 go build -o ./bin/boil-linux-amd64 main.go
	echo "building binary for linux arm"
	GOOS=linux GOARCH=arm go build -o ./bin/boil-linux-arm main.go

clean:
	rm -rf bin

install:
	mkdir -p ~/.config/boil/cauldron
	cp ./cauldron/* ~/.config/boil/cauldron
	sudo go build -o /usr/bin/boil main.go
