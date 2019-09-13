# Transform

	package lit

	import (
		"bytes"
	)

The core of the lit tool is the `Transform` function. At this function we process the given source and comment out all non code lines by the given `commentStyle`.

	func Transform(code []byte, commentStyle []byte) []byte {
		tokens := Scanner(code)
		transformed := make([][]byte, len(tokens))
		for i := 0; i < len(tokens); i++ {
			switch tokens[i].Type {
			case TYPE_DOC, TYPE_CODESIGN:
				if len(tokens[i].Value) > 0 {
					transformed[i] = append(commentStyle, tokens[i].Value...)
				}
				break

			case TYPE_CODE:
				transformed[i] = tokens[i].Value
				break
			}
		}
		return bytes.Join(transformed, NEWLINE)
	}
