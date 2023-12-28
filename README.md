# Alpine.js 101

## prerequisites
see [golang docs](https://go.dev/doc/install) for instructions on how to install GO.

> GO isn't necessarily required. Any program or method that can start a *local* web server should work.

## how to run
```sh
go mod tidy
GIN_MODE=release go run . & open http://localhost:8000
```
> this should work on `linux` and `macos`
