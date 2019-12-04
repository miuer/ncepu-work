package lexical

import (
	"regexp"
	"strings"
)

var (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*/().,:;=<> \n	#"
)

func isValid(char byte) bool {
	return strings.Contains(charset, string(char))
}

func isDigit(char byte) bool {
	matched, _ := regexp.MatchString(`[0-9]`, string(char))
	return matched
}

func isLetter(char byte) bool {
	matched, _ := regexp.MatchString(`[a-z]|[A-Z]`, string(char))
	return matched
}

func isSymbol(char byte) bool {
	matched, _ := regexp.MatchString(`:|>|<`, string(char))
	return matched
}

// --- useless
func isWhitespace(char byte) bool {
	matched, _ := regexp.MatchString(`\s`, string(char))
	return matched
}

func isLineBreaks(char byte) bool {
	matched, _ := regexp.MatchString(`\n`, string(char))
	return matched
}

func getMachineCode(name []byte) (machineCode int) {
	machineCode, ok := MachineMap[string(name)]
	if ok {
		return machineCode
	} else if strings.Contains(string(name), `.`) {
		return 20
	} else if isDigit(name[0]) {
		return 19
	} else {
		return 18
	}
}
