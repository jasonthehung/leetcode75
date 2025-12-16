package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Maximum Product Subarray (Go Version)
// ======================================================================
// Description:
// Given an integer array `nums`, find a subarray that has the largest
// product, and return the product.
//
// 📋 Rules:
// 1. A subarray is a contiguous part of an array.
// 2. Return the maximum product found.
// 3. Handle negative numbers (negative * negative = positive) and zeros.
//
// 💡 Examples:
// - MaxProduct([2,3,-2,4]) => 6  ([2,3])
// - MaxProduct([-2,0,-1])  => 0  (The result cannot be 2, because [-2,-1] is not a subarray)
// - MaxProduct([-2,3,-4])  => 24 ([-2,3,-4])
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
func MaxProduct(nums []int) int {
	// TODO: Implement your solution here.
	return 0
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
	expected int
}

func runTests() {
	testCases := []TestCase{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
		{[]int{-4, -3, -2}, 12},
		{[]int{0, 2}, 2},
		{[]int{-2}, -2},
		{[]int{2, -5, -2, -4, 3}, 24},
	}

	fmt.Printf("\n🧪 Testing your [MaxProduct] function...\n\n")

	header := fmt.Sprintf("%-30s | %-10s | %-10s | Status", "Input nums", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := MaxProduct(tc.nums)

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
		"func MaxProduct(nums []int) int {",
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
