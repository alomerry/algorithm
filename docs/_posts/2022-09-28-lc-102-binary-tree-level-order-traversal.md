---
title: 二叉树的层序遍历
problem_no: 102
date: 2022-09-27
categories:
  - LeetCode
tags:
  - Y2022
  - LeetCode
  - Medium
  - BFS
  - BinaryTree
---

## [Description](https://leetcode.cn/problems/binary-tree-level-order-traversal/)

Difficulty: **中等**

Related Topics: [树](https://leetcode.cn/tag/tree/), [广度优先搜索](https://leetcode.cn/tag/breadth-first-search/), [二叉树](https://leetcode.cn/tag/binary-tree/)


给你二叉树的根节点 `root` ，返回其节点值的 **层序遍历** 。 （即逐层地，从左到右访问所有节点）。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
```

**示例 2：**

```
输入：root = [1]
输出：[[1]]
```

**示例 3：**

```
输入：root = []
输出：[]
```

**提示：**

*   树中节点数目在范围 `[0, 2000]` 内
*   `-1000 <= Node.val <= 1000`


## Solution

Language: **C++**

{% include_file "https://gitee.com/alomerry/algorithm/raw/master/code/leet-code/102-main.cpp" syntax="cpp" %}
