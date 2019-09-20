test:
	go test ./...

test-ci:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: test test-ci
