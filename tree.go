package main

import (
	"fmt"
)

// Tree describes nodes of tree all of which have value (Token) and left and right subtree (nil if leaf)
type Tree struct {
	value       Token
	left, right *Tree
}

// FormTree forms parse tree (/concrete syntax tree/derivation tree) from given tokens
func FormTree(tokens []Token) *Tree {
	//fmt.Println(tokens)

	if len(tokens) == 0 {
		return nil
	}
	/*
		if tokens[0].tokenType == "RPAR" && (tokens[1].tokenType == "TERM" || tokens[1].tokenType == "NEG") {
			t := Tree{}
			t.value = tokens
			t.left = FormTree(tokens[0:lparInd])
			t.right = FormTree(tokens[lparInd:])
		}*/

	lparInd := 0
	for index, value := range tokens {
		if value.tokenType == "LPAR" {
			lparInd = index
			break
		}
	}
	//fmt.Println(lparInd)

	if lparInd > 1 {
		t := Tree{}
		i := 1

		// Count and take in account negations in front of parenthesis
		for lparInd-i > 0 && tokens[lparInd-i].tokenType == "NEG" {
			i++
		}

		t.value = tokens[lparInd-i]
		t.left = FormTree(tokens[0 : lparInd-i])
		t.right = FormTree(tokens[lparInd-i+1:])

		return &t
	} else if lparInd == 1 {
		t := Tree{}
		t.value = tokens[0]
		t.left = FormTree(tokens[1:])
		return &t
	}

	if tokens[lparInd].tokenType == "LPAR" {
		countPars := 0
		rparInd := 0
		for ind, val := range tokens {
			if val.tokenType == "LPAR" {
				countPars++
			} else if val.tokenType == "RPAR" {
				countPars--
				if countPars == 0 {
					rparInd = ind
					break
				}
			}
		}

		if rparInd == 0 {
			fmt.Println("Forming tree failed because cannot find pair for left parenthesis.")
			return nil
		}
		//else
		t := Tree{}

		if rparInd == len(tokens)-1 {
			t.value = tokens[lparInd]
			t.left = FormTree(tokens[lparInd+1 : rparInd])
			return &t
		}
		t.value = tokens[rparInd+1]
		t.left = FormTree(tokens[lparInd : rparInd+1])
		t.right = FormTree(tokens[rparInd+2:])
		return &t
	}

	if tokens[0].tokenType == "NEG" {
		if len(tokens) < 2 {
			fmt.Println("There's something wrong within this token sequence:", tokens)
			return nil
		}
		if tokens[1].tokenType != "TERM" && tokens[1].tokenType != "LPAR" && tokens[1].tokenType != "NEG" {
			fmt.Println("There's something wrong within this token sequence:", tokens)
			return nil
		}
		if len(tokens) > 1 {
			/*	t := Tree{}
				t.value = tokens[2]
				t.left = FormTree(tokens[0:2])
				t.right = FormTree(tokens[3:])
				return &t
			} else if len(tokens) == 2 {*/
			t := Tree{}
			t.value = tokens[0]
			t.left = FormTree(tokens[1:])
			return &t
		}
		//Else
		fmt.Println("There's something wrong within this token sequence:", tokens)
		return nil

	} else if tokens[0].tokenType == "TERM" {
		if len(tokens) > 1 {
			if tokens[1].tokenType == "TERM" || tokens[1].tokenType == "NEG" {
				fmt.Println("There's something wrong within this token sequence:", tokens)
				return nil
			}
			t := Tree{}
			t.value = tokens[1]
			t.left = FormTree(tokens[0:1])
			t.right = FormTree(tokens[2:])
			return &t
		} else if len(tokens) == 1 {
			t := Tree{}
			t.value = tokens[0]
			return &t
		} else {
			fmt.Println("There's something wrong within this token sequence:", tokens)
			return nil
		}
	} else {
		fmt.Println("There's something wrong within this token sequence:", tokens)
		return nil
	}

}

// PrintTree prints tree but it's not pretty
func PrintTree(t *Tree) {
	fmt.Println(t.value)
	if t.left != nil {
		fmt.Print("left ")
		PrintTree(t.left)
	}
	if t.right != nil {
		fmt.Print("right ")
		PrintTree(t.right)
	}
}

/*
rules:
colons(A): A=(A)
imp(A,B): true if
	A and B are true
	A is false
eq(A,B): true if
	A and B are true
	A and B are false
and(A,B): true if
	A and B are true
or(A,B): true if
	A and B are true
	A is true
	B is true
neg(A): true if
	A is false
*/
