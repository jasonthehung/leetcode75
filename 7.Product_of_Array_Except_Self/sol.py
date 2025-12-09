import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Product of Array Except Self (Python Version)
# ======================================================================
# Description:
# Given an integer array `nums`, return an array `answer` such that
# `answer[i]` is equal to the product of all the elements of `nums`
# except `nums[i]`.
#
# 📋 Rules:
# 1. You must write an algorithm that runs in O(n) time.
# 2. You CANNOT use the division operation.
# 3. Follow up: Can you solve it with O(1) extra space complexity?
#    (The output array does not count as extra space).
#
# 💡 Examples:
# - product_except_self([1,2,3,4])       => [24,12,8,6]
# - product_except_self([-1,1,0,-3,3])   => [0,0,9,0,0]
# ======================================================================

# region [📚 Reference Solutions] (Solutions hidden as requested)


def product_except_self1(nums: List[int]) -> List[int]:
    n = len(nums)
    answer = [1] * n

    # 1. Left Pass (你是對的)
    for i in range(1, n):
        answer[i] = answer[i - 1] * nums[i - 1]

    # 2. Right Pass (修正部分)
    right_product = 1

    # 從最後一個元素開始 (n-1)，一路走到 0
    for i in range(n - 1, -1, -1):
        # 關鍵：是「乘上去」(*=)，不是「覆蓋」(=)
        # answer[i] 目前是左邊積，right_product 是右邊積
        answer[i] *= right_product

        # 更新累積的右邊積
        right_product *= nums[i]

    return answer


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def product_except_self(nums):
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
        "def product_except_self(nums):\n",
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
        ([1, 2, 3, 4], [24, 12, 8, 6]),
        ([-1, 1, 0, -3, 3], [0, 0, 9, 0, 0]),
        ([0, 0], [0, 0]),  # Double zero case
        ([1, 0], [0, 1]),  # Single zero case
    ]

    print(f"\n🧪 Testing your [product_except_self] function...\n")

    header = f"{'Input nums':<25} | {'Expected':<20} | {'Actual':<20} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for nums, expected in test_cases:
        try:
            result = product_except_self(nums)
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
