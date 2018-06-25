package main

import (
	"fmt"
)

type Error struct{}

func (e Error) ErrorMessage(msg string, err error) {
	fmt.Println("Internal Error: " + msg)
	panic(err.Error())
}
