package gemini

import (
	"fmt"
)

//Error is just a standard error message
type Error struct {
	Reason	string
	Message string `json:"message"`
}

//Error returns a standard error using the Error struct.
func (e Error) Error() string {
	return fmt.Sprintf("[%v] %v", e.Reason, e.Message)
}
