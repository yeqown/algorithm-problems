class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def gen_linked_list(l: list) -> ListNode:
    head = None
    next = None
    for v in l:
        node = ListNode(v)
        if head == None:
            next = node
            head = node
        else:
            next.next = node
            next = node
    return head


def print_linked_list(root: ListNode):
    l = []
    while root is not None:
        l.append(root.val)
        root = root.next

    print(l)

# 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
# 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
#  
# 进阶：
# 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
#  

# 示例：
# 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
# 输出：7 -> 8 -> 0 -> 7


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 is None:
            return l2
        if l2 is None:
            return l1

        l1r = self.reverse(l1)
        l2r = self.reverse(l2)
        print_linked_list(l1r)
        print_linked_list(l2r)

        ptr = None
        head = None
        c = 0

        while l1r != None and l2r != None:
            s = l1r.val + l2r.val + c
            v = s % 10
            c = int(s / 10)

            print(l1r.val, l2r.val, c, s, v, c)
            node = ListNode(v)
            if head is None:
                head = node
                ptr = node
            else:
                ptr.next = node
                ptr = node

            l1r = l1r.next
            l2r = l2r.next

        while l1r is not None:
            s = l1r.val + c
            v = s % 10
            c = int(s / 10)
            node = ListNode(v)
            ptr.next = node
            ptr = node
            l1r = l1r.next

        while l2r is not None:
            s = l2r.val + c
            v = s % 10
            c = int(s / 10)
            node = ListNode(v)
            ptr.next = node
            ptr = node
            l2r = l2r.next

        if c != 0:
            ptr.next = ListNode(c)

        return self.reverse(head)

    def reverse(self, root: ListNode) -> ListNode:
        if root is None:
            return None

        pre = None
        next = None

        while root is not None:
            next = root.next
            root.next = pre
            pre = root
            root = next
        return pre


class Solution2:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 is None:
            return l2
        if l2 is None:
            return l1
        # TODO:
        return None


if __name__ == "__main__":
    l1 = gen_linked_list([7, 2, 4, 3])
    l2 = gen_linked_list([5, 6, 4])
    # print_linked_list(l1)
    # print_linked_list(l2)

    sol = Solution()

    # print_linked_list(sol.reverse(l1))
    print_linked_list(sol.addTwoNumbers(l1, l2))
