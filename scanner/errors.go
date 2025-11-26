package scanner

import "errors"

var ErrInput = errors.New("Error input. Please try again...")
var ErrNoComand = errors.New("No command found. Please try again, or type /help")
