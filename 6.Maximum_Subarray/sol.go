package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Maximum Subarray (Go Version)
// ======================================================================
// Description:
// Given an integer array `nums`, find the subarray with the largest sum,
// and return its sum.
//
// 📋 Rules:
// 1. A subarray is a contiguous part of an array.
// 2. Return the maximum sum found.
// 3. Handle cases with negative numbers correctly.
//
// 💡 Examples:
// - MaxSubArray([-2,1,-3,4,-1,2,1,-5,4]) => 6
// - MaxSubArray([1])                     => 1
// - MaxSubArray([5,4,-1,7,8])            => 23
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)

/**
 * Method: Kadane's Algorithm
 * Complexity: Time O(N) | Space O(1)
 */
func MaxSubArray1(nums []int) int {
	// Initialize with the first element
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Logic: curSum = max(num, curSum + num)
		// If adding the previous sum makes the result smaller than num itself
		// (which happens if curSum is negative), we discard the previous sum.
		if currentSum < 0 {
			currentSum = num
		} else {
			currentSum += num
		}

		// Update global max
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func MaxSubArray(nums []int) int {
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
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1}, -1},     // Edge case: single negative number
		{[]int{-2, -1}, -1}, // Edge case: all negatives
	}

	fmt.Printf("\n🧪 Testing your [MaxSubArray] function...\n\n")

	header := fmt.Sprintf("%-30s | %-10s | %-10s | Status", "Input nums", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := MaxSubArray(tc.nums)

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
		"func MaxSubArray(nums []int) int {",
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
