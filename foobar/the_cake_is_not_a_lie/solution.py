import unittest


def solution(s):
    if not is_valid(s):
        return 0
    # if len(s) % 2 != 0: TODO
    return divide_rec(s)


def is_valid(s):
    if s == "":
        return False
    min_c, max_c = 'a', 'z'
    for c in s:
        if min_c <= c <= max_c:
            continue
        return False
    return True


def divide_rec(s, pieces=1):
    str_len = len(s)
    if str_len == 1:
        return pieces
    middle = str_len/2
    left, right = s[0:middle], s[middle:str_len]
    if left == right:
        pieces *= 2
    return divide_rec(left, pieces)


class TestSolution(unittest.TestCase):
    def test1(self):
        self.assertEqual(solution(""), 0)

    def test2(self):
        self.assertEqual(2, solution("abccbaabccba"))

    def test3(self):
        self.assertEqual(4, solution("abcabcabcabc"))

    def test4(self):
        self.assertEqual(8, solution("aaaaaaaa"))

    def test5(self):
        self.assertEqual(1, solution("abcdefg"))

    def test6(self):
        self.assertEqual(0, solution("abccb8aabccba"))

    def test7(self):
        self.assertEqual(5, solution("abcabcabcabcabc"))

    def test8(self):
        self.assertEqual(3, solution("abcabcabc"))

    def test9(self):
        self.assertEqual(0, solution("ababT"))


if __name__ == "__main__":
    unittest.main()
