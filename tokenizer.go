package main

import (
	"fmt"
	"regexp"
)

// Token is an object which has type and value properties
type Token struct {
	tokenType string
	value     string
}

// TokenizeInput converts string into a slice of tokens
func TokenizeInput(input string) []Token {
	// Let's ensure that every left parenthesis has right pair
	parenthesis := 0
	var tokens []Token
	input += "E"
	for input != "" {
		i := 0

		switch string(input[i]) {
		case "(":
			// If there are parenthesis the first one has to be "("
			if parenthesis < 0 {
				fmt.Println("There's something wrong with the parenthesis")
				return []Token{}
			}
			parenthesis++
			tokens = append(tokens, Token{"LPAR", string(input[i])})
			input = input[1:]
		case ")":
			parenthesis--
			tokens = append(tokens, Token{"RPAR", string(input[i])})
			input = input[1:]
		case "&":
			tokens = append(tokens, Token{"AND", string(input[i])})
			input = input[1:]
		case "v":
			tokens = append(tokens, Token{"OR", string(input[i])})
			input = input[1:]
		// There was a problem with tilde. That's why it's separated in 2 cases.
		case "∼":
			// This tilde is copied from moodle.
			fmt.Println("Wrong kind of tilde. Copy and use this one instead ~")
			return []Token{}
		case "~":
			tokens = append(tokens, Token{"NEG", string(input[i])})
			input = input[1:]
		// This case is actually <->
		case "<":
			if string(input[1]) == "-" && string(input[2]) == ">" {
				tokens = append(tokens, Token{"EQUAL", "<->"})
			} else {
				fmt.Printf(`"<" should be followed by characters "-" and ">". Instead there was characters "%v" and "%v"
`, string(input[1]), string(input[2]))
				return []Token{}
			}
			input = input[3:]
		case "-":
			if string(input[1]) == ">" {
				tokens = append(tokens, Token{"IMP", "->"})
			} else {
				fmt.Printf(`"-" should be followed by character ">". Instead there was character "%v"
`, string(input[1]))
				return []Token{}
			}
			input = input[2:]
		case "p":
			// Prevent out of bounds with limiter
			limit := 5
			if len(string(input)) < 5 {
				limit = len(string(input))
			}

			match, _ := regexp.Compile("[0-9]+")
			digits := match.FindString(string(input[2:limit]))
			if string(input[1]) == "_" && len(digits) > 0 {
				tokens = append(tokens, Token{"TERM", string(input[i]) + string(input[1]) + digits})
			} else {
				fmt.Printf(`"p" should be followed by characters "_" AND digits. Instead there was characters "%v" and "%v"
`, string(input[1]), digits)
				return []Token{}
			}
			input = input[2+len(digits):]
		// End of string reached
		case "E":
			input = input[1:]
			break
		default:
			fmt.Println("Unidentified symbol in user input", string(input[0]))
			return nil
		}
	}

	if parenthesis == 0 && len(input) == 0 {
		return tokens
	}
	// Else
	fmt.Println("There's something wrong with the input")
	fmt.Println("Number of pairless parenthesis:", parenthesis, "Length of false input:", len(input))
	return nil
}
