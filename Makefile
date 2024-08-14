go-all: go-update go-generate go-all-tests
go-all-tests: go-lint go-unit-tests

go-dependencies:
	$(eval GOBIN=$(shell go env GOPATH 2>/dev/null)/bin)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) latest
	#
	go install github.com/onsi/ginkgo/v2/ginkgo@latest

go-update: go-dependencies
	go mod tidy && go get -t -v -u ./...

go-generate:
	go generate ./...
	$(MAKE) go-update

go-lint:
	golangci-lint run

go-unit-tests:
	ginkgo -race --cover --coverprofile="ginkgo-coverage-unit.out" --junit-report="junit-report.xml" ./...
	go tool cover -func "ginkgo-coverage-unit.out" -o "coverage-unit.out"
	go tool cover -html "ginkgo-coverage-unit.out" -o "coverage-unit.html"

	cat coverage-unit.out
