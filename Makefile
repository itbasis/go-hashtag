go-dependencies:
	# https://asdf-vm.com/
	asdf install golang
	#
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint@latest && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go get github.com/onsi/ginkgo/v2@latest && go install github.com/onsi/ginkgo/v2/ginkgo
	# https://github.com/securego/gosec
	go get -u github.com/securego/gosec/v2/cmd/gosec@latest && go install github.com/securego/gosec/v2/cmd/gosec@latest
	#
	go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest && go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	#
	go get -u -t -v ./... || :

go-generate:
	go generate ./...

go-test:
	go vet -vettool=$(which shadow) ./...
	ginkgo ./...
	golangci-lint run
	gosec ./...

go-all: go-dependencies go-generate go-test
	go mod tidy || :
