import os
import sys

# ======================================================================
# 🧠 CHALLENGE: Contains Duplicate (Python Version)
# ======================================================================
# Description:
# Given an integer array `nums`, return `True` if any value appears
# at least twice in the array, and return `False` if every element
# is distinct.
#
# 📋 Rules:
# 1. Return True if duplicates exist.
# 2. Return False if all elements are unique.
#
# 💡 Examples:
# - contains_duplicate([1,2,3,1])               => True
# - contains_duplicate([1,2,3,4])               => False
# - contains_duplicate([1,1,1,3,3,4,3,2,4,2])   => True
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)


def contains_duplicate_1(nums: list[int]) -> bool:
    """
    Method 2: Hash Set
    Best for: Cases where duplicate appears early in the array.
    """
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def contains_duplicate(nums):
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
        "def contains_duplicate(nums):\n",
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
        ([1, 2, 3, 1], True),
        ([1, 2, 3, 4], False),
        ([1, 1, 1, 3, 3, 4, 3, 2, 4, 2], True),
    ]

    print(f"\n🧪 Testing your [contains_duplicate] function...\n")

    header = f"{'Input nums':<30} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, expected in test_cases:
        try:
            result = contains_duplicate(nums)
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
