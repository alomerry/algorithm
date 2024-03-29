---
title: 二叉树的层序遍历 II
problem_no: 107
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

## [Description](https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/)

Difficulty: **中等**

Related Topics: [树](https://leetcode.cn/tag/tree/), [广度优先搜索](https://leetcode.cn/tag/breadth-first-search/), [二叉树](https://leetcode.cn/tag/binary-tree/)


给你二叉树的根节点 `root` ，返回其节点值 **自底向上的层序遍历** 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
输入：root = [3,9,20,null,null,15,7]
输出：[[15,7],[9,20],[3]]
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

{% include_file "https://gitee.com/alomerry/algorithm/raw/master/code/leet-code/107-main.cpp" syntax="cpp" %}
