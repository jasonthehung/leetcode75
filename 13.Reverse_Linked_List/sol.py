import os
import sys
from typing import Optional

# ======================================================================
# 🧠 CHALLENGE: Reverse Linked List (Python Version)
# ======================================================================
# Description:
# Given the `head` of a singly linked list, reverse the list, and return
# the reversed list.
#
# 📋 Rules:
# 1. Reverse the direction of the pointers.
# 2. Return the new head of the reversed list.
#
# 💡 Examples:
# - Input: [1,2,3,4,5] => Output: [5,4,3,2,1]
# - Input: [1,2]       => Output: [2,1]
# - Input: []          => Output: []
# ======================================================================

# region [📚 Reference Solutions] (Solutions hidden as requested)


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverse_list1(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Method: Iterative (Three Pointers)
    Complexity: Time O(N) | Space O(1)
    """
    prev = None
    curr = head

    while curr:
        # 1. 暫存下一步 (Save next)
        # 因為等等要切斷 curr.next，不存起來會斷鍊
        next_node = curr.next

        # 2. 反轉箭頭 (Reverse pointer)
        # 將當前節點指回前一個節點
        curr.next = prev

        # 3. 前進 (Move forward)
        # prev 往前一步變成 curr
        prev = curr
        # curr 往前一步變成原本的 next
        curr = next_node

    # 最後 curr 會變成 None，prev 則會停在原本的最後一個節點（新的頭）
    return prev


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def reverse_list(head):
    # TODO: Implement your solution here.
    return None


# <PRACTICE_END>
# endregion

# ======================================================================
#  region [🚀 Test Runner & Auto-Reset] (Do not modify below this line)
# ======================================================================


def to_linked_list(arr):
    if not arr:
        return None
    head = ListNode(arr[0])
    current = head
    for val in arr[1:]:
        current.next = ListNode(val)
        current = current.next
    return head


def from_linked_list(head):
    arr = []
    current = head
    while current:
        arr.append(current.val)
        current = current.next
    return arr


def reset_practice_area():
    print("\n🔄 Resetting Practice Area to default state...")
    MARKER_START = "# <PRACTICE_" + "START>"
    MARKER_END = "# <PRACTICE_" + "END>"

    default_code_lines = [
        "def reverse_list(head):\n",
        "    # TODO: Implement your solution here.\n",
        "    return None\n",
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
        ([1, 2, 3, 4, 5], [5, 4, 3, 2, 1]),
        ([1, 2], [2, 1]),
        ([], []),
    ]

    print(f"\n🧪 Testing your [reverse_list] function...\n")

    header = f"{'Input List':<20} | {'Expected':<20} | {'Actual':<20} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for arr_in, arr_out in test_cases:
        try:
            head = to_linked_list(arr_in)
            result_head = reverse_list(head)
            result_arr = from_linked_list(result_head)
        except Exception as e:
            result_arr = "Error"

        is_match = result_arr == arr_out
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        in_str = str(arr_in)
        if len(in_str) > 18:
            in_str = in_str[:15] + "..."
        out_str = str(arr_out)
        if len(out_str) > 18:
            out_str = out_str[:15] + "..."
        res_str = str(result_arr)
        if len(res_str) > 18:
            res_str = res_str[:15] + "..."

        print(f"{in_str:<20} | {out_str:<20} | {res_str:<20} | {status_icon}")

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
