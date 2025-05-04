package helper1

import "fmt"

// global level scope
var Tickes = 50 // exporting the variables by capitalizing the first letter

func FunctionFromAnotherPackage() { // by capitalizing the first letter of function, it will be export for whole package
	fmt.Println("FunctionFromAnotherPackage")
}
