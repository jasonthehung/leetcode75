package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Merge Intervals (Go Version)
// ======================================================================
// Description:
// Given an array of `intervals` where intervals[i] = [start_i, end_i],
// merge all overlapping intervals, and return an array of the
// non-overlapping intervals that cover all the intervals in the input.
//
// 📋 Rules:
// 1. Intervals [a, b] and [c, d] overlap if a <= d and c <= b.
// 2. Intervals touching at boundaries (e.g., [1,4] and [4,5]) are considered overlapping.
// 3. The input intervals might not be sorted.
//
// 💡 Examples:
// - Merge([[1,3],[2,6],[8,10],[15,18]]) => [[1,6],[8,10],[15,18]]
// - Merge([[1,4],[4,5]])                => [[1,5]]
// - Merge([[4,7],[1,4]])                => [[1,7]]
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
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals...)
		} else {
			result[len(result)-1][1] = math.Max()
		}
	}

	return [][]int{}
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
	intervals [][]int
	expected  [][]int
}

func runTests() {
	testCases := []TestCase{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{4, 7}, {1, 4}}, [][]int{{1, 7}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}

	fmt.Printf("\n🧪 Testing your [Merge] function...\n\n")

	header := fmt.Sprintf("%-30s | %-20s | %-20s | Status", "Input Intervals", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		// Deep copy input to avoid modification side effects in test display
		inputCopy := make([][]int, len(tc.intervals))
		for i, v := range tc.intervals {
			row := make([]int, len(v))
			copy(row, v)
			inputCopy[i] = row
		}

		result := Merge(inputCopy)

		isMatch := reflect.DeepEqual(result, tc.expected)
		statusIcon := "✅ PASS"
		if !isMatch {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		// Helper to format slices
		fmtSlice := func(s [][]int) string {
			return fmt.Sprintf("%v", s)
		}

		inpStr := fmtSlice(tc.intervals)
		if len(inpStr) > 28 {
			inpStr = inpStr[:25] + "..."
		}
		expStr := fmtSlice(tc.expected)
		if len(expStr) > 18 {
			expStr = expStr[:15] + "..."
		}
		resStr := fmtSlice(result)
		if len(resStr) > 18 {
			resStr = resStr[:15] + "..."
		}

		fmt.Printf("%-30s | %-20s | %-20s | %s\n",
			inpStr, expStr, resStr, statusIcon)
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
		"func Merge(intervals [][]int) [][]int {",
		"\t// TODO: Implement your solution here.",
		"\treturn [][]int{}",
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
