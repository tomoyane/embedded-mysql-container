package exception

import (
	"fmt"
)

type ErrorHandler struct{}

func (e ErrorHandler) ErrorMessage(msg string, err error) {
	fmt.Println("Internal Error: " + msg)
	panic(err.Error())
}
