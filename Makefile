build-docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-static-server .
	docker build -t jasonmichels/go-static-server:develop .
test:
	go test -v ./... -bench . -cover