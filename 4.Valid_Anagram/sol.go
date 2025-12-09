package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ======================================================================
// 🧠 CHALLENGE: Valid Anagram (Go Version)
// ======================================================================
// Description:
// Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`,
// and `false` otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters of a
// different word or phrase, typically using all the original letters exactly once.
//
// 📋 Rules:
// 1. Return true if t is an anagram of s.
// 2. Return false otherwise.
//
// 💡 Examples:
// - IsAnagram("anagram", "nagaram") => true
// - IsAnagram("rat", "car")         => false
// ======================================================================

// #region [📚 Reference Solutions] (Solutions hidden as requested)

func IsAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 在 Go 中，讀取不存在的 key 會自動回傳 0
	// 這行為跟 Python 的 defaultdict(int) 一模一樣
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	for _, char := range t {
		counts[char]--
		// 優化：如果該字母數量變成負數，表示 t 這個字母比 s 多 (或是 s 根本沒有)
		if counts[char] < 0 {
			return false
		}
	}

	return true
}

// #endregion

// ======================================================================
//
//	#region [✍️ Practice Area]
//	Please write your solution between the markers below.
//
// ======================================================================
// <PRACTICE_START>
func IsAnagram(s string, t string) bool {
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
	s        string
	t        string
	expected bool
}

func runTests() {
	testCases := []TestCase{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
		{"ab", "a", false},
		{"", "", true},
	}

	fmt.Printf("\n🧪 Testing your [IsAnagram] function...\n\n")

	header := fmt.Sprintf("%-10s | %-10s | %-10s | %-10s | Status", "s", "t", "Expected", "Actual")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", len(header)))

	allPass := true

	for _, tc := range testCases {
		result := IsAnagram(tc.s, tc.t)

		isMatch := result == tc.expected
		statusIcon := "✅ PASS"
		if !isMatch {
			statusIcon = "❌ FAIL"
			allPass = false
		}

		sStr := tc.s
		if len(sStr) > 10 {
			sStr = sStr[:7] + "..."
		}
		tStr := tc.t
		if len(tStr) > 10 {
			tStr = tStr[:7] + "..."
		}

		fmt.Printf("%-10s | %-10s | %-10v | %-10v | %s\n",
			sStr, tStr, tc.expected, result, statusIcon)
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
		"func IsAnagram(s string, t string) bool {",
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
