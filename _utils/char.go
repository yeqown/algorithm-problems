/*
 主要提供给array.Calculator使用
*/
package _utils

func ConvRuneNum2Int(c rune) int {
	return int(c - '0')
}

// `0123456789` check
func IsNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

// `+ - * / ` check
func IsPlus(c rune) bool {
	return c == '+'
}

func IsSubtract(c rune) bool {
	return c == '-'
}

func IsMultiplicati(c rune) bool {
	return c == '*'
}

func IsDivide(c rune) bool {
	return c == '/'
}

func IsOpChar(c rune) bool {
	return IsPlus(c) || IsSubtract(c) || IsMultiplicati(c) || IsDivide(c)
}

// `(` `)` check
func IsParenthesis(c rune) bool {
	return IsLeftParenthesis(c) || IsRightParenthesis(c)
}

func IsLeftParenthesis(c rune) bool {
	return c == '('
}

func IsRightParenthesis(c rune) bool {
	return c == ')'
}
