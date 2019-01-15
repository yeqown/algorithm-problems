package utils

// ConvRuneNum2Int ...
func ConvRuneNum2Int(c rune) int {
	return int(c - '0')
}

// IsNumber `0// 123456789` check
func IsNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

// IsPlus `+ - * / ` check
func IsPlus(c rune) bool {
	return c == '+'
}

// IsSubtract ...
func IsSubtract(c rune) bool {
	return c == '-'
}

// IsMultiplicati ...
func IsMultiplicati(c rune) bool {
	return c == '*'
}

// IsDivide ...
func IsDivide(c rune) bool {
	return c == '/'
}

// IsOpChar ...
func IsOpChar(c rune) bool {
	return IsPlus(c) || IsSubtract(c) || IsMultiplicati(c) || IsDivide(c)
}

// IsParenthesis `(` `)` check
func IsParenthesis(c rune) bool {
	return IsLeftParenthesis(c) || IsRightParenthesis(c)
}

// IsLeftParenthesis ...
func IsLeftParenthesis(c rune) bool {
	return c == '('
}

// IsRightParenthesis ...
func IsRightParenthesis(c rune) bool {
	return c == ')'
}
