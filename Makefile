PACKAGES = $(shell go list ./...)
GO := CGO_ENABLED=0 go

.PHONY: all
all: build

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -rf ./bin/

.PHONY: format
format: go/fmt

.PHONY: go/fmt
go/fmt:
	$(GO) fmt $(PACKAGES)

.PHONY: test
test:
	@for PKG in $(PACKAGES); do $(GO) test -cover $$PKG || exit 1; done;

.PHONY: generate
generate: build
	./bin/githubhook-gen --output=githubevents

.PHONY: build
build: \
	bin/gen

.PHONY: bin/gen
bin/gen:
	mkdir -p bin
	$(GO) build -v -ldflags '-w $(LDFLAGS)' -o ./bin/githubhook-gen ./gen/*.go
