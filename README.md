<p align="center">
  <h3 align="center">lit</h3>
  <p align="center">the universal literate programming tool</p>
  <p align="center">
    <a href="https://travis-ci.org/literate-programming/lit"><img alt="Travis" src="https://img.shields.io/travis/literate-programming/lit/master.svg?style=flat-square"> </a>
    <a href="https://github.com/literate-programming/lit/releases"><img alt="Software Release" src="https://img.shields.io/github/v/release/literate-programming/lit.svg?style=flat-square"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square"></a>
  </p>
</p>

---

## Installation

**Prebuild binaries**

Go to the [release page](https://github.com/literate-programming/lit/releases/tag/v0.3.0) and download the binary.

**Building from Source**

`lit` is written in [Golang](https://golang.org). You need to install go before you can continue with the next steps.

```sh
git clone git@github.com:literate-programming/lit.git
cd cmd/lit
make download-darwin # or download-linux
make bootstrap
make build
```

## Usage

To transform a file and write it back to a file, run:

```sh
lit source.go.md source.go
```

If you want to print it to stdout, run:

```sh
lit source.go.md
```
