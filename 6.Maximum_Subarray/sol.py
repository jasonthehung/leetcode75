import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Maximum Subarray (Python Version)
# ======================================================================
# Description:
# Given an integer array `nums`, find the subarray with the largest sum,
# and return its sum.
#
# 📋 Rules:
# 1. A subarray is a contiguous part of an array.
# 2. Return the maximum sum found.
# 3. Handle cases with negative numbers correctly.
#
# 💡 Examples:
# - max_sub_array([-2,1,-3,4,-1,2,1,-5,4]) => 6  ([4,-1,2,1])
# - max_sub_array([1])                     => 1
# - max_sub_array([5,4,-1,7,8])            => 23
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def max_sub_array1(nums: List[int]) -> int:
    """
    Method: Kadane's Algorithm
    Complexity: Time O(N) | Space O(1)
    """
    # Initialize with the first element (handles array of size 1)
    current_sum = nums[0]
    max_sum = nums[0]

    # Iterate starting from the second element
    for num in nums[1:]:
        # Decision: Start new subarray at num, or extend previous sum?
        # If current_sum is negative, adding it decreases value, so we restart at num.
        current_sum = max(num, current_sum + num)

        # Update global max
        max_sum = max(max_sum, current_sum)

    return max_sum


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def max_sub_array(nums):
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
        "def max_sub_array(nums):\n",
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
        ([-2, 1, -3, 4, -1, 2, 1, -5, 4], 6),
        ([1], 1),
        ([5, 4, -1, 7, 8], 23),
        ([-1], -1),  # Edge case: single negative number
        ([-2, -1], -1),  # Edge case: all negatives
    ]

    print(f"\n🧪 Testing your [max_sub_array] function...\n")

    header = f"{'Input nums':<30} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, expected in test_cases:
        try:
            result = max_sub_array(nums)
        except Exception as e:
            result = "Error"

        is_match = result == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        nums_str = str(nums)
        if len(nums_str) > 28:
            nums_str = nums_str[:25] + "..."

        print(
            f"{nums_str:<30} | {str(expected):<10} | {str(result):<10} | {status_icon}"
        )

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
