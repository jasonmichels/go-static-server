### Go Static Server
Work In Progress
Static server to serve single page javascript applications.  The use is for microservice front end apps.

### Optional Environment Variables
- APP_ENV=local
- NEW_RELIC_NAME=changeMe
- NEW_RELIC_KEY=changeMe
- PATH_PREFIX="/front-end-service-one"

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