import unittest


def solution(l):
    if len(l) <= 2:
        return 0
    curr_idx = len(l)
    res = 0
    while curr_idx >= 2:
        curr_idx -= 1
        dividers = [n for n in l[:curr_idx] if l[curr_idx] % n == 0]
        multiples = [n for n in l[curr_idx+1:] if n % l[curr_idx] == 0]
        res += len(dividers) * len(multiples)
    return res


def naive_solution(l):
    if len(l) == 0:
        return 0
    res = []
    for i in xrange(len(l)-2):
        for j in xrange(i+1, len(l)-1):
            if l[j] % l[i] != 0:
                continue
            for k in xrange(j+1, len(l)):
                if l[k] % l[j] != 0:
                    continue
                res.append([l[i], l[j], l[k]])
    return len(res)


def naive_solution_improved(l):
    if len(l) == 0:
        return 0
    res = 0
    for i in xrange(len(l)-2):
        multiples = naive_solution_improved_find_multiples(i, l)
        for j in xrange(len(multiples)):
            res += len(naive_solution_improved_find_multiples(j, multiples))
    return res


def naive_solution_improved_find_multiples(idx, l):
    return [n for n in l[idx+1:len(l)] if n % l[idx] == 0]


class TestSolution(unittest.TestCase):
    def test1(self):
        self.assertEqual(3, solution([1, 2, 3, 4, 5, 6]))

    def test2(self):
        self.assertEqual(1, solution([1, 1, 1]))

    def test3(self):
        self.assertEqual(0, solution([]))

    def test4(self):
        self.assertEqual(0, solution([1, 2, 5]))

    def test5(self):
        self.assertEqual(0, solution([1, 2]))

    def test6(self):
        self.assertEqual(0, solution([1, 5, 6]))

    def test7(self):
        self.assertEqual(40888, solution(range(1, 2001)))


if __name__ == '__main__':
    unittest.main()
