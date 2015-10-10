# api2go-gorm-gin-crud-example
[![Build Status](https://travis-ci.org/hnakamur/api2go-gorm-gin-crud-example.svg?branch=master)](https://travis-ci.org/hnakamur/api2go-gorm-gin-crud-example)

This is a CRUD example using [jinzhu/gorm](https://github.com/jinzhu/gorm) and [gin-gonic/gin](https://github.com/gin-gonic/gin).
This example is a fork of the one in [manyminds/api2go](https://github.com/manyminds/api2go).

## Examples

Examples can be found [here](https://github.com/manyminds/api2go/blob/master/examples/crud_example.go).

## Database setup

Before running the server or running tests, copy .envrc.example to .envrc and edit .envrc for your need.
Two envrinment variables DB_DIALECT and DB_PARAMS are passed to sql.Open(driverName, datasourceName string) (https://golang.org/pkg/database/sql/#Open).

After editing, run the following command to set environment variables DB_DIALECT and DB_PARAMS.

```sh
source .envrc
```

## Tests

```sh
source .envrc
go test ./...
ginkgo -r                # Alternative
ginkgo watch -r -notify  # Watch for changes
```
