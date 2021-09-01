package main

import (
	"fmt"
	"os"
	"strings"
)

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

//Stringer is implemented by any value that has a String method,
//which defines the ``native'' format for that value.
//The String method is used to print values passed as an operand
//to any format that accepts a string or to an unformatted printer
//such as Print.
//type fmt.Stringer interface {
//	String() string
//}

func main() {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}
