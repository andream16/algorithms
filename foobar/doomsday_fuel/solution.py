from __future__ import division
import unittest
from fractions import Fraction, gcd
from functools import reduce


def normalise_matrix(m):

    num_absorbing_states = 0

    # reorder rows - absorbing states first, then the rest.
    # transform entries in fractions to represent probability.
    current_absorbent = 0
    for i in range(len(m)):
        s = sum(m[i])
        if s == 0:
            m[i][i] = 1
            m = stop_n_swap(m, i, current_absorbent)
            current_absorbent += 1
            num_absorbing_states += 1
        else:
            for j in range(len(m[i])):
                m[i][j] = Fraction(m[i][j]/s).limit_denominator(1000)

    return m, num_absorbing_states


def stop_n_swap(m, x, y):
    m2 = m[:y] + [m[x]] + [m[y]] + m[y+1:x] + m[x+1:]

    for i in range(len(m)):
        m2[i] = m2[i][:y] + [m2[i][x]] + [m2[i][y]] + m2[i][y+1:x] + m2[i][x+1:]

    return m2


def is_absorbing_state(nums):
    for n in nums:
        if n != 0:
            return False
    return True


def get_fraction(n):
    return Fraction(n).limit_denominator(1000)


def sub_matrix(m, r1, c1, r2, c2):
    sub_m = []
    for i in range(r1, r2):
        curr = []
        for j in range(c1, c2):
            curr.append(m[i][j])
        if len(curr) > 0:
            sub_m.append(curr)
    return sub_m


def sub(i, q):
    res = []
    for k in range(0, len(i)):
        curr = []
        for j in range(0, len(i[0])):
            curr.append(i[k][j]-q[k][j])
        res.append(curr)
    return res


def get_f(m):
    determinant = get_determinant(m)
    if determinant == 0:
        determinant = 1
    if len(m) == 2:
        return [[m[1][1]/determinant, -1*m[0][1]/determinant],
                [-1*m[1][0]/determinant, m[0][0]/determinant]]

    cofactors = []
    for r in range(len(m)):
        cofactor_row = []
        for c in range(len(m)):
            minor = get_minor(m, r, c)
            cofactor_row.append(((-1)**(r+c)) * get_determinant(minor))
        cofactors.append(cofactor_row)
    cofactors = transpose(cofactors)
    for r in range(len(cofactors)):
        for c in range(len(cofactors)):
            cofactors[r][c] = cofactors[r][c]/determinant
    return cofactors


def transpose(m):
    return map(list, zip(*m))


def get_minor(m, i, j):
    return [row[:j] + row[j+1:] for row in (m[:i]+m[i+1:])]


def get_determinant(m):
    if len(m) == 2:
        return m[0][0]*m[1][1]-m[0][1]*m[1][0]

    determinant = 0
    for c in range(len(m)):
        determinant += ((-1)**c)*m[0][c]*get_determinant(get_minor(m, 0, c))
    return determinant


def product(m1, m2):
    return [[sum(x * y for x, y in zip(m1_r, m2_c)) for m2_c in zip(*m2)] for m1_r in m1]


def get_identity(rows, cols):
    m = []
    for i in range(rows):
        row = []
        for j in range(cols):
            if i == j:
                row.append(1)
                continue
            row.append(0)
        m.append(row)
    return m


