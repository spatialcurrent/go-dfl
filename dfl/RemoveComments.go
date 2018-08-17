package dfl

// RemoveComments removes comments from a multi-line dfl expression
func RemoveComments(in string) string {
	out := ""
	singlequotes := 0
	doublequotes := 0
	comment := 0
	for _, c := range in {

		if comment == 0 {

			if singlequotes == 0 && doublequotes == 0 {
				if c == '"' {
					doublequotes += 1
				} else if c == '\'' {
					singlequotes += 1
				} else if c == '#' {
					comment += 1
				}
			} else if doublequotes == 1 && c == '"' {
				doublequotes -= 1
			} else if singlequotes == 1 && c == '\'' {
				singlequotes -= 1
			}

			if comment == 0 {
				out += string(c)
			}

		}

		if c == '\n' {
			comment = 0
		}

	}

	return out
}
