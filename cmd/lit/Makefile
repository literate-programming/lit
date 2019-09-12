VERSION=0.3.0

build: pre
	@echo "build lit v${VERSION}"
	@go build -ldflags "-X main.version=${VERSION}" -o lit

pre:
	@echo "pre-processing main.go.md"
	@./lit-bootstrap main.go.md main.go

bootstrap:
	@echo "bootstrap lit v${VERSION}..."
	@go build -ldflags "-X main.version=${VERSION}" -o lit-bootstrap
	@./lit-bootstrap version

test: build
	@./lit version
	@mkdir -p tmp
	@echo "build test sample"
	@./lit sample/sample.go.md tmp/main.go
	@echo "run test sample\n---"
	@go run tmp/main.go
	@echo "---"
	@rm -rf tmp

release:
	GOOS=darwin GOARCH=386 go build -ldflags "-X main.version=${VERSION}" -o lit_darwin_386
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o lit_darwin_amd64
	GOOS=linux GOARCH=386 go build -ldflags "-X main.version=${VERSION}" -o lit_linux_386
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o lit_linux_amd64
	GOOS=linux GOARCH=arm go build -ldflags "-X main.version=${VERSION}" -o lit_linux_arm
	GOOS=windows GOARCH=386 go build -ldflags "-X main.version=${VERSION}" -o lit_windows_386
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o lit_windows_amd64

clean:
	@rm -f lit
	@rm -f lit-bootstrap

.PHONY: build pre bootstrap test release clean
