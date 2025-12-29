package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Container With Most Water (Go Version)
// ======================================================================
// Description:
// Given an integer array `height` of length n, there are n vertical lines
// drawn such that the endpoints of the i-th line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the
// container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice:
// - You may not slant the container.
//
// 📋 Rules:
// 1. Choose two indices i < j.
// 2. Area = (j - i) * min(height[i], height[j]).
// 3. Return the maximum possible area.
//
// 💡 Examples:
// - Input: [1,8,6,2,5,4,8,3,7] => Output: 49
// - Input: [1,1]               => Output: 1
//
// Constraints:
// - 2 <= n <= 1e5
// - 0 <= height[i] <= 1e4
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)
func maxArea1(height []int) int {
	// Method: Two Pointers (Greedy)
	// Complexity: Time O(N) | Space O(1)
	l, r := 0, len(height)-1
	best := 0

	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		area := h * (r - l)
		if area > best {
			best = area
		}

		// Move the pointer with smaller height
		if height[l] <= height[r] {
			l++
		} else {
			r--
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
func MaxArea(height []int) int {
	// TODO: Implement your solution here.
	return -1
}
// <PRACTICE_END>
// #endregion

// ======================================================================
//  #region [🚀 Test Runner & Auto-Reset] (Do not modify below this line)
// ======================================================================

type TestCase struct {
	height   []int
	expected int
}

func main() {
	runTests()
}

func runTests() {
	testCases := []TestCase{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{0, 0}, 0},
	}

	fmt.Printf("\n🧪 Testing your [MaxArea] function...\n\n")

	header := fmt.Sprintf("%-35s | %-10s | %-10s | %s", "Input height", "Expected", "Actual", "Status")
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
			actual = MaxArea(tc.height)
		}()

		if actual != tc.expected {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		inStr := fmt.Sprintf("%v", tc.height)
		if len(inStr) > 33 {
			inStr = inStr[:30] + "..."
		}
		fmt.Printf("%-35s | %-10d | %-10d | %s\n", inStr, tc.expected, actual, statusIcon)
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
		"func MaxArea(height []int) int {",
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
