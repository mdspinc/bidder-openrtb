.PHONY: default all test regenerate clean_ffjson full_regenerate

default: test

deps:
	go get -t ./...

test: deps
	go test ./...

regenerate:
	go generate `go list ./... | grep -v '/vendor/'`

clean_ffjson:
	find -name '*ffjson*' -not -path './vendor/*' -delete

full_regenerate: clean_ffjson regenerate
	go generate `go list ./... | grep -v '/vendor/'`
