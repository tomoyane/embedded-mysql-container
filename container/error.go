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
	fmt.Println(error.Error(e.error))
}
