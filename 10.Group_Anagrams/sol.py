import os
import sys
from collections import defaultdict
from typing import List

# ======================================================================
# 🧠 CHALLENGE: Group Anagrams (Python Version)
# ======================================================================
# Description:
# Given an array of strings `strs`, group the anagrams together.
# You can return the answer in any order.
#
# An Anagram is a word or phrase formed by rearranging the letters of a
# different word or phrase, typically using all the original letters exactly once.
#
# 📋 Rules:
# 1. Return a list of lists, where each inner list contains words that are anagrams.
# 2. The order of the groups and the order of words within groups does not matter.
#
# 💡 Examples:
# - group_anagrams(["eat","tea","tan","ate","nat","bat"])
#   => [["bat"],["nat","tan"],["ate","eat","tea"]]
# - group_anagrams([""])  => [[""]]
# - group_anagrams(["a"]) => [["a"]]
# ======================================================================


# region [📚 Reference Solutions] (Solutions hidden as requested)
def group_anagrams1(strs: List[str]) -> List[List[str]]:
    """
    Method: Sorting + Hash Map
    Complexity: Time O(N * K log K) | Space O(N * K)
    (N = number of strings, K = max length of a string)
    """
    # Key: Sorted String (e.g., "aet")
    # Value: List of original strings (e.g., ["eat", "tea", "ate"])
    anagram_map = defaultdict(list)

    for word in strs:
        # 1. Generate the key by sorting characters
        # sorted("eat") becomes ['a', 'e', 't']
        # "".join(...) turns it back into "aet"
        sorted_key = "".join(sorted(word))

        # 2. Append the original word to the correct group
        anagram_map[sorted_key].append(word)

    # 3. Return just the groups (values)
    return list(anagram_map.values())


def group_anagrams2(strs: list[str]) -> list[list[str]]:
    # 建議改名為 groups 或 anagram_map
    groups = defaultdict(list)

    for word in strs:
        # 1. 建立計數陣列 (Frequency Array)
        # 這裡用 List 是為了可以修改內容
        count = [0] * 26

        for char in word:
            count[ord(char) - ord("a")] += 1

        # 2. 關鍵：轉成 Tuple 才能當 Dictionary Key
        # key 像是 (1, 0, 0, 0, 1, ..., 0)
        key = tuple(count)

        groups[key].append(word)

    return list(groups.values())


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def group_anagrams(strs):
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
        "def group_anagrams(strs):\n",
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
        (
            ["eat", "tea", "tan", "ate", "nat", "bat"],
            [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
        ),
        ([""], [[""]]),
        (["a"], [["a"]]),
        (
            ["abc", "bca", "cba", "xyz", "zyx", "klm"],
            [["abc", "bca", "cba"], ["xyz", "zyx"], ["klm"]],
        ),
    ]

    print(f"\n🧪 Testing your [group_anagrams] function...\n")

    header = f"{'Input strs':<30} | {'Status':<10}"
    print(header)
    print("-" * 45)

    all_pass = True
    for strs, expected in test_cases:
        try:
            result = group_anagrams(strs)
        except Exception as e:
            result = "Error"

        # Canonicalize results for comparison (Sort inner lists, then sort the outer list)
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
        strs_str = str(strs)
        if len(strs_str) > 28:
            strs_str = strs_str[:25] + "..."

        print(f"{strs_str:<30} | {status_icon}")
        if not is_match:
            print(f"   Expected: {expected}")
            print(f"   Actual:   {result}")

    print("-" * 45)
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
