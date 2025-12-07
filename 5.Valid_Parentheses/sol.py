import os
import sys

# ======================================================================
# 🧠 CHALLENGE: Valid Parentheses (Python Version)
# ======================================================================
# Description:
# Given a string `s` containing just the characters '(', ')', '{', '}', '[' and ']',
# determine if the input string is valid.
#
# An input string is valid if:
# 1. Open brackets must be closed by the same type of brackets.
# 2. Open brackets must be closed in the correct order.
# 3. Every close bracket has a corresponding open bracket of the same type.
#
# 📋 Rules:
# - Return True if the string is valid.
# - Return False otherwise.
#
# 💡 Examples:
# - is_valid("()")       => True
# - is_valid("()[]{}")   => True
# - is_valid("(]")       => False
# - is_valid("([])")     => True
# - is_valid("([)]")     => False
# ======================================================================

# region [📚 Reference Solutions] (Solutions hidden as requested)


def is_valid1(s: str) -> bool:
    """
    Method: Stack with Hash Map
    Complexity: Time O(N) | Space O(N)
    """
    stack = []
    # Map: Closing -> Opening
    # This makes looking up the expected "pair" easier
    pairs = {")": "(", "]": "[", "}": "{"}

    for char in s:
        # If char is a closing bracket (it exists in pairs keys)
        if char in pairs:
            # 1. Check if stack is empty (Invalid case: starts with closing)
            # 2. Pop the top element
            if not stack:
                return False

            top_element = stack.pop()

            # 3. Check if the popped element matches the mapping
            if top_element != pairs[char]:
                return False
        else:
            # If it's an opening bracket, push to stack
            stack.append(char)

    # Finally, stack must be empty (all opened brackets are closed)
    return len(stack) == 0


# endregion


# ======================================================================
#  region [✍️ Practice Area]
#  Please write your solution between the markers below.
# ======================================================================
# <PRACTICE_START>
def is_valid(s):
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
        "def is_valid(s):\n",
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
        ("()", True),
        ("()[]{}", True),
        ("(]", False),
        ("([])", True),
        ("([)]", False),
        ("]", False),
        ("", True),  # Empty string is technically valid (no unmatched brackets)
    ]

    print(f"\n🧪 Testing your [is_valid] function...\n")

    header = f"{'Input s':<20} | {'Expected':<10} | {'Actual':<10} | Status"
    print(header)
    print("-" * len(header))

    all_pass = True
    for s, expected in test_cases:
        try:
            result = is_valid(s)
        except Exception as e:
            result = "Error"

        is_match = result == expected
        status_icon = "✅ PASS" if is_match else "❌ FAIL"
        if not is_match:
            all_pass = False

        # Format output
        s_str = s if len(s) < 20 else s[:17] + "..."

        print(f"{s_str:<20} | {str(expected):<10} | {str(result):<10} | {status_icon}")

    print("-" * len(header))
    if all_pass:
        print("\n🎉 Fantastic! All test cases passed.")
        reset_practice_area()
    else:
        print("\n⚠️  Some tests failed. Keep trying! (The file will not reset yet)")

# endregion
