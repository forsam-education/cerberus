buildk:
	GO111MODULE=on go build -ldflags="-s -w" -o bin/kerberos ./main.go

install: buildk
	sudo cp bin/kerberos /usr/local/bin