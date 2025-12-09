import os
import sys
from collections import defaultdict

# ======================================================================
# 🧠 CHALLENGE: Valid Anagram (Python Version)
# ======================================================================
# Description:
# Given two strings `s` and `t`, return `True` if `t` is an anagram of `s`,
# and `False` otherwise.
#
# An Anagram is a word or phrase formed by rearranging the letters of a
# different word or phrase, typically using all the original letters exactly once.
#
# 📋 Rules:
# 1. Return True if t is an anagram of s.
# 2. Return False otherwise.
#
# 💡 Examples:
# - is_anagram("anagram", "nagaram") => True
# - is_anagram("rat", "car")         => False
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)


def is_anagram1(s: str, t: str) -> bool:
    """
    Method: Hash Map using defaultdict
    Best for: Cleaner syntax than standard dict (no need for .get checks).
    """
    if len(s) != len(t):
        return False

    # int 的預設值是 0
    count = defaultdict(int)

    # 1. 記錄 s 的字母出現次數
    for char in s:
        count[char] += 1

    # 2. 扣除 t 的字母出現次數
    for char in t:
        count[char] -= 1

        # 優化：如果減到 < 0，代表 t 有額外的字母，直接回傳 False
        if count[char] < 0:
            return False

    return True


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def is_anagram(s, t):
    # TODO: Implement your solution here.
    return False
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
        "def is_anagram(s, t):\n",
        "    # TODO: Implement your solution here.\n",
        "    return False\n",
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
        ("anagram", "nagaram", True),
        ("rat", "car", False),
        ("a", "ab", False),
        ("ab", "a", False),
        ("", "", True),
    ]

    print(f"\n🧪 Testing your [is_anagram] function...\n")

    header = f"{'s':<10} | {'t':<10} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for s, t, expected in test_cases:
        try:
            result = is_anagram(s, t)
        except Exception as e:
            result = "Error"

        is_match = result == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        s_str = s if len(s) < 10 else s[:7] + "..."
        t_str = t if len(t) < 10 else t[:7] + "..."

        print(
            f"{s_str:<10} | {t_str:<10} | {str(expected):<10} | {str(result):<10} | {status_icon}"
        )

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
