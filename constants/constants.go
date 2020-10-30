package constants

const (
	// JSONComma in a JSON object
	JSONComma = ","
	// JSONColon in a JSON object
	JSONColon = ";"
	// JSONLeftBracket in a JSON object
	JSONLeftBracket = "["
	// JSONRightBracket in a JSON object
	JSONRightBracket = "]"
	// JSONLeftBrace in a JSON object
	JSONLeftBrace = "{"
	// JSONRightBrace in a JSON object
	JSONRightBrace = "}"
	// JSONQuote in a JSON object
	JSONQuote = '"'
	// FalseLen is the character length of word 'false'
	FalseLen = len("false")
	// TrueLen is the character length of word 'true'
	TrueLen = len("true")
	// NullLen is the character length of word 'null'
	NullLen = len("null")
)

// JSONWhitespace possibilities
var JSONWhitespace = [...]string{" ", "\t", "\b", "\n", "\r"}

// JSONSyntax possibilities
var JSONSyntax = [...]string{JSONComma, JSONColon, JSONLeftBracket, JSONRightBracket, JSONLeftBrace, JSONRightBrace}
