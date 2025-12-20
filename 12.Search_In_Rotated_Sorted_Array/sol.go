package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Search in Rotated Sorted Array (Go Version)
// ======================================================================
// Description:
// There is an integer array `nums` sorted in ascending order (with distinct values).
// Prior to being passed to your function, `nums` is possibly rotated at an
// unknown index k.
//
// Given the array `nums` after the possible rotation and an integer `target`,
// return the index of `target` if it is in `nums`, or -1 if it is not.
//
// 📋 Rules:
// 1. You must write an algorithm with O(log n) runtime complexity.
// 2. Return the index of the target, or -1.
//
// 💡 Examples:
// - Search([4,5,6,7,0,1,2], 0) => 4
// - Search([4,5,6,7,0,1,2], 3) => -1
// - Search([1], 0)             => -1
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
func Search(nums []int, target int) int {
	// TODO: Implement your solution here.
	return -1
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
	nums     []int
	target   int
	expected int
}

func runTests() {
	testCases := []TestCase{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{3, 1}, 1, 1},
		{[]int{5, 1, 3}, 5, 0},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4},
	}

	fmt.Printf("\n🧪 Testing your [Search] function...\n\n")

	header := fmt.Sprintf("%-25s | %-6s | %-10s | %-10s | Status", "Input nums", "Target", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := Search(tc.nums, tc.target)

		isMatch := result == tc.expected
		statusIcon := "✅ PASS"
		if !isMatch {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		// Helper to format slices
		fmtSlice := func(s []int) string {
			return fmt.Sprintf("%v", s)
		}

		numsStr := fmtSlice(tc.nums)
		if len(numsStr) > 23 {
			numsStr = numsStr[:20] + "..."
		}

		fmt.Printf("%-25s | %-6d | %-10v | %-10v | %s\n",
			numsStr, tc.target, tc.expected, result, statusIcon)
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
		"func Search(nums []int, target int) int {",
		"\t// TODO: Implement your solution here.",
		"\treturn -1",
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
