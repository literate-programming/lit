test:
	@./cmd/lit/lit -n scanner.go.md scanner.go
	@./cmd/lit/lit -n transform.go.md transform.go
	go test -v

.PHONY: test
