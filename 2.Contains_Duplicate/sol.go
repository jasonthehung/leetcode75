package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Contains Duplicate (Go Version)
// ======================================================================
// Description:
// Given an integer array `nums`, return `true` if any value appears
// at least twice in the array, and return `false` if every element
// is distinct.
//
// 📋 Rules:
// 1. Return true if duplicates exist.
// 2. Return false if all elements are unique.
//
// 💡 Examples:
// - ContainsDuplicate([1,2,3,1])               => true
// - ContainsDuplicate([1,2,3,4])               => false
// - ContainsDuplicate([1,1,1,3,3,4,3,2,4,2])   => true
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)
/**
 * Method: Hash Map as Set
 * Why: Go has no built-in 'Set'. We use map[int]struct{} instead.
 * struct{} takes 0 bytes of storage, making it efficient.
 */
func ContainsDuplicate1(nums []int) bool {
	// 1. Create a map using empty struct to minimize memory usage
	seen := make(map[int]struct{})

	for _, num := range nums {
		// 2. Check if key exists in map
		// The `_` ignores the value, `exists` is a boolean
		if _, exists := seen[num]; exists {
			return true
		}

		// 3. Add to map (Set)
		seen[num] = struct{}{}
	}

	return false
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func ContainsDuplicate(nums []int) bool {
	// TODO: Implement your solution here.
	return false
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
	expected bool
}

func runTests() {
	testCases := []TestCase{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	fmt.Printf("\n🧪 Testing your [ContainsDuplicate] function...\n\n")

	header := fmt.Sprintf("%-30s | %-10s | %-10s | Status", "Input nums", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := ContainsDuplicate(tc.nums)

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
		if len(numsStr) > 28 {
			numsStr = numsStr[:25] + "..."
		}

		fmt.Printf("%-30s | %-10v | %-10v | %s\n",
			numsStr, tc.expected, result, statusIcon)
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
		"func ContainsDuplicate(nums []int) bool {",
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
