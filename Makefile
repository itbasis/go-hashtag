go-dependencies:
	# https://asdf-vm.com/
	asdf install golang
	#
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint@latest && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go get -u github.com/onsi/ginkgo/v2@latest && go install github.com/onsi/ginkgo/v2/ginkgo@latest
	# https://github.com/securego/gosec
	go get -u github.com/securego/gosec/v2/cmd/gosec@latest && go install github.com/securego/gosec/v2/cmd/gosec@latest
	#
	go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest && go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	#
	go get -u -t -v ./... || :
	asdf reshim golang

go-generate:
	go generate ./...

go-test:
	golangci-lint run ./...
	go vet -vettool=$(which shadow) ./...
	gosec ./...
#	go test ./...
	ginkgo -r -race --cover --coverprofile=.coverage-details.out ./...
	go tool cover -func=.coverage-details.out -o=.coverage.out
	cat .coverage.out

go-all: go-dependencies go-generate go-test
	go mod tidy || :
