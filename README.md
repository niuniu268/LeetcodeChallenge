# LEETCODE PRACTICES


## Array

### Bisection

### Subarray

### Remove Element

### Sorted Squares

### Spiral Matrix

![img](https://code-thinking-1253855093.file.myqcloud.com/pics/20220922102236.png)

## List

## Binary Tree

![img](https://camo.githubusercontent.com/3fd14087ae6ed14819b3ee6443ed5125161da981dbb8f2b48c10887cd9a5cfb7/68747470733a2f2f696d672d626c6f672e6373646e696d672e636e2f32303231303231393139303830393435312e706e67)

## Reverse-search Algorithm

![img](https://camo.githubusercontent.com/1531017a62378c14e8731434dbd48ad05a8336606ea72927416ee87191e2bd0e/68747470733a2f2f696d672d626c6f672e6373646e696d672e636e2f32303231303231393139323035303636362e706e67)
![img](https://code-thinking-1253855093.file.myqcloud.com/pics/20210130173631174.png)
```
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}

```
## Greedy Algorithm

![img](https://camo.githubusercontent.com/a77a640405ca924ee1a18028fef3c05c21c15df33f1d91852a3f9f64d37b7366/68747470733a2f2f636f64652d7468696e6b696e672d313235333835353039332e66696c652e6d7971636c6f75642e636f6d2f706963732f32303231303931373130343331352e706e67)

## Dynamic Programming

![img](https://camo.githubusercontent.com/173c771640343e8fd7c0429d56f7fbd1d02bfef90c0bd1d409df945fe668ff79/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f706963732fe58aa8e68081e8a784e588922de680bbe7bb93e5a4a7e7bab2312e6a7067)

Logical Procedure:

1. 确定dp数组（dp table）以及下标的含义
2. 确定递推公式
3. dp数组如何初始化
4. 确定遍历顺序
5. 举例推导dp数组

### Backpack problem

![img](https://camo.githubusercontent.com/a671fd4e891752c2b447ec882cd8bf075e1f603122964fde0797c22cc80da5bf/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f706963732fe58aa8e68081e8a784e588922de8838ce58c85e997aee9a298e680bbe7bb932e706e67)

- 问能否能装满背包（或者最多装多少）：dp[j] = max(dp[j], dp[j - nums[i]] + nums[i]); ，对应题目如下：

- 问装满背包有几种方法：dp[j] += dp[j - nums[i]] ，对应题目如下：

- 问背包装满最大价值：dp[j] = max(dp[j], dp[j - weight[i]] + value[i]); ，对应题目如下：

- 问装满背包所有物品的最小个数：dp[j] = min(dp[j - coins[i]] + 1, dp[j]); ，对应题目如下：

- 01 backpack
  一维dp数组的背包在遍历顺序上和二维dp数组实现的01背包其实是有很大差异的，大家需要注意！

- complete backpack
  如果求组合数就是外层for循环遍历物品，内层for遍历背包。

  如果求排列数就是外层for遍历背包，内层for循环遍历物品

### Stock problem

![img](https://camo.githubusercontent.com/006c670b6d92f56c90ef5de18883717ceff43526f28e40c82a77e017b1f30cee/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f706963732fe882a1e7a5a8e997aee9a298e680bbe7bb932e6a7067)

### Subsequence problem

![img](https://camo.githubusercontent.com/8766139b53670fa351812ce4d6208914c614e5417026e15dff36802f59dbc86e/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f706963732fe58aa8e68081e8a784e588922de5ad90e5ba8fe58897e997aee9a298e680bbe7bb932e6a7067)

### DFS

```
void dfs(参数) {
    处理节点
    dfs(图，选择的节点); // 递归
    回溯，撤销处理结果
}
```
BackTracking: 

```
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }
    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}

```
Framework:

```
void dfs(参数) {
if (终止条件) {
存放结果;
return;
}

    for (选择：本节点所连接的其他节点) {
        处理节点;
        dfs(图，选择的节点); // 递归
        回溯，撤销处理结果
    }
}
```