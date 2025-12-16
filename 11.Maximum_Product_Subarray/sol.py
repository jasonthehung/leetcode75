import os
import sys

# ======================================================================
# 🧠 CHALLENGE: Maximum Product Subarray (Python Version)
# ======================================================================
# Description:
# Given an integer array `nums`, find a subarray that has the largest
# product, and return the product.
#
# The test cases are generated so that the answer will fit in a 32-bit integer.
#
# 📋 Rules:
# 1. A subarray is a contiguous part of an array.
# 2. Return the maximum product found.
# 3. Handle negative numbers (negative * negative = positive) and zeros.
#
# 💡 Examples:
# - max_product([2,3,-2,4]) => 6  ([2,3])
# - max_product([-2,0,-1])  => 0  (The result cannot be 2, because [-2,-1] is not a subarray)
# - max_product([-2,3,-4])  => 24 ([-2,3,-4])
# ======================================================================

# region [📚 Reference Solutions] (Solutions hidden as requested)
# (Focus on implementing your own logic in the Practice Area below!)
# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def max_product(nums):
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
        "def max_product(nums):\n",
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
        ([2, 3, -2, 4], 6),
        ([-2, 0, -1], 0),
        ([-2, 3, -4], 24),
        ([-4, -3, -2], 12),
        ([0, 2], 2),
        ([-2], -2),
        ([2, -5, -2, -4, 3], 24),
    ]

    print(f"\n🧪 Testing your [max_product] function...\n")

    header = f"{'Input nums':<30} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, expected in test_cases:
        try:
            result = max_product(nums)
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
