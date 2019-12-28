buildwin: build test

WRIT_VERSION ?= 1.0

export GOPATH=$(CURDIR)/
export GOBIN=$(CURDIR)/.temp/
export WRIT_VERSION

clean:
	@rm -rf ./.output/

init: clean
	go get ./src/writ/...

build: init
	go build -o ./.output/writ .

test:
	go test
	go test -bench=.

fmt:
	@go fmt .
	@go fmt ./src/writ

dist: build test

	export GOOS=linux; \
	export GOARCH=amd64; \
	go build -o ./.output/writ .

	export GOOS=darwin; \
	export GOARCH=amd64; \
	go build -o ./.output/writ_osx .

	export GOOS=windows; \
	export GOARCH=amd64; \
	go build -o ./.output/writ.exe .

package: dist fpmPackage

fpmPackage: versionTest fpmTest

	fpm \
		--log error \
		-s dir \
		-t deb \
		-v $(WRIT_VERSION) \
		-n writ \
		--after-install=package/install.sh \
		./.output/writ=/usr/local/bin/writ \
		./docs/writ.7=/usr/share/man/man7/writ.7 \
		./autocomplete/writ=/etc/bash_completion.d/writ \
		./package/init.d.sh=/etc/init.d/writ \
		./package/defaults.sh=/etc/default/writ

	@mv ./*.deb ./.output/

dockerPackage: containerized_package dockerTest
	docker build . -t writ:latest

fpmTest:
ifeq ($(shell which fpm), )
	@echo "FPM is not installed, no packages will be made."
	@echo "https://github.com/jordansissel/fpm"
	@exit 1
endif

versionTest:
ifeq ($(WRIT_VERSION), )

	@echo "No 'WRIT_VERSION' was specified."
	@echo "Export a 'WRIT_VERSION' environment variable to perform a package"
	@exit 1
endif

dockerTest:
ifeq ($(shell which docker), )
	@echo "Docker is not installed."
	@exit 1
endif

containerized_build: dockerTest

	docker run \
		--rm \
		-v "$(CURDIR)":"/srv/build":rw \
		-u "$(shell id -u $(whoami)):$(shell id -g $(whoami))" \
		-e WRIT_VERSION=$(WRIT_VERSION) \
		golang:1.10 \
		bash -c \
		"cd /srv/build; make build"

containerized_dist: dockerTest

	docker run \
		--rm \
		-v "$(CURDIR)":"/srv/build":rw \
		-u "$(shell id -u $(whoami)):$(shell id -g $(whoami))" \
		-e WRIT_VERSION=$(WRIT_VERSION) \
		golang:1.10 \
		bash -c \
		"cd /srv/build; make dist"

containerized_package: dockerTest containerized_dist

	docker run \
		-v "$(CURDIR)":"/srv/build" \
		-u "$(shell id -u $(whoami)):$(shell id -g $(whoami))" \
		-e WRIT_VERSION=$(WRIT_VERSION) \
		alanfranz/fpm-within-docker:debian-stretch \
		bash -c \
		"cd /srv/build; make fpmPackage"
