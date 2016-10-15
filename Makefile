PROJECT = github.com/thngkaiyuan/look-at-my-site

.PHONY: build
build:
	go build ${PROJECT}

.PHONY: format
format:
	go fmt ${PROJECT}

.PHONY: lint
lint:
	go vet ${PROJECT}

.PHONY: pre-commit
pre-commit: lint build test

CURRENT_DIR = "$(shell pwd)"
EXPECTED_DIR = "${GOPATH}/src/github.com/thngkaiyuan/look-at-my-site"

.PHONY: check
check:
ifeq (${CURRENT_DIR}, ${EXPECTED_DIR})
	@echo "PASS: Current directory is in \$$GOPATH."
else
	@echo "FAIL"
	@echo "Expected: ${EXPECTED_DIR}"
	@echo "Actual: ${CURRENT_DIR}"
endif

.PHONY: serve
serve: check
	go run *.go

.PHONY: test
test:
	go test -v -cover "${PROJECT}"
