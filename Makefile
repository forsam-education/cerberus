build: main.go
	packr2 && GO111MODULE=on go build -ldflags="-s -w -X github.com/forsam-education/cerberus/utils.VersionHash=`git rev-parse HEAD` -X github.com/forsam-education/cerberus/utils.BuildTime=`date +%Y-%m-%dT%T%z` -X github.com/forsam-education/cerberus/utils.VersionIdentifier=`git describe --tags`" -o bin/cerberus .
