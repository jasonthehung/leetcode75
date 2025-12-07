import os
import sys
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Best Time to Buy and Sell Stock (Python Version)
# ======================================================================
# Description:
# You are given an array `prices` where `prices[i]` is the price of a
# given stock on the ith day.
#
# You want to maximize your profit by choosing a single day to buy one
# stock and choosing a different day in the future to sell that stock.
#
# 📋 Rules:
# 1. Return the maximum profit you can achieve.
# 2. If you cannot achieve any profit, return 0.
#
# 💡 Examples:
# - max_profit([7,1,5,3,6,4]) => 5 (Buy at 1, sell at 6)
# - max_profit([7,6,4,3,1])   => 0 (No profit possible)
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def max_profit1(prices: List[int]) -> int:
    """
    Method: One Pass (Greedy)
    Complexity: Time O(N) | Space O(1)
    """
    if not prices:
        return 0

    # Initialize min_price to infinity so the first price becomes the minimum
    min_price = float("inf")
    max_profit = 0

    for price in prices:
        # Case 1: Found a lower buy price? Update min_price.
        if price < min_price:
            min_price = price
        # Case 2: Found a higher sell price? Check if profit is greater.
        # Note: We use elif because we can't sell on the same day we bought (profit 0)
        elif (price - min_price) > max_profit:
            max_profit = price - min_price

    return max_profit  # type: ignore


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def max_profit(prices):
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
        "def max_profit(prices):\n",
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
        ([7, 1, 5, 3, 6, 4], 5),
        ([7, 6, 4, 3, 1], 0),
        ([1, 2], 1),
        ([2, 4, 1], 2),
    ]

    print(f"\n🧪 Testing your [max_profit] function...\n")

    header = f"{'Input prices':<20} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for prices, expected in test_cases:
        try:
            result = max_profit(prices)
        except Exception as e:
            result = "Error"

        is_match = result == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        prices_str = str(prices)
        if len(prices_str) > 18:
            prices_str = prices_str[:15] + "..."

        print(
            f"{prices_str:<20} | {str(expected):<10} | {str(result):<10} | {status_icon}"
        )

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
