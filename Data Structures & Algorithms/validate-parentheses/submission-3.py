class Solution:
    def isValid(self, s: str) -> bool:
        close_to_open = {
            ")": "(",
            "]": "[",
            "}": "{"}
        stack = []
        for b in s:
            if b not in close_to_open:
                stack.append(b)
                continue
            if not len(stack):
                return False
            if stack.pop() != close_to_open[b]:
                return False
        return len(stack) == 0