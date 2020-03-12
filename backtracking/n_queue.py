# 给你一个 N×N 的棋盘，让你放置 N 个皇后，使得它们不能互相攻击。
# PS：皇后可以攻击同一行、同一列、同一斜线的任意单位。
import copy


# 工具函数
def mat_constructor(row: int, col: int, default_value: int):
    mat = []
    for _ in range(row):
        r = [default_value for _ in range(col)]
        mat.append(r)
    return mat


# 工具函数
def print_mat(mat, without_wrapper: bool):
    if without_wrapper:
        N = len(mat) - 2

        for idx, row in enumerate(mat):
            if idx == 0 or idx == N+1:
                continue
            print(row[1:N+1])
        return

    # with wrapper
    for row in mat:
        print(row)


# 工具函数
def debug_print(*args):
    if debug:
        print(*args)


debug = True
result = []


def permutate_n_queue(N: int):
    # calculate result
    choice = mat_constructor(N+2, N+2, 0)
    backtrack(N, choice, N)

    # output result
    for idx, mat in enumerate(result):
        print("solution: ", idx + 1)
        print_mat(mat, True)


# 改为斜线检测
def is_valid(n, i, choice, N) -> bool:
    # 左上左下 [仅左下]
    step = i - 1
    for vn in range(n+1, N+1):
        if step < 0:
            break
        if choice[vn][step] != 0:
            debug_print("(%d, %d) 左下不合法" % (n, i))
            return False
        step -= 1

    # 右上右下 [仅右下]
    step = i + 1
    for vn in range(n+1, N+1):
        if step > N+1:
            break
        if choice[vn][step] != 0:
            debug_print("(%d, %d) 右下不合法" % (n, i))
            return False
        step += 1

    # 同一行, 不必要
    # if sum(choice[n]) != 0:
    #     return False

    # 同一列
    # 只需要考虑n以下的列
    for vn in range(n+1, N+1):
        if choice[vn][i] != 0:
            debug_print("(%d, %d) 该列不合法" % (vn, i))
            return False

    return True


# 从 N -> 1
def backtrack(n: int, choice, N: int):
    if n == 0:
        # 找到一种可能行
        # backup = choice.copy()           # 浅拷贝
        deep_copy = copy.deepcopy(choice)  # 深拷贝
        result.append(deep_copy)
        # print_mat(backup)
        return

    # 存在外层防溢出层
    for i in range(1, N+1):
        # 非法选择筛选，剪枝
        # 同一行，同一列，左上，左下，右下，右上
        valid = is_valid(n, i, choice, N)

        # debug日志
        # if debug:
        #     print_mat(choice, True)
        # debug_print(n, i, valid)

        if not valid:
            continue

        choice[n][i] = 1
        backtrack(n-1, choice, N)
        choice[n][i] = 0


if __name__ == "__main__":
    debug = False
    permutate_n_queue(8)
