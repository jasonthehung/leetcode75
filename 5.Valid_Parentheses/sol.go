package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Valid Parentheses (Go Version)
// ======================================================================
// Description:
// Given a string `s` containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// 3. Every close bracket has a corresponding open bracket of the same type.
//
// 📋 Rules:
// - Return true if the string is valid.
// - Return false otherwise.
//
// 💡 Examples:
// - IsValid("()")       => true
// - IsValid("()[]{}")   => true
// - IsValid("(]")       => false
// - IsValid("([])")     => true
// - IsValid("([)]")     => false
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)

func IsValid1(s string) bool {
	// Stack of runes (characters)
	stack := []rune{}

	// Map: Closing -> Opening
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		// Check if the character is a closing bracket
		if expectedOpen, isClose := pairs[char]; isClose {
			// Case 1: Stack is empty but we found a closing bracket -> Invalid
			if len(stack) == 0 {
				return false
			}

			// Case 2: Pop from stack
			lastIndex := len(stack) - 1
			topElement := stack[lastIndex]
			stack = stack[:lastIndex] // Remove last element

			// Case 3: Check if they match
			if topElement != expectedOpen {
				return false
			}
		} else {
			// It is an opening bracket, push to stack
			stack = append(stack, char)
		}
	}

	// If stack is empty, it means all brackets were matched correctly
	return len(stack) == 0
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func IsValid(s string) bool {
	st := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if _, exist := pairs[char]; exist {
			if len(st) == 0 {
				return false
			}

			top := st[len(st)-1]
			st = st[:len(st)-1]

			if top != pairs[char] {
				return false
			}
		} else {
			st = append(st, char)
		}
	}
	return len(st) == 0
}

// <PRACTICE_END>
// #endregion

// ======================================================================
//  #region [🚀 Test Runner & Auto-Reset] (Do not modify below this line)
// ======================================================================

func main() {
	runTests()
}

type TestCase struct {
	s        string
	expected bool
}

func runTests() {
	testCases := []TestCase{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
		{"]", false},
		{"", true},
	}

	fmt.Printf("\n🧪 Testing your [IsValid] function...\n\n")

	header := fmt.Sprintf("%-20s | %-10s | %-10s | Status", "Input s", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := IsValid(tc.s)

		isMatch := result == tc.expected
		statusIcon := "✅ PASS"
		if !isMatch {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		sStr := tc.s
		if len(sStr) > 20 {
			sStr = sStr[:17] + "..."
		}

		fmt.Printf("%-20s | %-10v | %-10v | %s\n",
			sStr, tc.expected, result, statusIcon)
	}

	fmt.Println(strings.Repeat("-", len(header)))

	if allPass {
		fmt.Println("\n🎉 Fantastic! All test cases passed.")
		resetPracticeArea()
	} else {
		fmt.Println("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")
	}
}

func resetPracticeArea() {
	fmt.Println("\n🔄 Resetting Practice Area to default state...")

	markerStart := "// <PRACTICE_" + "START>"
	markerEnd := "// <PRACTICE_" + "END>"

	defaultCode := []string{
		"func IsValid(s string) bool {",
		"\t// TODO: Implement your solution here.",
		"\treturn false",
		"}",
	}

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("⚠️ Error: Could not determine file path. Reset cancelled.")
		return
	}

	content, err := os.ReadFile(currentFile)
	if err != nil {
		fmt.Printf("⚠️ Error reading file: %v\n", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	startIdx := -1
	endIdx := -1

	for i, line := range lines {
		if strings.Contains(line, markerStart) {
			startIdx = i
		} else if strings.Contains(line, markerEnd) {
			endIdx = i
		}
	}

	if startIdx == -1 || endIdx == -1 || startIdx >= endIdx {
		fmt.Println("⚠️ Error: Markers not found or invalid. Reset cancelled.")
		return
	}

	newLines := make([]string, 0)
	newLines = append(newLines, lines[:startIdx+1]...)
	newLines = append(newLines, defaultCode...)
	newLines = append(newLines, lines[endIdx:]...)

	output := strings.Join(newLines, "\n")
	err = os.WriteFile(currentFile, []byte(output), 0644)
	if err != nil {
		fmt.Printf("⚠️ Error writing file: %v\n", err)
		return
	}

	fmt.Println("✨ Reset complete! The file is ready for a fresh start.")
}

// #endregion
