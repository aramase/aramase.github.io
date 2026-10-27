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

.PHONY: build build-draft preview clean
