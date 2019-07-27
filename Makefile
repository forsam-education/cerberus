buildcerberus:
	GO111MODULE=on go build -ldflags="-s -w" -o bin/cerberus ./main.go

buildfront:
	(cd ./web && yarn install && yarn build)

install: buildcerberus
	sudo cp bin/cerberus /usr/local/bin
