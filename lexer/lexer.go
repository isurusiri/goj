package lexer

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/isurusiri/goj/constants"
	"github.com/isurusiri/goj/utils"
)

// LexString First it checks if the string starts with a quote,
// then iterate over the string until it finds the ending quote.
// If there's no ending quote it returns an error, or if it finds
// an initial quote and an ending quote, return the string within
// the quotes and the rest of the unchecked input string.
// Params: inputString string - input string to check if contains
//                              a json string
// Returns: string - single part of JSON string identified in
//                   inputString by starting and ending quotes.
//          string - unchecked part of the inputString.
//          error  - throws this if ending quote is missing in
//                   the inputString
func LexString(inputString string) (string, string, error) {
	jsonString := ""

	if inputString[0] == constants.JSONQuote {
		inputString = inputString[1:]
	} else {
		return "", inputString, nil
	}

	for pos, char := range inputString {
		if char == constants.JSONQuote {
			return jsonString, inputString[pos+1:], nil
		}

		jsonString += string(char)
	}

	return "", inputString, errors.New("lexer: Expected end-of-string quote")
}

// LexNumber Iterates over the inputString until it finds a
// character that cannot be a number. Then it returns an integer
// (limitation) of number characters detected so far and the
// rest of the string. Return nil and the inputString at it is
// if there aren't any numbers in the inputString.
// Params: inputString string - input string to check if contains
//                              a json string
// Returns: *int   - pointer to the number identified in the
//                   inputString.
//          string - unchecked part of the inputString.
func LexNumber(inputString string) (*int, string) {
	jsonNumber := ""

	numberCharacters := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true}

	for _, char := range inputString {
		character := string(char)
		if numberCharacters[character] {
			jsonNumber += character
		} else {
			break
		}
	}

	rest := inputString[len(jsonNumber):]

	if len(jsonNumber) <= 0 {
		return nil, inputString
	}

	parsedNumber, _ := strconv.Atoi(jsonNumber)
	return &parsedNumber, rest
}

// LexBoolean Check if the inputString starts with a boolean
// value 'true' or 'false' by checking the string length and
// by a simple string matching. If the inputString starts with
// a boolean value, returns the boolean value and rest of the
// inputString.
// Params: inputString string - input string to check if contains
//                              a json string
// Returns: *bool  - pointer to the boolean value identified
//                   in the inputString.
//          string - unchecked part of the inputString.
func LexBoolean(inputString string) (*bool, string) {
	stringLength := len(inputString)
	boolValue := false

	if (stringLength >= constants.TrueLen) && inputString[:constants.TrueLen] == "true" {
		boolValue = true
		return &boolValue, inputString[constants.TrueLen:]
	} else if (stringLength >= constants.FalseLen) && inputString[:constants.FalseLen] == "false" {
		return &boolValue, inputString[constants.FalseLen:]
	}

	return nil, inputString
}

// LexNull Check if the inputString starts with null as it's
// value by checking the string length and by a simple string
// matching. If the inputString starts with null, returns true
// with the rest of the inputString, otherwise returns false
// and the inputString without a change.
// Params: inputString string - input string to check if contains
//                              a json string
// Returns: bool   - True if the inputString starts with null.
//          string - unchecked part of the inputString.
func LexNull(inputString string) (bool, string) {
	stringLength := len(inputString)

	if (stringLength >= constants.NullLen) && inputString[:constants.NullLen] == "null" {
		return true, inputString[constants.NullLen:]
	}

	return false, inputString
}

// Lex performs the lexical analysis of the inputString that
// represents the JSON object.
// Params: inputString string - input string to perform lexical
//                              analysis.
// Returns: map[int]interface{} - a map representing the parsed
//                                json object that contains all
//                                elements and their position.
//          error               - not null if there is an error
//                                occurred during the process.
func Lex(inputString string) (map[int]interface{}, error) {
	tokens := make(map[int]interface{})

	for i := 0; i >= 0; i++ {

		jsonString, inputString, _ := LexString(inputString)
		if jsonString != "" {
			tokens[i] = jsonString
			continue
		}

		jsonNumber, inputString := LexNumber(inputString)
		if jsonNumber != nil {
			tokens[i] = &jsonNumber
			continue
		}

		jsonBool, inputString := LexBoolean(inputString)
		if jsonBool != nil {
			tokens[i] = &jsonBool
			continue
		}

		isJSONNull, inputString := LexNull(inputString)
		if isJSONNull {
			tokens[i] = nil // should this be "Null" ?
			continue
		}

		if len(inputString) <= 0 {
			break
		}

		char := inputString[0]

		if utils.IsWhitesapce(char) {
			// ignore whitespace
			inputString = inputString[1:]
		} else if utils.IsValidJSONSyntax(char) {
			tokens[i] = char
			inputString = inputString[1:]
		} else {
			return nil, fmt.Errorf("lexer: Expected character %v", char)
		}

	}

	return tokens, nil
}
