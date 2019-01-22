// Package snakecase - Super-Fast snake-case implementation.
package snakecase

const underscorByte = '_'

// Snakecase the given string.
func Snakecase(s string) string {
	idx := 0

	// loop through all good characters:
	// - lowercase
	// - digit
	// - underscore (as long as the next character is lowercase or digit)
	for idx < len(s) && ((isLower(s[idx]) || isDigit(s[idx])) || (s[idx] == underscorByte && idx < len(s)-1 && (isLower(s[idx+1]) || isDigit(s[idx+1])))) {
		idx++
	}

	// if we haven't gone through all of the characters then we must need to manipulate the string
	if idx < len(s) {
		b := make([]byte, 0, 64)
		// handles digit followed by an uppercase character
		if idx > 0 && isDigit(s[idx-1]) {
			idx--
		}
		b = append(b, s[:idx]...)

		for idx < len(s) {
			if !isAlphanumeric(s[idx]) {
				idx++
				continue
			}

			if len(b) > 0 {
				b = append(b, underscorByte)
			}

			for idx < len(s) && (isUpper(s[idx]) || isDigit(s[idx])) {
				b = append(b, toLower(s[idx]))
				idx++
			}

			for idx < len(s) && (isLower(s[idx]) || isDigit(s[idx])) {
				b = append(b, s[idx])
				idx++
			}
		}
		return string(b) // return manipulated string
	}
	return s // no changes needed, can just borrow the string
}

func isAlphanumeric(c byte) bool {
	return isLower(c) || isUpper(c) || isDigit(c)
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func toLower(c byte) byte {
	if isUpper(c) {
		return c + ('a' - 'A')
	}
	return c
}
