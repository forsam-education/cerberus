buildfront:
	(cd ./web && yarn install && yarn build && cd .. && packr2)

buildcerberus:
	GO111MODULE=on go build -ldflags="-s -w" -o bin/cerberus ./main.go

install: buildcerberus
	sudo cp bin/cerberus /usr/local/bin
