# Perfect Matches

Given an array of integers, containig some perfect matches. Find out these perfect matches and return the result.

A perfect match should satisfy the following conditions:

1. if nums[n] == nums[m]
2. 'n' is less than 'm'

Example 1:

```plain
Input nums = [1, 1, 1, 3, 2, 3]
Output: 4
Explanation: There are 4 perfect matches (0, 1), (0, 2), (1, 2), (3, 5)
```

Example 2:

```plain
Input nums = [4, 5, 6]
Output: 0
Explanation: No match found
```

Constrainsts:

- 1 <= nums.length <= 100
- 1 <= nums[i] <= 100
