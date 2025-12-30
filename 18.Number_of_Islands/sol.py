import os
import sys
from collections import deque
from typing import Deque, List, Tuple

# ======================================================================
# 🧠 CHALLENGE: Number of Islands (Python Version)
# ======================================================================
# Description:
# Given an m x n 2D binary grid `grid` which represents a map of '1's (land)
# and '0's (water), return the number of islands.
#
# An island is surrounded by water and is formed by connecting adjacent lands
# horizontally or vertically. You may assume all four edges are surrounded by water.
#
# 📋 Rules:
# 1. Cells are connected 4-directionally (up/down/left/right).
# 2. Count connected components of '1'.
#
# 💡 Examples:
# - grid = [
#   ["1","1","1","1","0"],
#   ["1","1","0","1","0"],
#   ["1","1","0","0","0"],
#   ["0","0","0","0","0"]
# ] => 1
#
# - grid = [
#   ["1","1","0","0","0"],
#   ["1","1","0","0","0"],
#   ["0","0","1","0","0"],
#   ["0","0","0","1","1"]
# ] => 3
#
# Constraints:
# - 1 <= m, n <= 300
# - grid[i][j] is '0' or '1'
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def num_islands1(grid: List[List[str]]) -> int:
    """
    Method: BFS Flood Fill (Mutate grid to mark visited)
    Complexity: Time O(M*N) | Space O(min(M*N)) (queue worst-case)
    """
    if not grid or not grid[0]:
        return 0

    m, n = len(grid), len(grid[0])
    islands = 0
    q: Deque[Tuple[int, int]] = deque()

    for i in range(m):
        for j in range(n):
            if grid[i][j] != "1":
                continue

            islands += 1
            grid[i][j] = "0"  # mark visited by sinking land
            q.append((i, j))

            while q:
                x, y = q.popleft()
                if x > 0 and grid[x - 1][y] == "1":
                    grid[x - 1][y] = "0"
                    q.append((x - 1, y))
                if x + 1 < m and grid[x + 1][y] == "1":
                    grid[x + 1][y] = "0"
                    q.append((x + 1, y))
                if y > 0 and grid[x][y - 1] == "1":
                    grid[x][y - 1] = "0"
                    q.append((x, y - 1))
                if y + 1 < n and grid[x][y + 1] == "1":
                    grid[x][y + 1] = "0"
                    q.append((x, y + 1))

    return islands


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def num_islands(grid: List[List[str]]) -> int:
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
        "def num_islands(grid: List[List[str]]) -> int:\n",
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


def deep_copy_grid(g: List[List[str]]) -> List[List[str]]:
    return [row[:] for row in g]


if __name__ == "__main__":
    test_cases = [
        (
            [
                ["1", "1", "1", "1", "0"],
                ["1", "1", "0", "1", "0"],
                ["1", "1", "0", "0", "0"],
                ["0", "0", "0", "0", "0"],
            ],
            1,
        ),
        (
            [
                ["1", "1", "0", "0", "0"],
                ["1", "1", "0", "0", "0"],
                ["0", "0", "1", "0", "0"],
                ["0", "0", "0", "1", "1"],
            ],
            3,
        ),
        ([["0"]], 0),
        ([["1"]], 1),
        (
            [
                ["1", "0", "1"],
                ["0", "1", "0"],
                ["1", "0", "1"],
            ],
            5,
        ),
    ]

    print(f"\n🧪 Testing your [num_islands] function...\n")

    header = f"{'Grid size':<12} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for grid, expected in test_cases:
        try:
            # Make a copy because most solutions will mutate the grid
            actual = num_islands(deep_copy_grid(grid))
        except Exception:
            actual = "Error"

        is_match = actual == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        size_str = f"{len(grid)}x{len(grid[0])}"
        print(
            f"{size_str:<12} | {str(expected):<10} | {str(actual):<10} | {status_icon}"
        )

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")
# endregion
