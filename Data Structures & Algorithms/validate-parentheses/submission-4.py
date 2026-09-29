class Solution:
	def isValid(self, s: str) -> bool:
		stack = []
		for b in s:
			if b in '({[':
				stack.append(b)
				continue
			if (b == '}' and (not stack or stack.pop() != '{')) or \
			(b == ']' and (not stack or stack.pop() != '[')) or \
			(b == ')' and (not stack or stack.pop() != '(')):
				return False
		return len(stack) == 0