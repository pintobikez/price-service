# price service
price service is a small app to deal with price in a multi-channel manner.
The database used to store the data is a mysql one.
The pub/sub used is rabbitmq.

## Requirements
App requires Golang 1.9 or later, Glide Package Manager and Docker (for building)

## Installation
- Install [Golang](https://golang.org/doc/install)
- Install [Glide](https://glide.sh)
- Install [Docker](htts://docker.com)
- Install [Mysql-server](https://hub.docker.com/_/mysql/)
- Install [Rabbitmq](https://hub.docker.com/_/rabbitmq/)


## Build
For building binaries please use make, look at the commands bellow:

```
// Build the binary in your environment
$ make build

// Build with another OS. Default Linux
$ make OS=darwin build

// Build with custom version.
$ make APP_VERSION=0.1.0 build

// Build with custom app name.
$ make APP_NAME=price-service build

// Passing all flags
$ make OS=darwin APP_NAME=price-service APP_VERSION=0.1.0 build

// Clean Up
$ make clean

// Configure. Install app dependencies.
$ make configure

// Check if docker exists.
$ make depend

// Create a docker image with application
$ make pack

// Pack with custom Docker namespace. Default gfgit
$ make DOCKER_NS=gfgit pack

// Pack with custom version.
$ make APP_VERSION=0.1.0 pack

// Pack with custom app name.
$ make APP_NAME=price-service pack

// Pack passing all flags
$ make APP_NAME=price-service APP_VERSION=0.1.0 DOCKER_NS=gfgit pack
```

## Development
```
// Running tests
$ make test

// Running tests with coverage. Output coverage file: coverage.html
$ make test-coverage

// Running tests with junit report. Output coverage file: report.xml
$ make test-report
```

## Run it
```
// Run and launch docker
$ make build; docker build -t price-service-docker .; docker-compose up;
```

## Usage:

* PUT price CALL
curl -v -X PUT http://localhost:8000/price -H 'content-type: application/json' -d '{"id":"ABCDEFGH","prices":[{"price":100.00,"specialPrice":90.00,"specialFrom":"2017-11-02T15:04:05Z","specialTo":"2017-11-22T15:04:05Z","channel":"loja1"}]}'

* GET price CALL
curl -v -X GET http://localhost:8000/price/ABCDE