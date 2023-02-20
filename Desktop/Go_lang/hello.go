package main

import "fmt"

// this is the package level variable so it si accessible through out the folder and can be accessible in different files as well
var temp = 53

// to convert the number to the string we use the below code
var r = string(rune(temp))

func main() {
	// msg := "golang"  ---- this is short way to declare the variable
	var msg = "hello world !"
	// var msg string = "work"  --- this is this type declaration of the variable

	fmt.Println(msg)
	fmt.Println(temp)
	fmt.Println("calling foo()")
	fmt.Println(r)
	foo()
	func1()
}

// we cannot change the const  decalred variable
// type declaration is used to increase the readability in the long run
// we can declare the type as
var apple int8 // this we do as to make our code more compatible to different machines
// as the int value depends on machine so to use the code in the different machine we declare the int size before hand in go

// scope
func func1() *int { // here we will have to declar ethe type of the return typep otherwise we will get the error
	var a int = 80 // short variable declaration
	return &a      // this is pointer concept like C    // the return type is pointer
} // here (*int) is the pointer [ return type]
// here we will get th eerro in case of keeping the unused variable

// form of identifiers -- we cannot start identifier with a number
// if we cd into .. then we go to previous directory
// if we cd in . then we go into the current directory

// IMPORTANT CMD FOR GIT
// git init   -- initializing the git repository
// git clone repository_name -- we get existing reporitory
// git status -- to see which files are merged etc we use this
// git add hello.go go.mod
// git commit -- it will ask for the our email and password to show our identity
// git config --global ==== for seperation we can use it
// git log  -- to get the changes wi e have made up till now in that file on the git

// TO check the changes made by the user in git
// git diff
