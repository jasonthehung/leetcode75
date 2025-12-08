package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Two Sum (Go Version)
// ======================================================================
// Description:
// Given an array of integers `nums` and an integer `target`, return
// indices of the two numbers such that they add up to `target`.
//
// 📋 Rules:
// 1. You may assume that each input would have EXACTLY one solution.
// 2. You may not use the same element twice.
// 3. You can return the answer in any order.
// 4. Follow-up: Can you come up with an algorithm that is less than
//    O(n^2) time complexity?
//
// 💡 Examples:
// - TwoSum([2,7,11,15], 9) => [0, 1]
// - TwoSum([3,2,4], 6)     => [1, 2]
// - TwoSum([3,3], 6)       => [0, 1]
// ======================================================================

// #region [📚 Reference Solutions]

/**
 * Method: One-pass Hash Map
 * Complexity: Time O(N) | Space O(N)
 */
func TwoSum1(nums []int, target int) []int {
	// Create a map to store value -> index
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		// Check if complement exists in map
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}

		// Store current number and index
		seen[num] = i
	}

	return nil // Should not be reached
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func TwoSum(nums []int, target int) []int {
	// TODO: Implement your solution here.
	return []int{}
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
	expected []int
}

func runTests() {
	testCases := []TestCase{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	fmt.Printf("\n🧪 Testing your [TwoSum] function...\n\n")

	header := fmt.Sprintf("%-20s | %-6s | %-10s | %-10s | Status", "Input nums", "Target", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := TwoSum(tc.nums, tc.target)

		// Create copies to sort for comparison without modifying originals (though these are local)
		resCopy := make([]int, len(result))
		copy(resCopy, result)
		expCopy := make([]int, len(tc.expected))
		copy(expCopy, tc.expected)

		sort.Ints(resCopy)
		sort.Ints(expCopy)

		isMatch := reflect.DeepEqual(resCopy, expCopy)
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
		if len(numsStr) > 18 {
			numsStr = numsStr[:15] + "..."
		}

		fmt.Printf("%-20s | %-6d | %-10s | %-10s | %s\n",
			numsStr, tc.target, fmtSlice(tc.expected), fmtSlice(result), statusIcon)
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
		"func TwoSum(nums []int, target int) []int {",
		"\t// TODO: Implement your solution here.",
		"\treturn []int{}",
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
