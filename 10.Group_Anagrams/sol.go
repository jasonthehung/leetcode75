package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Group Anagrams (Go Version)
// ======================================================================
// Description:
// Given an array of strings `strs`, group the anagrams together.
// You can return the answer in any order.
//
// 📋 Rules:
// 1. Return a slice of slices, where each inner slice contains words that are anagrams.
// 2. The order of the groups and the order of words within groups does not matter.
//
// 💡 Examples:
// - GroupAnagrams(["eat","tea","tan","ate","nat","bat"])
//   => [["bat"],["nat","tan"],["ate","eat","tea"]]
// - GroupAnagrams([""])  => [[""]]
// - GroupAnagrams(["a"]) => [["a"]]
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)
// (Focus on implementing your own logic in the Practice Area below!)
// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func GroupAnagrams(strs []string) [][]string {
	// TODO: Implement your solution here.
	return [][]string{}
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
	strs     []string
	expected [][]string
}

func runTests() {
	testCases := []TestCase{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	fmt.Printf("\n🧪 Testing your [GroupAnagrams] function...\n\n")

	header := fmt.Sprintf("%-30s | %-10s", "Input strs", "Status")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", 45))

	allPass := true

	for _, tc := range testCases {
		result := GroupAnagrams(tc.strs)

		// Canonicalize results for comparison
		canonicalize := func(input [][]string) string {
			if len(input) == 0 {
				return "[]"
			}
			// Copy to sort safely
			copySlice := make([][]string, len(input))
			for i, v := range input {
				innerCopy := make([]string, len(v))
				copy(innerCopy, v)
				sort.Strings(innerCopy)
				copySlice[i] = innerCopy
			}

			// Sort the groups themselves for deterministic comparison
			sort.Slice(copySlice, func(i, j int) bool {
				// Sort by the first element of each group (since inner groups are sorted)
				if len(copySlice[i]) > 0 && len(copySlice[j]) > 0 {
					return copySlice[i][0] < copySlice[j][0]
				}
				return len(copySlice[i]) < len(copySlice[j])
			})
			return fmt.Sprintf("%v", copySlice)
		}

		resStr := canonicalize(result)
		expStr := canonicalize(tc.expected)

		isMatch := resStr == expStr
		statusIcon := "✅ PASS"
		if !isMatch {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		// Helper to format string for display
		inputStr := fmt.Sprintf("%v", tc.strs)
		if len(inputStr) > 28 {
			inputStr = inputStr[:25] + "..."
		}

		fmt.Printf("%-30s | %s\n", inputStr, statusIcon)
		if !isMatch {
			// Print details if failed, as simple table might not show enough
			fmt.Printf("   Expected: %v\n", tc.expected)
			fmt.Printf("   Actual:   %v\n", result)
		}
	}

	fmt.Println(strings.Repeat("-", 45))

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
		"func GroupAnagrams(strs []string) [][]string {",
		"\t// TODO: Implement your solution here.",
		"\treturn [][]string{}",
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
