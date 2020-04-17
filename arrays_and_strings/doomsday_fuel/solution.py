from fractions import Fraction, gcd
from functools import reduce


def solution(m):
    if len(m) == 0 or len(m[0]) == 0:
        return []

    num_absorbing_states = 0

    for i in range(0, len(m)):
        sum_row = 0
        if is_absorbing_state(m[i]):
            m[i][i] = 1
            num_absorbing_states += 1
        else:
            for j in range(0, len(m[i])):
                sum_row += m[i][j]
            for j in range(0, len(m[i])):
                if m[i][j] != 0:
                    m[i][j] = float(m[i][j]) / float(sum_row)

    i = get_sub_matrix(m, len(m)-num_absorbing_states, num_absorbing_states-1)
    q = get_sub_matrix(m, 0, len(m)-num_absorbing_states-1)

    iq = sub(i, q)

    f = get_f(iq)

    r = sub_matrix(m, 0, len(m)-num_absorbing_states, len(m)-num_absorbing_states-1, len(m))

    fr = product(f, r)

    denominators = []
    for n in fr[0]:
        if n != 0:
            denominators.append(Fraction(n).limit_denominator(100000000000).denominator)

    denominator = reduce(lambda a, b: a*b // gcd(a, b), denominators)

    res = []
    for n in fr[0]:
        if n == 0:
            res.append(0)
        else:
            c = Fraction(n).limit_denominator(100000000000)
            res.append(int(c.numerator * (denominator / c.denominator)))

    res.append(denominator)
    return res


def get_sub_matrix(m, begin, end):
    nm = [[]]
    cr = 0
    found = False
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            if begin <= i <= end and begin <= j <= end:
                if len(nm[cr]) == 0:
                    nm.append([])
                nm[cr].append(m[i][j])
                found = True
        if found:
            cr += 1
            found = False
    return nm[:len(nm)-1]


def sub_matrix(m, r1, c1, r2, c2):
    sub_m = [[]]
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            if r1 <= i <= r2 and c1 <= j <= c2:
                if len(sub_m[i]) == 0:
                    sub_m.append([])
                sub_m[i].append(m[i][j])
    return sub_m[:len(sub_m)-1]


def is_absorbing_state(nums):
    for n in nums:
        if n != 0:
            return False
    return True


def sub(i, q):
    res = [[]]
    for k in range(0, len(i)):
        for j in range(0, len(i[k])):
            if len(res[k]) == 0:
                res.append([])
            res[k].append(i[k][j]-q[k][j])
    return res[:len(res)-1]


def get_f(iq):

    det = 1 / get_determinant(iq)

    trasp = traspose(iq)

    for i in range(0, len(trasp)):
        for j in range(0, len(trasp[i])):
            trasp[i][j] = det * trasp[i][j]

    return trasp


def traspose(m):
    trasp = []
    if len(m) == 2 and len(m[0]) == 2:
        trasp = [[m[1][1], -m[0][1]], [-m[1][0], m[0][0]]]
    return trasp


def get_matrix_minor(m, i, j):
    return [row[:j] + row[j+1:] for row in (m[:i]+m[i+1:])]


def get_determinant(m):
    if len(m) == 2:
        return m[0][0]*m[1][1]-m[0][1]*m[1][0]

    determinant = 0
    for c in range(len(m)):
        determinant += ((-1)**c)*m[0][c]*get_determinant(get_matrix_minor(m, 0, c))
    return determinant


def product(m1, m2):
    return [[sum(x * y for x, y in zip(m1_r, m2_c)) for m2_c in zip(*m2)] for m1_r in m1]


def main():
    # [0, 3, 2, 9, 14]
    print(solution([[0, 1, 0, 0, 0, 1], [4, 0, 0, 3, 2, 0], [0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0]]))
    # [7, 6, 8, 21]
    print(solution([[0, 2, 1, 0, 0], [0, 0, 0, 3, 4], [0, 0, 0, 0, 0], [0, 0, 0, 0,0], [0, 0, 0, 0, 0]]))


if __name__ == "__main__":
    main()
