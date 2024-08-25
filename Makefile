.PHONY: test

all: test

test/verbose:
	@go test -v -failfast -cover ./internal/...

test:
	@go test -failfast -cover ./internal/...
