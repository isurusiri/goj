package utils

import "github.com/isurusiri/goj/constants"

// IsValidJSONSyntax check if the inputString is valid
// character representing Json syntax.
// Params: inputString interface{} - input string to check if
//                              contains a valid json syntax
// Returns: bool   - True if the inputString is a char
//                   representing Json syntax.
func IsValidJSONSyntax(inputString interface{}) bool {
	for _, syntax := range constants.JSONSyntax {
		if inputString == syntax {
			return true
		}
	}
	return false
}

// IsWhitesapce check if the inputString represents a
// whitespace.
// Params: inputString interface{} - input string to check if
//                              its a whitespace.
// Returns: bool   - True if the inputString is a whitespace.
func IsWhitesapce(inputString interface{}) bool {
	for _, whitespace := range constants.JSONWhitespace {
		if inputString == whitespace {
			return true
		}
	}
	return false
}
