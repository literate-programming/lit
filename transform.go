package lit

import (
	"bytes"
)

func Transform(code []byte, commentStyle []byte, noDocs bool) []byte {
	tokens := Scanner(code)
	transformed := make([][]byte, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Type {
		case TYPE_DOC, TYPE_CODESIGN:
			if len(tokens[i].Value) == 0 {
				transformed = append(transformed, []byte(""))
			} else {
				if !noDocs {
					transformed = append(transformed, append(commentStyle, tokens[i].Value...))
				}
			}
			break

		case TYPE_CODE:
			transformed = append(transformed, tokens[i].Value)
			break
		}
	}
	return bytes.Join(transformed, NEWLINE)
}
