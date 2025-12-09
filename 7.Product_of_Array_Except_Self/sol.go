package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Product of Array Except Self (Go Version)
// ======================================================================
// Description:
// Given an integer array `nums`, return an array `answer` such that
// `answer[i]` is equal to the product of all the elements of `nums`
// except `nums[i]`.
//
// 📋 Rules:
// 1. You must write an algorithm that runs in O(n) time.
// 2. You CANNOT use the division operation.
// 3. Follow up: Can you solve it with O(1) extra space complexity?
//    (The output array does not count as extra space).
//
// 💡 Examples:
// - ProductExceptSelf([1,2,3,4])       => [24,12,8,6]
// - ProductExceptSelf([-1,1,0,-3,3])   => [0,0,9,0,0]
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)

func ProductExceptSelf1(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 1. Left Pass
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 2. Right Pass
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		// 關鍵：將右邊累積的積「乘」進結果中
		answer[i] *= rightProduct

		// 更新右邊累積值
		rightProduct *= nums[i]
	}

	return answer
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range n {
		ans[i] = 1
	}

	leftProduct := 1
	for i := range n {
		ans[i] = leftProduct * ans[i]
		leftProduct = leftProduct * nums[i]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] = rightProduct * ans[i]
		rightProduct = rightProduct * nums[i]
	}
	return ans
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
	expected []int
}

func runTests() {
	testCases := []TestCase{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{1, 0}, []int{0, 1}},
	}

	fmt.Printf("\n🧪 Testing your [ProductExceptSelf] function...\n\n")

	header := fmt.Sprintf("%-20s | %-15s | %-15s | Status", "Input nums", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := ProductExceptSelf(tc.nums)

		isMatch := reflect.DeepEqual(result, tc.expected)
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
		expStr := fmtSlice(tc.expected)
		if len(expStr) > 13 {
			expStr = expStr[:10] + "..."
		}
		resStr := fmtSlice(result)
		if len(resStr) > 13 {
			resStr = resStr[:10] + "..."
		}

		fmt.Printf("%-20s | %-15s | %-15s | %s\n",
			numsStr, expStr, resStr, statusIcon)
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
		"func ProductExceptSelf(nums []int) []int {",
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
