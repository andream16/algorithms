import unittest
from fractions import Fraction, gcd
from functools import reduce


def solution(m):

    n_rows = len(m)

    if n_rows == 0 or len(m[0]) == 0:
        return []

    num_absorbing_states = 0

    for i in range(0, n_rows):
        sum_row = 0
        if is_absorbing_state(m[i]):
            m[i][i] = 1
            num_absorbing_states += 1
            continue
        else:
            for j in range(0, len(m[i])):
                sum_row += m[i][j]
            for j in range(0, len(m[i])):
                if m[i][j] != 0:
                    m[i][j] = get_fraction(float(m[i][j]) / float(sum_row))

    if num_absorbing_states == 1:
        return [1, 1]

    # F = (I-Q)^-1
    i = sub_matrix(m, n_rows-num_absorbing_states, n_rows-num_absorbing_states, n_rows, n_rows)
    q = sub_matrix(m, 0, 0, n_rows-num_absorbing_states, n_rows-num_absorbing_states)
    # I-Q
    diff = sub(i, q)
    # (diff)^-1
    f = get_f(diff)
    r = sub_matrix(m, 0, n_rows-num_absorbing_states, n_rows-num_absorbing_states, n_rows)
    fr = product(f, r)

    if len(fr) == 0:
        return []

    denominators = []
    for n in fr[0]:
        if n != 0:
            denominators.append(n.denominator)

    # Least Common Multiple
    denominator = reduce(lambda a, b: a*b // gcd(a, b), denominators)

    res = []
    for n in fr[0]:
        if n == 0:
            res.append(0)
        else:
            res.append(int(n.numerator * (denominator / n.denominator)))

    res.append(denominator)
    return res


def is_absorbing_state(nums):
    for n in nums:
        if n != 0:
            return False
    return True


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
    min_row, min_col = min(len(i), len(q)), min(len(i[0]), len(q[0]))
    res = []
    for k in range(0, min_row):
        curr = []
        for j in range(0, min_col):
            curr.append(i[k][j]-q[k][j])
        res.append(curr)
    return res


def get_f(m):
    determinant = get_determinant(m)
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


def get_fraction(n):
    return Fraction(n).limit_denominator(100000000000)


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


if __name__ == "__main__":
    unittest.main()

