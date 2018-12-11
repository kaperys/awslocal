fmt:
	go fmt
	go vet

test: fmt
	@mkdir -p artifacts
	go test -race -short -cover -coverprofile=artifacts/coverage.out

coverage: test
	go tool cover -html=artifacts/coverage.out
