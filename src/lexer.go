package main

import (
	"fmt"
	"strconv"
	"strings"
)

//input: 123 + 456
//return: ["Number": 123, "Operation": add, "Number": 456]

const (
	//DATA TYPES
	INT = "INT"
	//MATHEMATICAL OPERATIONS
	OP_ADD = "ADD"
	OP_SUBT = "SUBTRACT"
	OP_MULT = "MULTIPLY"
	OP_DIV = "DIVIDE"
	//CONTROL FLOW
	L_PAREN = "OPEN_PAREN"
	R_PAREN = "CLOSE_PAREN"
)

type Token struct {
	Type string
	val interface {}
}

func tokenFromWord(currWord string) Token {

	if currWord == "+" {
		return Token {
			Type: OP_ADD,
		}
	} else if currWord == "-" {
		return Token {
			Type: OP_SUBT,
		}
	} else if currWord == "*" {
		return Token {
			Type: OP_MULT,
		}
	} else if currWord == "/" {
		return Token {
			Type: OP_DIV,
		}
	} else if currWord == "(" {
		return Token {
			Type: L_PAREN,
		}
	} else if currWord == ")" {
		return Token {
			Type: R_PAREN,
		}
	} else {
		n, _ := strconv.Atoi(currWord)
		return Token {
			Type: INT,
			val: n,
		}
	}
}

func Tokenize(input string) []Token {
	tokens := make([]Token, 0)
	input = strings.TrimSpace(input)
	/*
	if current character is not a space add it to the current word
	else classify the currWord as a token and append it to tokens
	*/
	
	currWord := ""
	for _, c := range input {
		if c != ' ' {
			currWord = currWord + string(c)
		} else {
			tokens = append(tokens, tokenFromWord(currWord))
			currWord = ""
		}
	}
	tokens = append(tokens, tokenFromWord(currWord)) //add final token

	return tokens
}

func PrintTokens(tokens []Token) {
	for _, tok := range tokens {
		if tok.val != nil {
			fmt.Printf("Type: %s = %v\n", tok.Type, tok.val)
		} else {
			fmt.Printf("Type: %s\n", tok.Type)
		}
	}
}


// func main() {
// 	input := " 1 + 2 * ( 5 / 9 ) - 2"
// 	fmt.Println(input)
// 	tokens := Tokenize(input)
// 	PrintTokens(tokens)
// }