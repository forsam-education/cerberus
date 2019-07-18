build:
	GO111MODULE=on go build -ldflags="-s -w" -o bin/kerberos ./main.go

install: build
	sudo cp bin/kerberos /usr/local/bin