package lit

import (
	"bytes"
)

var (
	TAB             = []byte("\t")
	NEWLINE         = []byte("\n")
	SPACE           = []byte(" ")
	FOUR_SPACES     = []byte("    ")
	THREE_BACKTICKS = []byte("```")
)

const (
	TYPE_DOC      = "DOC"
	TYPE_CODE     = "CODE"
	TYPE_CODESIGN = "CODESIGN"
)

type Token struct {
	Type  string
	Decl  string
	Value []byte
}

func Scanner(source []byte) []Token {
	lines := bytes.Split(source, NEWLINE)
	tokens := make([]Token, len(lines))

	lastType := ""
	for i := 0; i < len(lines); i++ {

		if bytes.HasPrefix(lines[i], FOUR_SPACES) {
			tokens[i].Type = TYPE_CODE
			tokens[i].Value = bytes.Replace(lines[i], FOUR_SPACES, []byte(""), 1)
			continue
		}
		if bytes.HasPrefix(lines[i], TAB) {
			tokens[i].Type = TYPE_CODE
			tokens[i].Value = bytes.Replace(lines[i], TAB, []byte(""), 1)
			continue
		}

		tokens[i].Value = lines[i]
		if len(lines[i]) >= 3 {
			if bytes.Equal(lines[i][0:3], THREE_BACKTICKS) {
				tokens[i].Type = TYPE_CODESIGN
				if lastType == "" {
					lastType = TYPE_CODESIGN
				} else {
					// reset the lastType
					lastType = ""
				}
				// get the declaration.
				// the declaration is the second string aster the three backticks.
				codesignDecl := bytes.Split(lines[i][3:], SPACE)
				if len(codesignDecl) > 1 {
					tokens[i].Decl = string(bytes.Join(codesignDecl[1:], SPACE))
				}
				continue
			}
		}

		if lastType == TYPE_CODESIGN {
			tokens[i].Type = TYPE_CODE
		} else {
			tokens[i].Type = TYPE_DOC
		}
	}

	return tokens
}
