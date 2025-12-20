import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Merge Intervals (Python Version)
# ======================================================================
# Description:
# Given an array of `intervals` where intervals[i] = [start_i, end_i],
# merge all overlapping intervals, and return an array of the
# non-overlapping intervals that cover all the intervals in the input.
#
# 📋 Rules:
# 1. Intervals [a, b] and [c, d] overlap if a <= d and c <= b.
# 2. Intervals touching at boundaries (e.g., [1,4] and [4,5]) are considered overlapping.
# 3. The input intervals might not be sorted.
#
# 💡 Examples:
# - merge([[1,3],[2,6],[8,10],[15,18]]) => [[1,6],[8,10],[15,18]]
# - merge([[1,4],[4,5]])                => [[1,5]]
# - merge([[4,7],[1,4]])                => [[1,7]]
# ======================================================================

# region [📚 Reference Solutions] (Solutions hidden as requested)


def merge1(intervals: List[List[int]]) -> List[List[int]]:
    # 1. Edge Case: Empty list
    if not intervals:
        return []

    # 2. Sort by start time
    intervals.sort(key=lambda x: x[0])

    result = [intervals[0]]

    # Skip the first one since it's already in result
    for interval in intervals[1:]:
        last_end = result[-1][1]
        current_start = interval[0]
        current_end = interval[1]

        # 3. Check for overlap
        # 正確邏輯：如果當前的「開始」比上一段的「結束」還晚，代表沒有重疊
        if current_start > last_end:
            result.append(interval)
        else:
            # 重疊了！合併方式是更新結尾
            # 取兩者結尾的最大值
            result[-1][1] = max(last_end, current_end)

    return result


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def merge(intervals):
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
        "def merge(intervals):\n",
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
        ([[1, 3], [2, 6], [8, 10], [15, 18]], [[1, 6], [8, 10], [15, 18]]),
        ([[1, 4], [4, 5]], [[1, 5]]),
        ([[4, 7], [1, 4]], [[1, 7]]),  # Unsorted input
        ([[1, 4], [0, 4]], [[0, 4]]),
        ([[1, 4], [2, 3]], [[1, 4]]),  # One interval fully inside another
    ]

    print(f"\n🧪 Testing your [merge] function...\n")

    header = f"{'Input Intervals':<30} | {'Expected':<20} | {'Actual':<20} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for intervals, expected in test_cases:
        try:
            # Pass a copy to avoid side-effects if user sorts in-place
            result = merge([x[:] for x in intervals])
        except Exception as e:
            result = "Error"

        # Compare results (Sorting usually required for deterministic comparison)
        is_match = False
        try:
            if isinstance(result, list):
                # We expect the result to be sorted by start time naturally
                is_match = result == expected
        except:
            is_match = False

        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        inp_str = str(intervals)
        if len(inp_str) > 28:
            inp_str = inp_str[:25] + "..."
        exp_str = str(expected)
        if len(exp_str) > 18:
            exp_str = exp_str[:15] + "..."
        res_str = str(result)
        if len(res_str) > 18:
            res_str = res_str[:15] + "..."

        print(f"{inp_str:<30} | {exp_str:<20} | {res_str:<20} | {status_icon}")

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
