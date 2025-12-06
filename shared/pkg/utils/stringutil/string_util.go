package stringutil

func ToSnake(s string) string {
	if s == "" {
		return s
	}

	// Preallocate: worst-case every char becomes "_x"
	b := make([]byte, 0, len(s)+5)

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= 'A' && c <= 'Z' {
			// insert underscore if not first character
			if i > 0 {
				b = append(b, '_')
			}
			// convert uppercase to lowercase using ASCII trick
			c = c + ('a' - 'A')
		}

		b = append(b, c)
	}

	return string(b)
}
