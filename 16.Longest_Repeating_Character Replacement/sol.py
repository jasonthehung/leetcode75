import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Longest Repeating Character Replacement (Python Version)
# ======================================================================
# Description:
# You are given a string s and an integer k. You can choose any character
# of the string and change it to any other uppercase English character.
# You can perform this operation at most k times.
#
# Return the length of the longest substring containing the same letter
# you can get after performing the above operations.
#
# 📋 Rules:
# 1. You may replace at most k characters in a substring.
# 2. After replacements, the substring should be all the same letter.
# 3. Return the maximum possible length.
#
# 💡 Examples:
# - Input: s="ABAB", k=2     => Output: 4
# - Input: s="AABABBA", k=1  => Output: 4
#
# Constraints:
# - 1 <= len(s) <= 1e5
# - s consists of only uppercase English letters
# - 0 <= k <= len(s)
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def character_replacement1(s: str, k: int) -> int:
    """
    Method: Sliding Window + Frequency Count
    Complexity: Time O(N) | Space O(1)   (since alphabet size is 26)
    """
    freq = [0] * 26
    l = 0
    max_freq = 0  # max frequency of a single char within the current window
    best = 0

    for r, ch in enumerate(s):
        idx = ord(ch) - ord("A")
        freq[idx] += 1
        if freq[idx] > max_freq:
            max_freq = freq[idx]

        # window_size - max_freq = number of replacements needed to unify window
        while (r - l + 1) - max_freq > k:
            left_idx = ord(s[l]) - ord("A")
            freq[left_idx] -= 1
            l += 1

        # now window is valid
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
def character_replacement(s: str, k: int) -> int:
    # TODO: Implement your solution here.
    return 0
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
        "def character_replacement(s: str, k: int) -> int:\n",
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
        ("ABAB", 2, 4),
        ("AABABBA", 1, 4),
        ("AAAA", 0, 4),
        ("ABCDE", 1, 2),
        ("BAAAB", 2, 5),
        ("", 3, 0),  # not in constraints, but helps robustness
    ]

    print(f"\n🧪 Testing your [character_replacement] function...\n")

    header = f"{'Input (s,k)':<30} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for s, k, expected in test_cases:
        try:
            actual = character_replacement(s, k)
        except Exception:
            actual = "Error"

        is_match = actual == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        in_str = f'("{s}", {k})'
        if len(in_str) > 28:
            in_str = in_str[:25] + "..."
        print(f"{in_str:<30} | {str(expected):<10} | {str(actual):<10} | {status_icon}")

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")
# endregion
