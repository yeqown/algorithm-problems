# 判断两个字符串是否是编辑距离为1或者0
#


def oneEditAway(first: str, second: str) -> bool:
    l1 = len(first)
    l2 = len(second)
    if abs(l1-l2) > 1:
        return False

    i = j = 0
    edited = False
    # print(i, j, l1, l2)

    while i < l1 and j < l2:
        # print(i, j, first[i], second[j], edited)
        if first[i] != second[j]:
            if not edited:
                edited = True
                # 移动更长字符串的指针
                if l1 > l2:
                    i += 1
                elif l1 == l2:
                    j += 1
                    i += 1
                else:
                    j += 1
            else:
                return False
        else:
            i += 1
            j += 1

    return True


if __name__ == "__main__":
    # res = oneEditAway("orse", "ros")
    print(oneEditAway("orse", "orse"))  # true
    print(oneEditAway("orse", "ors"))  # true
    print(oneEditAway("ore", "ors"))  # true
    print(oneEditAway("orse", "orsee"))  # true
    print(oneEditAway("orse", "rseo"))  # false
    print(oneEditAway("teacher", "beacher"))  # true
