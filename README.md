### Go Static Server
Static server to serve single page javascript applications

### Optional Environment Variables
- APP_ENV=local
- NEW_RELIC_NAME=changeMe
- NEW_RELIC_KEY=changeMe
- PATH_PREFIX="/work-queue"

### Testing Examples
```sh
$ make test
$ go test -v ./... -bench . -cover
$ go test -v ./entities/ -bench . -covermode=count -coverprofile=coverage.out
$ m

### Language/Tech
  - Go 1.9
  - Docker

Authors
----
- Jason Michels