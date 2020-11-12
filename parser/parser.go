package parser

import (
	"github.com/isurusiri/goj/constants"
)

// ParseArray iterates over the inputTokens and see if it
// starts with a left bracket and finishes with a right
// bracket with commas in between.
// Params:   inputTokens -  a map containing a dynamic set of
//                          values tokenized from the JSON.
// Returns: []interface{} - a map containing parsed
//                                array elements of the JSON
//                                array object.
//          []interface{} - unchecked elements of the
//                                inputTokens array.
func ParseArray(inputTokens []interface{}) ([]interface{}, []interface{}, error) {
	jsonArray := make([]interface{}, 0)

	token := inputTokens[0]
	if token == constants.JSONRightBracket {
		return jsonArray, inputTokens[1:], nil
	}

	return nil, nil, nil
}
