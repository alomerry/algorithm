---
title: 最后一个单词的长度
problem_no: 58
date: 2021-09-30
categories:
  - LeetCode
tags:
  - Y2021
  - LeetCode
  - Easy
---

<!-- Description. -->

<!-- more -->

## Problem

Source: [LeetCode 58](https://leetcode-cn.com/problems/length-of-last-word/){:target="_blank"}

### Description

给你一个字符串 `s`，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。

**单词** 是指仅由字母组成、不包含任何空格字符的最大子字符串。

示例 1：

```text
输入：s = "Hello World"
输出：5
```

示例 2：

```text
输入：s = "   fly me   to   the moon  "
输出：4
```

示例 3：

```text
输入：s = "luffy is still joyboy"
输出：6
```

提示：

- $1 <= s.length <= 10^4$
- `s` 仅有英文字母和空格 `' '` 组成
- `s` 中至少存在一个单词

## Solution

## Code

```cpp
class Solution {
public:
    int lengthOfLastWord(string s) {
        int len = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == ' ') {
                len = -abs(len);
                continue;
            }
            if (len < 0 ){
                len = 0;
            }
            len++;
        }
        return abs(len);
    }
};
```
