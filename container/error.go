package container

import (
	"fmt"
)

type ErrorContainer struct{
	msg string
	error error
}

func (e ErrorContainer) ErrorMessage() {
	fmt.Println("Internal Error: " + e.msg)
	panic(error.Error(e.error))
}
