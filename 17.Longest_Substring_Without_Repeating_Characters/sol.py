import os
import sys
from typing import Dict

# ======================================================================
# 🧠 CHALLENGE: Longest Substring Without Repeating Characters (Python)
# ======================================================================
# Description:
# Given a string s, find the length of the longest substring without
# duplicate characters.
#
# 📋 Rules:
# 1. Substring must be contiguous.
# 2. No repeated characters inside the chosen substring.
# 3. Return the maximum length.
#
# 💡 Examples:
# - Input: "abcabcbb" => Output: 3  ("abc")
# - Input: "bbbbb"    => Output: 1  ("b")
# - Input: "pwwkew"   => Output: 3  ("wke")
#
# Constraints:
# - 0 <= len(s) <= 5 * 10^4
# - s consists of English letters, digits, symbols, and spaces.
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def length_of_longest_substring1(s: str) -> int:
    """
    Method: Sliding Window + Last Seen Index (Hash Map)
    Complexity: Time O(N) | Space O(min(N, charset))
    """
    last = {}  # char -> last index
    l = 0
    best = 0

    for r, ch in enumerate(s):
        if ch in last and last[ch] >= l:
            l = last[ch] + 1
        last[ch] = r

        win_len = r - l + 1
        if win_len > best:
            best = win_len

    return best


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def length_of_longest_substring(s: str) -> int:
    last = {}
    l = 0
    res = 0

    for r, ch in enumerate(s):
        if ch in last and last[ch] >= l:
            l = last[ch] + 1

        last[ch] = r

        res = max(res, r - l + 1)

    return res


# <PRACTICE_END>
# endregion

# ======================================================================
#  region [🚀 Test Runner & Auto-Reset] (Do not modify below this line)
# ======================================================================


def reset_practice_area():
    print("\n🔄 Resetting Practice Area to default state...")
    MARKER_START = "# <PRACTICE_" + "START>"
    MARKER_END = "# <PRACTICE_" + "END>"

    default_code_lines = [
        "def length_of_longest_substring(s: str) -> int:\n",
        "    # TODO: Implement your solution here.\n",
        "    return 0\n",
    ]

    try:
        current_file = __file__
        with open(current_file, "r", encoding="utf-8") as f:
            lines = f.readlines()

        start_idx = -1
        end_idx = -1

        for i, line in enumerate(lines):
            if MARKER_START in line:
                start_idx = i
            if MARKER_END in line:
                end_idx = i

        if start_idx == -1 or end_idx == -1 or start_idx >= end_idx:
            print("⚠️ Error: Markers not found. Reset cancelled.")
            return

        new_content = lines[: start_idx + 1] + default_code_lines + lines[end_idx:]

        with open(current_file, "w", encoding="utf-8") as f:
            f.writelines(new_content)

        print("✨ Reset complete! The file is ready for a fresh start.")
    except Exception as e:
        print(f"⚠️ Reset failed: {e}")


if __name__ == "__main__":
    test_cases = [
        ("abcabcbb", 3),
        ("bbbbb", 1),
        ("pwwkew", 3),
        ("", 0),
        (" ", 1),
        ("dvdf", 3),  # "vdf"
        ("abba", 2),  # "ab" or "ba"
        ("tmmzuxt", 5),  # "mzuxt"
    ]

    print(f"\n🧪 Testing your [length_of_longest_substring] function...\n")

    header = f"{'Input s':<25} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for s, expected in test_cases:
        try:
            actual = length_of_longest_substring(s)
        except Exception:
            actual = "Error"

        is_match = actual == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        in_str = repr(s)
        if len(in_str) > 23:
            in_str = in_str[:20] + "..."
        print(f"{in_str:<25} | {str(expected):<10} | {str(actual):<10} | {status_icon}")

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")
# endregion
