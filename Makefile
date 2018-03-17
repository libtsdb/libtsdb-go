VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/utsdb

.PHONY: clean
	go clean -cache
	go clean -testcache

.PHONY: test
test:
	go test -v -cover ./libtsdb/...

.PHONY: fmt
fmt:
	gofmt -d -l -w ./libtsdb

.PHONY: generate
generate:
	gommon generate -v

.PHONY: loc
loc:
	cloc --exclude-dir=vendor,.idea,playground,vagrant,node_modules,libtsdbpb .

##--- docker ---#
.PHONY: docker-stop-all-containers
docker-stop-all-containers:
	docker stop $(shell docker ps -a -q)

.PHONY: docker-remove-all-containers
docker-remove-all-containers:
	docker rm $(shell docker ps -a -q)

.PHONY: docker-remove-all-volume
docker-remove-all-volume:
	docker volume prune

.PHONY: docker-clean-all
docker-clean-all:
	docker system prune
##--- docker ---#