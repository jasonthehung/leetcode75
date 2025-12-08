package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: 3Sum (Go Version)
// ======================================================================
// Description:
// Given an integer array `nums`, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
// 📋 Rules:
// 1. Return a slice of slices, where each inner slice is a triplet.
// 2. The order of the output and the order of the triplets does not matter.
// 3. Elements in a triplet must sum to 0.
//
// 💡 Examples:
// - ThreeSum([-1,0,1,2,-1,-4]) => [[-1,-1,2], [-1,0,1]]
// - ThreeSum([0,1,1])          => []
// - ThreeSum([0,0,0])          => [[0,0,0]]
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)

/**
 * Method: Sorting + Two Pointers
 * Complexity: Time O(N^2) | Space O(1)
 */
func ThreeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// Optimization: If current number is positive, the rest will be larger
		if nums[i] > 0 {
			break
		}

		// 1. Skip duplicate 'i'
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// Found a triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// 2. Skip duplicate 'left'
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// 3. Skip duplicate 'right'
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func ThreeSum(nums []int) [][]int {
	// TODO: Implement your solution here.
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
	nums     []int
	expected [][]int
}

func runTests() {
	testCases := []TestCase{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	fmt.Printf("\n🧪 Testing your [ThreeSum] function...\n\n")

	header := fmt.Sprintf("%-25s | %-20s | %-20s | Status", "Input nums", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := ThreeSum(tc.nums)

		// Canonicalize results for comparison
		canonicalize := func(input [][]int) string {
			if len(input) == 0 {
				return "[]"
			}
			// Copy to sort safely
			copySlice := make([][]int, len(input))
			for i, v := range input {
				innerCopy := make([]int, len(v))
				copy(innerCopy, v)
				sort.Ints(innerCopy)
				copySlice[i] = innerCopy
			}

			// Sort the list of triplets
			sort.Slice(copySlice, func(i, j int) bool {
				for k := 0; k < len(copySlice[i]) && k < len(copySlice[j]); k++ {
					if copySlice[i][k] != copySlice[j][k] {
						return copySlice[i][k] < copySlice[j][k]
					}
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
		formatDisp := func(s string) string {
			if len(s) > 18 {
				return s[:15] + "..."
			}
			return s
		}

		numsStr := fmt.Sprintf("%v", tc.nums)
		if len(numsStr) > 23 {
			numsStr = numsStr[:20] + "..."
		}

		fmt.Printf("%-25s | %-20s | %-20s | %s\n",
			numsStr, formatDisp(expStr), formatDisp(resStr), statusIcon)
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
		"func ThreeSum(nums []int) [][]int {",
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
