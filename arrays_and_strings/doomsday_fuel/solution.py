def solution(m):
    if len(m) == 0 or len(m[0]) == 0:
        return []

    num_absorbing_states = 0

    for i in range(0, len(m)):
        if is_absorbing_state(m[i]):
            m[i][i] = 1
            num_absorbing_states += 1

    i = get_sub_matrix(m, len(m)-num_absorbing_states, num_absorbing_states-1)
    q = get_sub_matrix(m, 0, len(m)-num_absorbing_states-1)

    iq = sub(i, q)
    f = get_f(iq)

    fr = product(f, get_sub_matrix(m, len(m)-num_absorbing_states-1, len(m)))

    print(fr[0])

    # [0, 3, 2, 9, 14]


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
    right, left = 1, 1

    for i in range(0, len(iq)):
        right *= iq[i][i]
        left *= iq[len(iq)-i-1][i]

    det = right - left
    trasp = [[iq[j][i] for j in range(len(iq))] for i in range(len(iq[0]))]

    for i in range(0, len(trasp)):
        for j in range(0, len(trasp[i])):
            trasp[i][j] = 1 / det * trasp[i][j]

    return trasp


def product(m1, m2):
    return [[sum(x * y for x, y in zip(m1_r, m2_c)) for m2_c in zip(*m2)] for m1_r in m1]


solution([[0, 1, 0, 0, 0, 1], [4, 0, 0, 3, 2, 0], [0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0]])
