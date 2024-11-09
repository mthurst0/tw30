.PHONY: all
all: build

.PHONY: build
build:
	./scripts/go-dev-bld

.PHONY: rel
rel:
	./scripts/go-rel-bld

.PHONY: completion
completion: build
	t=$(mktemp)
	./tw30 completion zsh > $t
	echo "source $t"

.PHONY: setup
setup:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/vektra/mockery/v2@v2.43.0
	go install github.com/a-h/templ/cmd/templ@v0.2.778

.PHONY: install
install: build
	./scripts/go-dev-bld install

.PHONY: docker-pull
docker-pull:	## - docker pull latest images
	@printf "\033[32m\xE2\x9c\x93 docker pull latest images\n\033[0m"
	@docker pull golang:alpine

.PHONY: docker
docker: docker-pull	## - Build the smallest and secured golang docker image based on scratch
	@printf "\033[32m\xE2\x9c\x93 Build the smallest and secured golang docker image based on scratch\n\033[0m"
	$(eval BUILDER_IMAGE=$(shell docker inspect --format='{{index .RepoDigests 0}}' golang:alpine))
	@export DOCKER_CONTENT_TRUST=1
	@docker build -f Dockerfile --build-arg "BUILDER_IMAGE=$(BUILDER_IMAGE)" -t tw30 .

.PHONY: get-htmx
get-htmx:
	curl -o ./internal/server/static/htmx.min.js https://unpkg.com/htmx.org@2.0.2/dist/htmx.min.js
	# We download the unminified version for development purposes, but exclude it from repo.
	curl -o ./internal/server/static/htmx.js https://unpkg.com/htmx.org@2.0.2/dist/htmx.js
