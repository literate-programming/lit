test:
	@./cmd/lit/lit -n transform.go.md transform.go
	go test -v
