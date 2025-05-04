package main // if its in same package means we need to mention here

// import "fmt"

func multipleReturnValues() (string, uint, bool) {
	return sharedVariableBTWFunc, 1, false
}

// func functionFromAnotherPackage() {
// 	fmt.Println("Dummy package")
// } // same package file and we don't need to import function at if its same package but in different file
