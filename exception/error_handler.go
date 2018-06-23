package exception

import (
	"fmt"
)

type ErrorHandler struct{}

func (e ErrorHandler) ErrorMessage(msg string) error {
	return fmt.Errorf("error %s", msg)
}
