# at first run the "make bootstrap" command to download the lit cli.
# this cli we use to transform the "main.go.md" and build the binary.
# here the two commands to run:
#
#   make download-darwin # or download-linux
#   make bootstrap
#   make build
#

VERSION=0.3.1

build: pre
	@echo "build lit v${VERSION}"
	@go build -ldflags "-X main.version=${VERSION}" -o lit

pre:
	@echo "pre-processing"
	@./lit-bootstrap ../../transform.go.md ../../transform.go
	@./lit-bootstrap main.go.md main.go


download-darwin:
	@curl -L https://github.com/literate-programming/lit/releases/download/v0.3.0/lit_darwin_386 > lit-bootstrap
	@chmod +x lit-bootstrap
download-linux:
	@curl -L https://github.com/literate-programming/lit/releases/download/v0.3.0/lit_linux_386 > lit-bootstrap
	@chmod +x lit-bootstrap
bootstrap:
	@echo "bootstrap lit"
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
