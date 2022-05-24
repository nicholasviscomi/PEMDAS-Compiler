package main

import (
	"fmt"
	"io/ioutil"
)

//use a stack and maybe a lexer to compute
//basic math operations
// 5 + 8 * 7 / (5 - 2)

func main() {
	fmt.Println("<PEMDAS Compiler>")
	content, err := ioutil.ReadFile("file.math")

	if err != nil { 
		fmt.Printf("Error reading from file\n")
		return
	}

   	fmt.Println(string(content))

	//input := "5 + 8 * 7 / ( 5 - 2 )"
	input := string(content)
	tokens := Tokenize(input)
	PrintTokens(tokens)
}