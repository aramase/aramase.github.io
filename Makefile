# CI pins the Hugo version in .github/workflows/hugo.yml (HUGO_VERSION); hugo.toml sets the minimum.
HUGO ?= hugo

build:
	$(HUGO) --gc --minify --cleanDestinationDir

build-draft:
	$(HUGO) --buildDrafts

preview:
	$(HUGO) server

clean:
	rm -rf public

# Regenerate data/ledger.yaml from kubernetes/enhancements on GitHub, or from a local
# clone with ENHANCEMENTS=<path>. See ledger/README.md.
ledger:
	cd tools/ledger && go run . -root ../.. $(if $(ENHANCEMENTS),-enhancements $(abspath $(ENHANCEMENTS)))

ledger-test:
	cd tools/ledger && go vet ./... && go test ./...

.PHONY: build build-draft preview clean ledger ledger-test
