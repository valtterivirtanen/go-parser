package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var allowedSymbols = []string{"p", "_", "1", "2", "3", "4", "5", "6", "7", "8", "9", "(", ")", "~", "∼", "&", "v", "-", ">", "<"}

func main() {
	fmt.Println(allowedSymbols)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Give me input:")
	var userInput string
	var x bool
	for scanner.Scan() {
		userInput = scanner.Text()
		userInput = removeSpaces(userInput)
		x = checkUserInput(userInput)
		if x {
			break
		}
		fmt.Println("Invalid input, please give another one")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//fmt.Printf("Input: %v\n", userInput)

	var input []Token
	if x {
		input = TokenizeInput(userInput)
		fmt.Println(input)
	}
	fmt.Println(x)

	fmt.Println("Would you like to test for tautology? (true/false/yes/no)")

	testForTautology := "false"
	for scanner.Scan() {
		testForTautology = scanner.Text()
		break
	}

	if testForTautology == "false" || testForTautology == "no" {
		fmt.Println("Give truth value (false, true, f, t) for each terminal")

		for i := range input {
			if input[i].tokenType == "TERM" && string(input[i].value[0]) == "p" {
				fmt.Print(input[i].value, ": ")
				for scanner.Scan() {
					newVal := scanner.Text()
					if len(newVal) == 0 || (newVal != "false" && newVal != "true" && newVal != "f" && newVal != "t") {
						fmt.Println("Allowed values are 'false'/'true'/'f'/'t'. Please try again.")
						fmt.Print(input[i].value, ": ")
					} else {
						input = replaceAll(string(input[i].value), newVal, input)
						break
					}
				}
			}
			//fmt.Println(input[i])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	if len(input) > 0 {
		if testForTautology == "true" || testForTautology == "yes" {
			fmt.Println("Is tautology:", CheckTautology(input))
			fmt.Println()
		} else {
			t := FormTree(input)

			if t != nil {

				fmt.Println()
				//PrintTree(t)

				finalResult := Translate(*t)
				fmt.Println("Proposition is", finalResult)

			} else {
				fmt.Println("Error, there's something wrong with the tree", t)
			}
		}

	} else {
		fmt.Println("Input length 0, something's wrong!")
	}
}

func checkUserInput(input string) bool {
	found := false
	for _, b := range input {
		found = false
		for _, a := range allowedSymbols {
			if string(b) == a {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Symbol not allowed %v\n", string(b))
			return false
		}
	}
	// Reject empty string ""
	if !found {
		return false
	}
	return true
}

func removeSpaces(s string) string {
	return strings.Join(strings.Fields(s), "")
}

func replaceAll(old string, new string, all []Token) []Token {
	if new == "t" || new == "true" {
		new = "true"
	} else if new == "f" || new == "false" {
		new = "false"
	} else {
		fmt.Println("There's something wrog with the truth values given by user")
		return nil
	}
	for t := range all {
		if all[t].value == old {
			all[t].value = new
		}
	}
	return all
}
