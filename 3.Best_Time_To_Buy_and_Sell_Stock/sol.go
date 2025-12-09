package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Best Time to Buy and Sell Stock (Go Version)
// ======================================================================
// Description:
// You are given an array `prices` where `prices[i]` is the price of a
// given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one
// stock and choosing a different day in the future to sell that stock.
//
// 📋 Rules:
// 1. Return the maximum profit you can achieve.
// 2. If you cannot achieve any profit, return 0.
//
// 💡 Examples:
// - MaxProfit([7,1,5,3,6,4]) => 5 (Buy at 1, sell at 6)
// - MaxProfit([7,6,4,3,1])   => 0 (No profit possible)
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)

/**
 * Method: One Pass (Greedy)
 * Fixes:
 * 1. Uses `_, price := range prices` to get the actual value.
 * 2. Removes float64 casting for better performance.
 */
func MaxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// Initialize minPrice to a very large number
	minPrice := math.MaxInt
	maxProfit := 0

	// ⚠️ Crucial Fix: range returns (index, value)
	// We ignore index (_) and capture value (price)
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func MaxProfit(prices []int) int {
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
	prices   []int
	expected int
}

func runTests() {
	testCases := []TestCase{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 4, 1}, 2},
	}

	fmt.Printf("\n🧪 Testing your [MaxProfit] function...\n\n")

	header := fmt.Sprintf("%-20s | %-10s | %-10s | Status", "Input prices", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := MaxProfit(tc.prices)

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

		pricesStr := fmtSlice(tc.prices)
		if len(pricesStr) > 18 {
			pricesStr = pricesStr[:15] + "..."
		}

		fmt.Printf("%-20s | %-10v | %-10v | %s\n",
			pricesStr, tc.expected, result, statusIcon)
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
		"func MaxProfit(prices []int) int {",
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
