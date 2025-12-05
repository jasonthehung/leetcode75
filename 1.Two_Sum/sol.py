import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Two Sum (Python Version)
# ======================================================================
# Description:
# Given an array of integers `nums` and an integer `target`, return
# indices of the two numbers such that they add up to `target`.
#
# 📋 Rules:
# 1. You may assume that each input would have EXACTLY one solution.
# 2. You may not use the same element twice.
# 3. You can return the answer in any order.
# 4. Follow-up: Can you come up with an algorithm that is less than
#    O(n^2) time complexity?
#
# 💡 Examples:
# - two_sum([2,7,11,15], 9) => [0, 1]
# - two_sum([3,2,4], 6)     => [1, 2]
# - two_sum([3,3], 6)       => [0, 1]
# ======================================================================

# #region [📚 Reference Solutions]


def two_sum1(nums: List[int], target: int) -> List[int]:
    """
    Method: One-pass Hash Table
    Complexity: Time O(N) | Space O(N)
    """
    # Dictionary to store number: index
    seen = {}

    for i, num in enumerate(nums):
        complement = target - num

        # Check if the complement exists in our dictionary
        if complement in seen:
            return [seen[complement], i]

        # Store current number and index
        seen[num] = i

    return []  # Should not happen per problem description


# #endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def two_sum(nums, target):
    # TODO: Implement your solution here.
    return []


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
        "def two_sum(nums, target):\n",
        "    # TODO: Implement your solution here.\n",
        "    return []\n",
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
        ([2, 7, 11, 15], 9, [0, 1]),
        ([3, 2, 4], 6, [1, 2]),
        ([3, 3], 6, [0, 1]),
    ]

    print(f"\n🧪 Testing your [two_sum] function...\n")

    header = f"{'Input nums':<20} | {'Target':<6} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, target, expected in test_cases:
        try:
            result = two_sum(nums, target)
        except Exception as e:
            result = "Error"

        # Compare sorted lists to allow returning indices in any order
        is_match = False
        if isinstance(result, list) and len(result) == 2:
            is_match = sorted(result) == sorted(expected)

        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        nums_str = str(nums)
        if len(nums_str) > 18:
            nums_str = nums_str[:15] + "..."

        print(
            f"{nums_str:<20} | {str(target):<6} | {str(expected):<10} | {str(result):<10} | {status_icon}"
        )

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