def solution(m):

    m, num_absorbing_states = normalise_matrix(m)

    if num_absorbing_states == 1:
        return [1, 1]

    n_rows = len(m)

    # F = (I-Q)^-1
    q = sub_matrix(m, num_absorbing_states, num_absorbing_states, n_rows, n_rows)
    i = get_identity(len(q), len(q[0]))

    # I-Q
    diff = sub(i, q)
    # (diff)^-1
    f = get_f(diff)
    r = sub_matrix(m, num_absorbing_states, 0, n_rows, num_absorbing_states)
    fr = product(f, r)

    if len(fr) == 0:
        return []

    denominators, denominator = [], 1
    for n in fr[0]:
        if n != 0:
            denominators.append(n.denominator)

    # Least Common Multiple
    if len(denominators) > 0:
        denominator = reduce(lambda a, b: a*b // gcd(a, b), denominators)

    res = []
    for n in fr[0]:
        if n == 0:
            res.append(0)
        else:
            if n.denominator == 0:
                n.denominator = 1
            res.append(int(n.numerator * (denominator / n.denominator)))

    res.append(denominator)
    return res


class TestGetF(unittest.TestCase):
    def test1(self):
        m = [[3, 0, 2], [2, 0, -2], [0, 1, 1]]
        inverse = [[0.2, 0.2, 0], [-0.2, 0.3, 1], [0.2, -0.3, 0]]
        self.assertEqual(get_f(m), inverse)


class TestGetDeterminant(unittest.TestCase):
    def test1(self):
        m = [[3, 8], [4, 6]]
        self.assertEqual(get_determinant(m), -14)

    def test2(self):
        m = [[3, 1, 7], [5, 4, 1], [1, 2, 9]]
        self.assertEqual(get_determinant(m), 100)

    def test3(self):
        m = [[1, 5, 0, 2], [3, 8, 4, 3], [1, 5, 0, 2], [2, 3, 0, 8]]
        self.assertEqual(get_determinant(m), 0)

    def test4(self):
        m = [[4, 0, 1, 2], [3, 1, -1, 3], [2, 0, 3, 1], [2, -2, 3, 1]]
        self.assertEqual(get_determinant(m), 30)


class TestSub(unittest.TestCase):
    def test1(self):
        a, b = [[5, 4, 3], [7, 8, 9], [3, 12, -5]], [[2, 3, 4], [4, -5, 6], [0, 7, 8]]
        self.assertEqual(sub(a, b), [[3, 1, -1], [3, 13, 3], [3, 5, -13]])

    def test2(self):
        a, b = [[1, 0, 0], [0, 1, 0], [0, 0, 1]], [[2, 3, 4], [4, -5, 6], [0, 7, 8]]
        self.assertEqual(sub(a, b), [[-1, -3, -4], [-4, 6, -6], [0, -7, -7]])


class TestProduct(unittest.TestCase):
    def test1(self):
        a, b = [[-1, 2, 3], [4, 0, 5]], [[5, -1], [-4, 0], [2, 3]]
        self.assertEqual(product(a, b), [[-7, 10], [30, 11]])

    def test2(self):
        a, b = [[3, 2, 1, 5], [9, 1, 3, 0]], [[2, 9, 0], [1, 3, 5], [2, 4, 7], [8, 1, 5]]
        self.assertEqual(product(a, b), [[50, 42, 42], [25, 96, 26]])


class TestSolution(unittest.TestCase):
    def test1(self):
        test_input = [
            [0, 1, 0, 0, 0, 1],
            [4, 0, 0, 3, 2, 0],
            [0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [0, 3, 2, 9, 14])

    def test2(self):
        test_input = [
            [0, 2, 1, 0, 0],
            [0, 0, 0, 3, 4],
            [0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [7, 6, 8, 21])

    def test3(self):
        test_input = [
            [1, 2, 3, 0, 0, 0],
            [4, 5, 6, 0, 0, 0],
            [7, 8, 9, 1, 0, 0],
            [0, 0, 0, 0, 1, 2],
            [0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [1, 2, 3])

    def test4(self):
        test_input = [
            [0]
        ]
        self.assertEqual(solution(test_input), [1, 1])

    def test5(self):
        test_input = [
            [0, 0, 12, 0, 15, 0, 0, 0, 1, 8],
            [0, 0, 60, 0, 0, 7, 13, 0, 0, 0],
            [0, 15, 0, 8, 7, 0, 0, 1, 9, 0],
            [23, 0, 0, 0, 0, 1, 0, 0, 0, 0],
            [37, 35, 0, 0, 0, 0, 3, 21, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [1, 2, 3, 4, 5, 15])

    def test6(self):
        test_input = [
            [0, 7, 0, 17, 0, 1, 0, 5, 0, 2],
            [0, 0, 29, 0, 28, 0, 3, 0, 16, 0],
            [0, 3, 0, 0, 0, 1, 0, 0, 0, 0],
            [48, 0, 3, 0, 0, 0, 17, 0, 0, 0],
            [0, 6, 0, 0, 0, 1, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [4, 5, 5, 4, 2, 20])

    def test7(self):
        test_input = [
            [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [1, 1, 1, 1, 1, 5])

    def test8(self):
        test_input = [
            [1, 1, 1, 0, 1, 0, 1, 0, 1, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 0, 1, 1, 1, 0, 1, 0, 1, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 0, 1, 0, 1, 1, 1, 0, 1, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 0, 1, 0, 1, 0, 1, 1, 1, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [1, 0, 1, 0, 1, 0, 1, 0, 1, 1],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [2, 1, 1, 1, 1, 6])

    def test9(self):
        test_input = [
            [0, 86, 61, 189, 0, 18, 12, 33, 66, 39],
            [0, 0, 2, 0, 0, 1, 0, 0, 0, 0],
            [15, 187, 0, 0, 18, 23, 0, 0, 0, 0],
            [1, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [6, 44, 4, 11, 22, 13, 100])

    def test10(self):
        test_input = [
            [0, 0, 0, 0, 3, 5, 0, 0, 0, 2],
            [0, 0, 4, 0, 0, 0, 1, 0, 0, 0],
            [0, 0, 0, 4, 4, 0, 0, 0, 1, 1],
            [13, 0, 0, 0, 0, 0, 2, 0, 0, 0],
            [0, 1, 8, 7, 0, 0, 0, 1, 3, 0],
            [1, 7, 0, 0, 0, 0, 0, 2, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        ]
        self.assertEqual(solution(test_input), [1, 1, 1, 2, 5])


if __name__ == "__main__":
    unittest.main()

