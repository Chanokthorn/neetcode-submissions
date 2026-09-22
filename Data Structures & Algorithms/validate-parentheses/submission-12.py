class Solution:
    def isValid(self, text: str) -> bool:
        stack = []
        for s in text:
            if s == "(":
                stack.append("(")
                continue
            if s == "[":
                stack.append("[")
                continue
            if s == "{":
                stack.append("{")
                continue
            if len(stack) == 0:
                return False
            top = stack.pop()
            if s == ")":
                if top != "(":
                    return False
            if s == "]":
                if top != "[":
                    return False
            if s == "}":
                if top != "{":
                    return False
        return len(stack) == 0

        