//
// Fast snake-case implementation.
//
package snakecase

import "strings"

// Snakecase the given `str`.
func Snakecase(str string) string {
	var b [1024]byte
	max := 1024
	l := len(str)
	ret := ""
	bi := 0
	i := 0

	for i < l {
		for i < l && !isWord(str[i]) {
			i++
		}

		for i < l && isU(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && isPart(str[i]) {
			if bi < max {
				b[bi] = str[i]
				bi++
			}
			i++
		}

		for i < l && !isWord(str[i]) {
			i++
		}

		ret += strings.ToLower(string(b[:bi])) + "_"
		bi = 0
	}

	if len(ret) > 0 {
		ret = ret[:len(ret)-1]
	}

	return ret
}

func isPart(c byte) bool {
	return isLower(c) || isDigit(c)
}

func isWord(c byte) bool {
	return isLetter(c) || isDigit(c)
}

func isLetter(c byte) bool {
	return isLower(c) || isU(c)
}

func isU(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
