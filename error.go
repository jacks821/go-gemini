package gemini

import (
	"fmt"
)

type Error struct {
	Reason	string
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("[%v] %v", e.Reason, e.Message)
}
