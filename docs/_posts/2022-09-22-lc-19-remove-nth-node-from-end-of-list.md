---
title: 删除链表的倒数第 N 个结点
problem_no: 19
date: 2022-09-22
categories:
  - LeetCode
tags:
  - Y2022
  - LeetCode
  - Medium
---

## [Description](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

Difficulty: **中等**

Related Topics: [链表](https://leetcode.cn/tag/linked-list/), [双指针](https://leetcode.cn/tag/two-pointers/)


给你一个链表，删除链表的倒数第 `n`个结点，并且返回链表的头结点。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
```

**示例 2：**

```
输入：head = [1], n = 1
输出：[]
```

**示例 3：**

```
输入：head = [1,2], n = 1
输出：[1]
```

**提示：**

*   链表中结点的数目为 `sz`
*   `1 <= sz <= 30`
*   `0 <= Node.val <= 100`
*   `1 <= n <= sz`

**进阶：** 你能尝试使用一趟扫描实现吗？


## Solution

Language: **C++**

空间替换时间：使用 map 记录“链”。key 为链表的序号，value 记录 next。删除倒数 n 个数字时，只需要将 n-1 的 value 指向 n+1 即可。

{% include_file "https://gitee.com/alomerry/algorithm/raw/master/code/leet-code/19-main.cpp" syntax="cpp" %}
