---
title: 对称二叉树
problem_no: 101
date: 2022-09-27
categories:
  - LeetCode
tags:
  - Y2022
  - LeetCode
  - Easy
  - DFS
  - BFS
  - BinaryTree
---

## [Description](https://leetcode.cn/problems/symmetric-tree/)

Difficulty: **简单**

Related Topics: [树](https://leetcode.cn/tag/tree/), [深度优先搜索](https://leetcode.cn/tag/depth-first-search/), [广度优先搜索](https://leetcode.cn/tag/breadth-first-search/), [二叉树](https://leetcode.cn/tag/binary-tree/)


给你一个二叉树的根节点 `root` ， 检查它是否轴对称。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg)

```
输入：root = [1,2,2,3,4,4,3]
输出：true
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg)

```
输入：root = [1,2,2,null,3,null,3]
输出：false
```

**提示：**

*   树中节点数目在范围 `[1, 1000]` 内
*   `-100 <= Node.val <= 100`

**进阶：**你可以运用递归和迭代两种方法解决这个问题吗？


## Solution

Language: **C++**

{% include_file "https://gitee.com/alomerry/algorithm/raw/master/code/leet-code/101-main.cpp" syntax="cpp" %}
