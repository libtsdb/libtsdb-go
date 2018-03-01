.PHONY: test
test:
	go test -v -cover ./libtsdb/...
.PHONY: fmt
fmt:
	gofmt -d -l -w ./libtsdb
.PHONY: generate
generate:
	gommon generate -v

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