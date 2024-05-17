GOOS :=
GOARCH :=
ENV := GOOS=$(GOOS) GOARCH=$(GOARCH)

.PHONY: build
build: nubitci-lint nubitci-release

nubitci-lint:
	env $(ENV) go build -o $@ ./cmd/$@

nubitci-release:
	env $(ENV) go build -o $@ ./cmd/$@

.PHONY: ci
ci:
	go run ./cmd/nubitci-lint

.PHONY: ci-fix
ci-fix:
	go run ./cmd/nubitci-lint -w

.PHONY: clean
clean:
	rm -rf nubitci-lint
