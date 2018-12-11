package main

import (
	"fmt"
	"strconv"
)

// Translate takes tree, determines it's value(boolean), and returns it
func Translate(t Tree) bool {

	switch t.value.tokenType {
	case "TERM":
		s, err := strconv.ParseBool(t.value.value)
		if err == nil {
			return s
		}
	case "NEG":
		if t.left != nil {
			return !Translate(*t.left)
		}
		return !Translate(*t.right)
	case "EQUAL":
		return areEqual(Translate(*t.left), Translate(*t.right))
	case "IMP":
		return !Translate(*t.left) || Translate(*t.right)
	case "AND":
		return Translate(*t.left) && Translate(*t.right)
	case "OR":
		return Translate(*t.left) || Translate(*t.right)
	case "LPAR":
		return Translate(*t.left)
	case "RPAR":
		return Translate(*t.left) || Translate(*t.right)
	default:
		fmt.Println("Invalid tokenType:", t.value.tokenType)
		return false
	}

	return false
}

func areEqual(a bool, b bool) bool {
	return (!a || b) && (!b || a)
}

// CheckTautology taken tokenized user input and return whether it's tautology or not
func CheckTautology(input []Token) bool {
	newVal := "true"
	for i := range input {
		if input[i].tokenType == "TERM" {
			input = replaceAll(string(input[i].value), newVal, input)
			break
		}

	}

	result := processInput(input)

	newVal = "false"
	for i := range input {
		if input[i].tokenType == "TERM" && result {
			input = replaceAll(string(input[i].value), newVal, input)
			result = processInput(input)
			break
		}

	}
	newVal = "true"
	for i := range input {
		if input[i].tokenType == "TERM" && result {
			input = replaceAll(string(input[i].value), newVal, input)
			result = processInput(input)
			break
		}
	}
	return result
}

func processInput(input []Token) bool {
	t := FormTree(input)
	finalResult := false
	if t != nil {

		fmt.Println()
		//PrintTree(t)

		finalResult = Translate(*t)

	} else {
		fmt.Println("Tree is nil!")
	}
	return finalResult
}
