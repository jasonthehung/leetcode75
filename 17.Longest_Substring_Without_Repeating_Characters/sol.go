package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Longest Substring Without Repeating Characters (Go)
// ======================================================================
// Description:
// Given a string s, find the length of the longest substring without
// duplicate characters.
//
// 📋 Rules:
// 1. Substring must be contiguous.
// 2. No repeated characters inside the chosen substring.
// 3. Return the maximum length.
//
// Constraints:
// - 0 <= len(s) <= 5 * 10^4
// - s consists of English letters, digits, symbols, and spaces.
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)
func lengthOfLongestSubstring1(s string) int {
	// Method: Sliding Window + Last Seen Index (Hash Map)
	// Complexity: Time O(N) | Space O(min(N, charset))

	last := make(map[byte]int, 128) // byte -> last index
	for i := 0; i < 128; i++ {
		last[byte(i)] = -1
	}

	l := 0
	best := 0

	for r := 0; r < len(s); r++ {
		ch := s[r]
		if last[ch] >= l {
			l = last[ch] + 1
		}
		last[ch] = r

		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func LengthOfLongestSubstring(s string) int {
	// TODO: Implement your solution here.
	return 0
}
// <PRACTICE_END>
// #endregion

// ======================================================================
//  #region [🚀 Test Runner & Auto-Reset] (Do not modify below this line)
// ======================================================================

type TestCase struct {
	s        string
	expected int
}

func main() {
	runTests()
}

func runTests() {
	testCases := []TestCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"dvdf", 3},    // "vdf"
		{"abba", 2},    // "ab" or "ba"
		{"tmmzuxt", 5}, // "mzuxt"
	}

	fmt.Printf("\n🧪 Testing your [LengthOfLongestSubstring] function...\n\n")

	header := fmt.Sprintf("%-25s | %-10s | %-10s | %s", "Input s", "Expected", "Actual", "Status")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		actual := -999999
		statusIcon := "✅ PASS"

		func() {
			defer func() {
				if r := recover(); r != nil {
					actual = -999999
					statusIcon = "❌ FAIL"
					allPass = false
				}
			}()
			actual = LengthOfLongestSubstring(tc.s)
		}()

		if actual != tc.expected {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		inStr := fmt.Sprintf("%q", tc.s)
		if len(inStr) > 23 {
			inStr = inStr[:20] + "..."
		}
		fmt.Printf("%-25s | %-10d | %-10d | %s\n", inStr, tc.expected, actual, statusIcon)
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
		"func LengthOfLongestSubstring(s string) int {",
		"\t// TODO: Implement your solution here.",
		"\treturn 0",
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
