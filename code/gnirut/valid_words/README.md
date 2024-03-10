# Valid Words

Given a set of unique caracteres ```s, k``` is an array that contains a set of valid and invalid words.

A workd is a valid if it was only generated from the set of characters in s.

Return the number of valid words, if no valid works are fond return 0.

Example 1:

```plain
Input: s = "cb", k = ["cd", "bd", "cccb", "bcc", "bcdcb"]
Output: 2
Explanation: Strings "cccb" and "bbc" are valid since they only contain characters 'c' and 'b'
```

Constrainsts:

- 1 <= k.length <= 10^4
- 1 <= s.length <= 20
- 1 <= k[i].length <= 10
- The caraccters in s are distinct
- k[i] and s contain only lowercase English letters
