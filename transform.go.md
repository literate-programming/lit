# Transform

	package lit

	import (
		"bytes"
	)

Define some variables we use at the `Transform` function

	var (
		tokenCode    = []byte("    ")
		tokenTab     = []byte("\t")
		tokenNewLine = []byte("\n")
	)

The core of the lit tool is the `Transform` function. At this function we process the given source and comment out all non code lines by the given `commentStyle`.

	func Transform(code []byte, commentStyle []byte) []byte {
		lines := make([][]byte, 0)
		lines = bytes.Split(code, tokenNewLine)

		for i, v := range lines {
			if len(v) != 0 {
				if !bytes.HasPrefix(v, tokenCode) {
					if !bytes.HasPrefix(v, tokenTab) {
						// change each non code line
						lines[i] = append(commentStyle, v...)
					} else {
						lines[i] = bytes.Replace(lines[i], tokenTab, []byte(""), 1)
					}
				} else {
					lines[i] = bytes.Replace(lines[i], tokenCode, []byte(""), 1)
				}
			}
		}
		return bytes.Join(lines, tokenNewLine)
	}
