import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: 3Sum (Python Version)
# ======================================================================
# Description:
# Given an integer array `nums`, return all the triplets [nums[i], nums[j], nums[k]]
# such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
#
# Notice that the solution set must not contain duplicate triplets.
#
# 📋 Rules:
# 1. Return a list of lists, where each inner list is a triplet.
# 2. The order of the output and the order of the triplets does not matter.
# 3. Elements in a triplet must sum to 0.
#
# 💡 Examples:
# - three_sum([-1,0,1,2,-1,-4]) => [[-1,-1,2], [-1,0,1]]
# - three_sum([0,1,1])          => []
# - three_sum([0,0,0])          => [[0,0,0]]
# ======================================================================

# region [📚 Reference Solutions] (Solutions hidden as requested)


def three_sum1(nums: List[int]) -> List[List[int]]:
    """
    Method: Sorting + Two Pointersg
    Complexity: Time O(N^2) | Space O(1) (depending on sort implementation)
    """
    nums.sort()
    result = []
    n = len(nums)

    for i in range(n):
        # Optimization: If the smallest number is > 0, we can never sum to 0.
        if nums[i] > 0:
            break

        # 1. Skip duplicate 'i' to avoid duplicate triplets
        if i > 0 and nums[i] == nums[i - 1]:
            continue

        left, right = i + 1, n - 1

        while left < right:
            total = nums[i] + nums[left] + nums[right]

            if total < 0:
                # Sum is too small, need a larger number
                left += 1
            elif total > 0:
                # Sum is too big, need a smaller number
                right -= 1
            else:
                # Found a triplet!
                result.append([nums[i], nums[left], nums[right]])
                left += 1
                right -= 1

                # 2. Skip duplicate 'left' values
                while left < right and nums[left] == nums[left - 1]:
                    left += 1

                # 3. Skip duplicate 'right' values (Optional optimization)
                # Note: The logic works without this loop because the outer loop handles it,
                # but adding it speeds up cases with many duplicates.
                while left < right and nums[right] == nums[right + 1]:
                    right -= 1

    return result


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def three_sum(nums):
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
        "def three_sum(nums):\n",
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
        ([-1, 0, 1, 2, -1, -4], [[-1, -1, 2], [-1, 0, 1]]),
        ([0, 1, 1], []),
        ([0, 0, 0], [[0, 0, 0]]),
    ]

    print(f"\n🧪 Testing your [three_sum] function...\n")

    header = f"{'Input nums':<25} | {'Expected':<20} | {'Actual':<20} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, expected in test_cases:
        try:
            result = three_sum(nums)
        except Exception as e:
            result = "Error"

        # Canonicalize results for comparison (Sort inner triplets, then sort the list of triplets)
        is_match = False
        try:
            if isinstance(result, list):
                res_sorted = sorted([sorted(x) for x in result])
                exp_sorted = sorted([sorted(x) for x in expected])
                is_match = res_sorted == exp_sorted
        except:
            is_match = False

        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        nums_str = str(nums)
        if len(nums_str) > 23:
            nums_str = nums_str[:20] + "..."
        exp_str = str(expected)
        if len(exp_str) > 18:
            exp_str = exp_str[:15] + "..."
        res_str = str(result)
        if len(res_str) > 18:
            res_str = res_str[:15] + "..."

        print(f"{nums_str:<25} | {exp_str:<20} | {res_str:<20} | {status_icon}")

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
