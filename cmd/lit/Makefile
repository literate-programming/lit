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

clean:
	@rm -f lit
	@rm -f lit-bootstrap
