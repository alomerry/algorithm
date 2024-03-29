---
title: 有效的括号
problem_no: 20
date: 2021-09-25
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

Source: [LeetCode 20](https://leetcode-cn.com/problems/valid-parentheses/){:target="_blank"}

### Description

给定一个只包括 `'('`，`')'`，`'{'`，`'}'`，`'['`，`']'` 的字符串 `s` ，判断字符串是否有效。

有效字符串需满足：

- 左括号必须用相同类型的右括号闭合。
- 左括号必须以正确的顺序闭合。

示例 1：

```text
输入：s = "()"
输出：true
```

示例 2：

```text
输入：s = "()[]{}"
输出：true
```

示例 3：

```text
输入：s = "(]"
输出：false
```

示例 4：

```text
输入：s = "([)]"
输出：false
```

示例 5：

```text
输入：s = "{[]}"
输出：true
```

提示：

- $1 <= s.length <= 10^4$
- `s` 仅由括号 `'()[]{}'` 组成

## Solution

## Code

```cpp
class Solution {
public:
    bool isValid(string s) {
        set<char> left;
        string ls = "({[",rs = ")}]";
        map<char,char> m;

        for (int i = 0; i < ls.size(); i++){
            left.insert(ls[i]);
            m[ls[i]] = rs[i];
        }

        stack<char> st;
        set<char>::iterator it;
        for (int i = 0; i < s.size(); i++){
            it = left.find(s[i]);
            if (it != left.end()){
                st.push(s[i]);
                continue;
            }
            if (st.empty()){
                return false;
            }
            char top = st.top();
            st.pop();
            if (m[top] != s[i]){
                return false;
            }
        }
    return st.empty();
}
};
```
