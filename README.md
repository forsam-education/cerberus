<p align="center">
    <img width="360" height="360" src="https://content.forsam.io/cerberus/logos/logo_360.png" alt="Cerberus" title="Cerberus" />
</p>
<br/><br/>
Cerberus is a simple yet powerful, modern and easily configurable reverse proxy solution.
<br/><br/>

![Lint & Build](https://github.com/forsam-education/cerberus/workflows/Lint%20&%20Build/badge.svg)
[![GoDoc](https://godoc.org/github.com/forsam-education/cerberus?status.svg)](https://godoc.org/github.com/forsam-education/cerberus)
[![Go Report Card](https://goreportcard.com/badge/github.com/forsam-education/cerberus)](https://goreportcard.com/report/github.com/forsam-education/cerberus)
![Version](https://img.shields.io/github/tag/forsam-education/cerberus?color=blue&label=alpha)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fforsam-education%2Fcerberus.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fforsam-education%2Fcerberus?ref=badge_shield)

## Disclaimer

This project is in a very early work in progress state, do not use in production.

## Contribute

If you want to contribute to this project , you can read the [code of conduct](./CODE_OF_CONDUCT.md) and [conventions](./CONVENTION.md), also check the current issues, and open pull requests.

You can also simply ask for features by creating issues.

## Dependencies

Cerberus depends on some services to store both it's temporary and persistent state.

The persistent state is stored in a MySQL 8 database, therefor you must have one available and set it properly in the configuration file.

For the temporary state storage (and sharing it between the different instances of Cerberus running in your stack), it uses a Redis server in it's fifth version.

## Quality assurance

To fix the basics of code format, you can run `go fmt ./...`.

For a bit more advanced code style checks, you can run `golint $(go list ./... | grep -v /vendor/)`. You'll have to run `go get -u golang.org/x/lint/golint` before.

## Dependencies upgrades

The dependendencies are automatically upgraded every Monday by Dependabot for both Javascript FrontEnd and Go Backend.

## Docker

We made a Docker image for the software, you can find it on [DockerHub](https://hub.docker.com/r/forsameducation/cerberus).
It is a two-stages build process, and the final build is based on `scratch` so it is as small as possible.

You can also find an example Docker Compose configuration file at `deployments/docker-compose.yml` and run it `docker-compose -f ./deployments/docker-compose.yml up`.

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fforsam-education%2Fcerberus.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fforsam-education%2Fcerberus?ref=badge_large)
