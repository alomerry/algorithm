---
title: 合并两个有序链表
problem_no: 21
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

将两个升序链表合并为一个新的 `升序` 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：

```text
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
```

示例 2：

```text
输入：l1 = [], l2 = []
输出：[]
```

示例 3：

```text
输入：l1 = [], l2 = [0]
输出：[0]
```

提示：

- 两个链表的节点数目范围是 `[0, 50]`
- `-100 <= Node.val <= 100`
- `l1` 和 `l2` 均按 **非递减顺序** 排列

## Solution

## Code

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        if (l1 == NULL) {
            return l2;
        }
        if (l2 == NULL){
            return l1;
        }

        if (l1->val <= l2->val) {
            l1->next = mergeTwoLists(l1->next, l2);
            return l1;
        }else{
            l2->next = mergeTwoLists(l1, l2->next);
            return l2;
        }
    }
};
```
