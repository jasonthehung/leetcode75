package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Find Minimum in Rotated Sorted Array (Go Version)
// ======================================================================
// Description:
// Suppose an array of length n sorted in ascending order is rotated between
// 1 and n times. Given the sorted rotated array `nums` of unique elements,
// return the minimum element of this array.
//
// You must write an algorithm that runs in O(log n) time.
//
// 📋 Rules:
// 1. nums contains unique integers.
// 2. nums is sorted ascending then rotated.
// 3. Return the minimum value in nums.
//
// 💡 Examples:
// - Input: [3,4,5,1,2]       => Output: 1
// - Input: [4,5,6,7,0,1,2]   => Output: 0
// - Input: [11,13,15,17]     => Output: 11
//
// Constraints:
// - 1 <= n <= 5000
// - -5000 <= nums[i] <= 5000
// - All integers are unique.
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)
func findMin1(nums []int) int {
	// Method: Binary Search (Compare with right boundary)
	// Complexity: Time O(log N) | Space O(1)

	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func FindMin(nums []int) int {
	// TODO: Implement your solution here.
	return 0
}

// <PRACTICE_END>
// #endregion

// ======================================================================
//  #region [🚀 Test Runner & Auto-Reset] (Do not modify below this line)
// ======================================================================

type TestCase struct {
	nums     []int
	expected int
}

func main() {
	runTests()
}

func runTests() {
	testCases := []TestCase{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{2, 1}, 1},
		{[]int{1}, 1},
		{[]int{5, 6, 7, 8, 1, 2, 3, 4}, 1},
	}

	fmt.Printf("\n🧪 Testing your [FindMin] function...\n\n")

	header := fmt.Sprintf("%-40s | %-10s | %-10s | %s", "Input nums", "Expected", "Actual", "Status")
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
			actual = FindMin(tc.nums)
		}()

		if actual != tc.expected {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		inStr := fmt.Sprintf("%v", tc.nums)
		if len(inStr) > 38 {
			inStr = inStr[:35] + "..."
		}
		fmt.Printf("%-40s | %-10d | %-10d | %s\n", inStr, tc.expected, actual, statusIcon)
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
		"func FindMin(nums []int) int {",
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
