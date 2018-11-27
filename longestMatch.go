package main

/*
   Parenthesis problems
   Balancing brackets
   Find longest valid match of paranthesis
*/

import (
	"fmt"
)

func main() {
	a := []string{"{()}"}
	fmt.Println(a, ":", isParenthesisValid(a))

	fmt.Println(isParenthesisValid("{()}"))
	fmt.Println(isParenthesisValid(")("))
	fmt.Println(isParenthesisValid("(())"))
	fmt.Println(isParenthesisValid("())"))
	fmt.Println(isParenthesisValid("))("))
	fmt.Println(isParenthesisValid("[{]}"))
	fmt.Println(isParenthesisValid("))"))
	fmt.Println(4, longestMatch("()((())")) // 4
	fmt.Println(4, longestMatch("()(()()")) // 4
	fmt.Println(4, longestMatch("()())()")) // 4
	fmt.Println(4, longestMatch("(())"))    // 4
	fmt.Println(4, longestMatch("()()"))    // 4
	fmt.Println(2, longestMatch(")()("))    // 2
	fmt.Println(2, longestMatch("())"))     // 2
	fmt.Println(2, longestMatch("(()"))     // 2
	fmt.Println(0, longestMatch(")("))      // 0

}

func longestMatch(s string) int {
	stack := make([]byte, 0, len(s)/2)
	count, lastMatch, max := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			continue
		}
		stack = stack[:0]
		count = 0
		for j := i; j < len(s); j++ {
			//	fmt.Println(i, j, s[i:])
			if len(stack) == 0 {
				lastMatch = count
				//		fmt.Println("last Match", count)
			}
			if s[j] == '(' {
				stack = push(stack, s[j])
			} else {
				if len(stack) > 0 {
					_, stack = pop(stack)
					count = count + 2
				} else {
					break
				}
			}
		}
		if len(stack) != 0 {
			count = lastMatch
		}
		if count > max {
			max = count
		}
	}
	if count > max {
		max = count
	}

	return max

}

func isParenthesisValid(parentheses string) bool {
	stack := make([]rune, 0, len(parentheses)/2)
	var p rune
	for _, c := range parentheses {
		if c == '{' || c == '(' || c == '[' {
			stack = push(stack, c)
		} else {
			// fmt.Printf("%s", stack)
			// If stack is empty, parenthesis is missing
			if len(stack) == 0 {
				return false
			}
			// fmt.Printf("%s",p)

			p, stack = pop(stack)
			if p != matches[c] {
				return false
			}
		}
	}
	return true
}

func push(stack []byte, c byte) []byte {
	stack = append(stack, c)
	return stack
}

func pop(stack []byte) (byte, []byte) {
	e := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return e, stack
}
