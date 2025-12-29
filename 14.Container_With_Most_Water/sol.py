import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Container With Most Water (Python Version)
# ======================================================================
# Description:
# Given an integer array `height` of length n, there are n vertical lines
# drawn such that the endpoints of the i-th line are (i, 0) and (i, height[i]).
#
# Find two lines that together with the x-axis form a container, such that the
# container contains the most water.
#
# Return the maximum amount of water a container can store.
#
# Notice:
# - You may not slant the container.
#
# 📋 Rules:
# 1. Choose two indices i < j.
# 2. Area = (j - i) * min(height[i], height[j]).
# 3. Return the maximum possible area.
#
# 💡 Examples:
# - Input: [1,8,6,2,5,4,8,3,7] => Output: 49
# - Input: [1,1]               => Output: 1
#
# Constraints:
# - 2 <= n <= 10^5
# - 0 <= height[i] <= 10^4
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def max_area1(height: List[int]) -> int:
    """
    Method: Two Pointers (Greedy)
    Complexity: Time O(N) | Space O(1)
    """
    l, r = 0, len(height) - 1
    best = 0

    while l < r:
        h = height[l] if height[l] < height[r] else height[r]
        width = r - l
        area = h * width
        if area > best:
            best = area

        # Move the pointer with smaller height
        if height[l] <= height[r]:
            l += 1
        else:
            r -= 1

    return best


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def max_area(height: List[int]) -> int:
    # TODO: Implement your solution here.
    return -1
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
        "def max_area(height: List[int]) -> int:\n",
        "    # TODO: Implement your solution here.\n",
        "    return -1\n",
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
        ([1, 8, 6, 2, 5, 4, 8, 3, 7], 49),
        ([1, 1], 1),
        ([4, 3, 2, 1, 4], 16),
        ([1, 2, 1], 2),
        ([0, 0], 0),
    ]

    print(f"\n🧪 Testing your [max_area] function...\n")

    header = f"{'Input height':<30} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for arr_in, expected in test_cases:
        try:
            actual = max_area(arr_in)
        except Exception:
            actual = "Error"

        is_match = actual == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        in_str = str(arr_in)
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
