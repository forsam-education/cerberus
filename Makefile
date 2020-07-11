buildcerberus:
	GO111MODULE=on go build -ldflags="-s -w -X github.com/forsam-education/cerberus/utils.VersionHash=`git rev-parse HEAD` -X github.com/forsam-education/cerberus/utils.BuildTime=`date +%Y-%m-%dT%T%z`" -o bin/cerberus

install: buildcerberus
	sudo cp bin/cerberus /usr/local/bin
