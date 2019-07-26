buildk:
	GO111MODULE=on go build -ldflags="-s -w" -o bin/cerberus ./main.go

install: buildk
	sudo cp bin/cerberus /usr/local/bin