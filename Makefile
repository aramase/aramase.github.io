build:
	hugo --gc --minify --cleanDestinationDir

build-draft:
	hugo --buildDrafts

deploy: build
	netlify deploy --prod

deploy-preview:
	netlify deploy --draft

preview:
	hugo server

clean:
	rm -r public

build-deps:
	go get -u github.com/netlify/netlifyctl
	go get -u github.com/gohugoio/hugo

.PHONY: build deploy clean
