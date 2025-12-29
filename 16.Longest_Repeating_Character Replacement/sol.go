package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Longest Repeating Character Replacement (Go Version)
// ======================================================================
// Description:
// You are given a string s and an integer k. You can choose any character
// of the string and change it to any other uppercase English character.
// You can perform this operation at most k times.
//
// Return the length of the longest substring containing the same letter
// you can get after performing the above operations.
//
// Constraints:
// - 1 <= len(s) <= 1e5
// - s consists of only uppercase English letters
// - 0 <= k <= len(s)
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)
func characterReplacement1(s string, k int) int {
	// Method: Sliding Window + Frequency Count
	// Complexity: Time O(N) | Space O(1) (alphabet size fixed at 26)

	freq := make([]int, 26)
	l := 0
	maxFreq := 0
	best := 0

	for r := 0; r < len(s); r++ {
		idx := int(s[r] - 'A')
		freq[idx]++
		if freq[idx] > maxFreq {
			maxFreq = freq[idx]
		}

		for (r-l+1)-maxFreq > k {
			leftIdx := int(s[l] - 'A')
			freq[leftIdx]--
			l++
		}

		winLen := r - l + 1
		if winLen > best {
			best = winLen
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
func CharacterReplacement(s string, k int) int {
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
	k        int
	expected int
}

func main() {
	runTests()
}

func runTests() {
	testCases := []TestCase{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AAAA", 0, 4},
		{"ABCDE", 1, 2},
		{"BAAAB", 2, 5},
		{"", 3, 0}, // not in constraints, but helps robustness
	}

	fmt.Printf("\n🧪 Testing your [CharacterReplacement] function...\n\n")

	header := fmt.Sprintf("%-30s | %-10s | %-10s | %s", "Input (s,k)", "Expected", "Actual", "Status")
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
			actual = CharacterReplacement(tc.s, tc.k)
		}()

		if actual != tc.expected {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		inStr := fmt.Sprintf("(\"%s\", %d)", tc.s, tc.k)
		if len(inStr) > 28 {
			inStr = inStr[:25] + "..."
		}
		fmt.Printf("%-30s | %-10d | %-10d | %s\n", inStr, tc.expected, actual, statusIcon)
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
		"func CharacterReplacement(s string, k int) int {",
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
