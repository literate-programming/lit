package lit

import (
	"bytes"
	"testing"
)

func TestScanner(t *testing.T) {
	source := []byte(`# test source
    echo hello
` + "```foo bar" + `
echo world
` + "```")
	tokens := Scanner(source)

	if len(tokens) != 5 {
		t.Errorf("tokens length not equal (has: %v must: 5)", len(tokens))
	}

	if tokens[0].Type != "DOC" {
		t.Error("token Type not DOC")
	}
	if tokens[0].Decl != "" {
		t.Error("token Decl not empty")
	}
	if !bytes.Equal(tokens[0].Value, []byte("# test source")) {
		t.Errorf("token Value not equal (%s)", tokens[1].Value)
	}

	if tokens[1].Type != "CODE" {
		t.Error("token Type not CODE")
	}
	if tokens[1].Decl != "" {
		t.Error("token Decl not empty")
	}
	if !bytes.Equal(tokens[1].Value, []byte("echo hello")) {
		t.Errorf("token Value not equal (%s)", tokens[1].Value)
	}

	if tokens[2].Type != "CODESIGN" {
		t.Error("token Type not CODESIGN")
	}
	if tokens[2].Decl != "bar" {
		t.Error("token Decl not bar")
	}
	if !bytes.Equal(tokens[2].Value, []byte("```foo bar")) {
		t.Errorf("token Value not equal (%s)", tokens[2].Value)
	}

	if tokens[3].Type != "CODE" {
		t.Error("token Type not CODE")
	}
	if tokens[3].Decl != "" {
		t.Error("token Type not equal")
	}
	if !bytes.Equal(tokens[3].Value, []byte("echo world")) {
		t.Errorf("token Value not equal (%s)", tokens[3].Value)
	}

	if tokens[4].Type != "CODESIGN" {
		t.Error("token Type not CODESIGN")
	}
	if tokens[4].Decl != "" {
		t.Error("token Type not equal")
	}
	if !bytes.Equal(tokens[4].Value, []byte("```")) {
		t.Errorf("token Value not equal (%s)", tokens[4].Value)
	}
}
