---
title: 搜索插入位置
date: 2020-08-18
problem_no: 35
categories:
  - LeetCode
tags:
  - Y2020
  - LeetCode
  - Medium
---

<!-- Description. -->

<!-- more -->

## Problem

Source: [LeetCode 35](https://leetcode-cn.com/problems/search-insert-position/){:target="_blank"}

## Description

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 `O(log n)` 的算法。

示例 1:

```text
输入: nums = [1,3,5,6], target = 5
输出: 2
```

示例 2:

```text
输入: nums = [1,3,5,6], target = 2
输出: 1
```

示例 3:

```text
输入: nums = [1,3,5,6], target = 7
输出: 4
```

示例 4:

```text
输入: nums = [1,3,5,6], target = 0
输出: 0
```

示例 5:

```text
输入: nums = [1], target = 0
输出: 0
```

提示:

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 为 **无重复元素** 的 **升序** 排列数组
- `-10^4 <= target <= 10^4`

## Solution

二分查找

### Code

```cpp
int binarySearch(vector<int> nums, int left, int right, int target) {
    if (target < nums[left])
        return left;
    if (target > nums[right])
        return right + 1;
    if (left >= right) {
        if (target <= nums[left]) {
            return left;
        } else {
            return left + 1;
        }
    }
    int middle = left + (right - left) / 2;
    if (target == nums[middle])return middle;
    else if (target < nums[middle]) {
        return binarySearch(nums, left, middle, target);
    } else {
        return binarySearch(nums, middle + 1, right, target);
    }
}

int searchInsert(vector<int> &nums, int target) {
     return binarySearch(nums, 0, nums.size() - 1, target);
}
```
