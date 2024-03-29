---
title: 电话号码的字母组合
problem_no: 17
date: 2023-01-23
categories:
  - LeetCode
tags:
  - Y2023
  - LeetCode
  - String
  - Medium
  - Backtracking
  - HashTable
---

## [Description](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)

Difficulty: **中等**

Related Topics: [哈希表](https://leetcode.cn/tag/hash-table/), [字符串](https://leetcode.cn/tag/string/), [回溯](https://leetcode.cn/tag/backtracking/)


给定一个仅包含数字`2-9`的字符串，返回所有它能表示的字母组合。答案可以按 **任意顺序** 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/11/09/200px-telephone-keypad2svg.png)

**示例 1：**

```
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**示例 2：**

```
输入：digits = ""
输出：[]
```

**示例 3：**

```
输入：digits = "2"
输出：["a","b","c"]
```

**提示：**

*   `0 <= digits.length <= 4`
*   `digits[i]` 是范围 `['2', '9']` 的一个数字。


## Solution

Language: **Golang**

递归遍历

{% include_file "https://gitee.com/alomerry/algorithm/raw/master/code/leet-code/17-main.go" syntax="go" %}
