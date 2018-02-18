.PHONY: test
test:
	go test -v -cover ./libtsdb/...
.PHONY: fmt
fmt:
	gofmt -d -l -w ./libtsdb