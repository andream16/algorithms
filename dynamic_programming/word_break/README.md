Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.
- Note that you are allowed to reuse a dictionary word.

Example 1:

- Input: `s = "leetcode"`, `wordDict = ["leet", "code"]`
- Output: `true`
- Explanation: Return true because `"leetcode"` can be segmented as `"leet code"`.

Example 2:

- Input: `s = "applepenapple"`, `wordDict = ["apple", "pen"]`
- Output: `true`
- Explanation: Return `true` because `"applepenapple"` can be segmented as `"apple pen apple"`.

Example 3:

- Input: `s = "catsandog"`, `wordDict = ["cats", "dog", "sand", "and", "cat"]`
- Output: `false`

Tags: [`Top 100 Questions`]

Companies [`Amazon`, `Facebook`, `Bloomberg`, `Google`, `Apple`, `Oracle`, `Microsoft`, `Qualtrics`, `Adobe`, 
`Snapchat`, `eBay`, `GoDaddy`]

Source `Leetcode` - https://leetcode.com/problems/word-break/
