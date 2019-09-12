package lit

import (
	"bytes"
	"testing"
)

func TestTransform(t *testing.T) {
	code := []byte(`# test code

this is a test string

    package main

    import "fmt"

    func() {
      fmt.Println("Hello literate programming")
    }
`)
	result := Transform(code, []byte("// "))
	expected := []byte(`// # test code

// this is a test string

package main

import "fmt"

func() {
  fmt.Println("Hello literate programming")
}
`)
	if !bytes.Equal(result, expected) {
		t.Error("result not equal")
		t.Logf("RESULT: %s", result)
		t.Logf("EXPECTED: %s", expected)
	}
}
