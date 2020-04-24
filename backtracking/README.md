### 回溯法思想(DFS)：

解决一个回溯问题，实际上就是一个决策树的遍历过程。你只需要思考 3 个问题：
1、路径：也就是已经做出的选择。
2、选择列表：也就是你当前可以做的选择。
3、结束条件：也就是到达决策树底层，无法再做选择的条件。

```sh
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
​
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```

例子：

* Binary Watch
* 全排列
* N皇后

### BFS思想：

```python
# 计算从起点 start 到终点 target 的最近距离
def BFS(start, target): {
    q = []       # 核心数据结构
    visited = {} # 避免走回头路

    q.offer(start) # 将起点加入队列
    visited.add(start)
    step = 0 # 记录扩散的步数

    while q:
        sz = len(q)
        # 将当前队列中的所有节点向四周扩散 
        for _ in range(sz): 
            cur = q.pop(0)
            # 划重点：这里判断是否到达终点
            if cur == target:
                return step
            # 将 cur 的相邻节点加入队列, 如果是数则是孩子节点。
            for x in cur.adj():
                if x not in visited:
                    q.offer(x);
                    visited.add(x);
        # 划重点：更新步数在这里
        step++;
    }
}
```

### Ref

* https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/hui-su-suan-fa-xiang-jie-xiu-ding-ban
* https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485134&idx=1&sn=fd345f8a93dc4444bcc65c57bb46fc35&chksm=9bd7f8c6aca071d04c4d383f96f2b567ad44dc3e67d1c3926ec92d6a3bcc3273de138b36a0d9&scene=21#wechat_redirect