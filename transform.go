package literateprogramming

import (
	"bytes"
)

var (
	tokenCode    = []byte("    ")
	tokenTab     = []byte("\t")
	tokenNewLine = []byte("\n")
)

func Transform(code []byte, commentStyle []byte) []byte {
	lines := make([][]byte, 0)
	lines = bytes.Split(code, tokenNewLine)

	for i, v := range lines {
		if len(v) != 0 {
			if !bytes.HasPrefix(v, tokenCode) {
				if !bytes.HasPrefix(v, tokenTab) {
					// change each non code line
					lines[i] = append(commentStyle, v...)
				}
			} else {
				lines[i] = bytes.Replace(lines[i], tokenCode, []byte(""), 1)
			}
		}
	}
	return bytes.Join(lines, tokenNewLine)
}
