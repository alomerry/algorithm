---
title: 括号生成
problem_no: 22
date: 2023-01-17
categories:
  - LeetCode
tags:
  - Y2023
  - LeetCode
  - Medium
  - DP
  - String
  - Backtracking
---

## [Description](https://leetcode.cn/problems/generate-parentheses/)

Difficulty: **中等**

Related Topics: [字符串](https://leetcode.cn/tag/string/), [动态规划](https://leetcode.cn/tag/dynamic-programming/), [回溯](https://leetcode.cn/tag/backtracking/)

数字 `n` 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 **有效的
** 括号组合。

**示例 1：**

```
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
```

**示例 2：**

```
输入：n = 1
输出：["()"]
```

**提示：**

* `1 <= n <= 8`

## Solution

Language: **Golang**

{% include_file "https://gitee.com/alomerry/algorithm/raw/master/code/leet-code/22-main.go" syntax="go" %}
