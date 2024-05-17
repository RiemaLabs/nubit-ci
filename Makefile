.PHONY: build
build: nubitci-lint

nubitci-lint:
	go build -o $@ ./cmd/nubitci-lint

.PHONY: ci
ci:
	go run ./cmd/nubitci-lint

.PHONY: ci-fix
ci-fix:
	go run ./cmd/nubitci-lint -w

.PHONY: clean
clean:
	rm -rf nubitci-lint
