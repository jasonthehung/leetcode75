import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Find Minimum in Rotated Sorted Array (Python Version)
# ======================================================================
# Description:
# Suppose an array of length n sorted in ascending order is rotated between
# 1 and n times. Given the sorted rotated array `nums` of unique elements,
# return the minimum element of this array.
#
# You must write an algorithm that runs in O(log n) time.
#
# 📋 Rules:
# 1. nums contains unique integers.
# 2. nums is sorted ascending then rotated.
# 3. Return the minimum value in nums.
#
# 💡 Examples:
# - Input: [3,4,5,1,2]       => Output: 1
# - Input: [4,5,6,7,0,1,2]   => Output: 0
# - Input: [11,13,15,17]     => Output: 11
#
# Constraints:
# - 1 <= n <= 5000
# - -5000 <= nums[i] <= 5000
# - All integers are unique.
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def find_min1(nums: List[int]) -> int:
    """
    Method: Binary Search (Compare with right boundary)
    Complexity: Time O(log N) | Space O(1)
    """
    l, r = 0, len(nums) - 1

    # Invariant: minimum is in [l, r]
    while l < r:
        mid = (l + r) // 2

        # If mid is greater than right, min is to the right of mid
        if nums[mid] > nums[r]:
            l = mid + 1
        else:
            # mid could be the minimum
            r = mid

    return nums[l]


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def find_min(nums: List[int]) -> int:
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
        "def find_min(nums: List[int]) -> int:\n",
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
        ([3, 4, 5, 1, 2], 1),
        ([4, 5, 6, 7, 0, 1, 2], 0),
        ([11, 13, 15, 17], 11),
        ([2, 1], 1),
        ([1], 1),
        ([5, 6, 7, 8, 1, 2, 3, 4], 1),
    ]

    print(f"\n🧪 Testing your [find_min] function...\n")

    header = f"{'Input nums':<35} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, expected in test_cases:
        try:
            actual = find_min(nums)
        except Exception:
            actual = "Error"

        is_match = actual == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        in_str = str(nums)
        if len(in_str) > 33:
            in_str = in_str[:30] + "..."
        print(f"{in_str:<35} | {str(expected):<10} | {str(actual):<10} | {status_icon}")

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")
# endregion
