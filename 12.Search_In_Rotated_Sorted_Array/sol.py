import os
import sys

# ======================================================================
# 🧠 CHALLENGE: Search in Rotated Sorted Array (Python Version)
# ======================================================================
# Description:
# There is an integer array `nums` sorted in ascending order (with distinct values).
# Prior to being passed to your function, `nums` is possibly rotated at an
# unknown index k.
#
# Given the array `nums` after the possible rotation and an integer `target`,
# return the index of `target` if it is in `nums`, or -1 if it is not.
#
# 📋 Rules:
# 1. You must write an algorithm with O(log n) runtime complexity.
# 2. Return the index of the target, or -1.
#
# 💡 Examples:
# - search([4,5,6,7,0,1,2], 0) => 4
# - search([4,5,6,7,0,1,2], 3) => -1
# - search([1], 0)             => -1
# ======================================================================

# region [📚 Reference Solutions] (Solutions hidden as requested)
# (Focus on implementing your own logic in the Practice Area below!)
# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def search(nums, target):
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
        "def search(nums, target):\n",
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
        ([4, 5, 6, 7, 0, 1, 2], 0, 4),
        ([4, 5, 6, 7, 0, 1, 2], 3, -1),
        ([1], 0, -1),
        ([1], 1, 0),
        ([3, 1], 1, 1),
        ([5, 1, 3], 5, 0),
        ([4, 5, 6, 7, 8, 1, 2, 3], 8, 4),
    ]

    print(f"\n🧪 Testing your [search] function...\n")

    header = f"{'Input nums':<25} | {'Target':<6} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, target, expected in test_cases:
        try:
            result = search(nums, target)
        except Exception as e:
            result = "Error"

        is_match = result == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        nums_str = str(nums)
        if len(nums_str) > 23:
            nums_str = nums_str[:20] + "..."

        print(
            f"{nums_str:<25} | {str(target):<6} | {str(expected):<10} | {str(result):<10} | {status_icon}"
        )

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
