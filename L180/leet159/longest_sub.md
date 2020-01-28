[LeetCode] 159. Longest Substring with At Most Two Distinct Characters 最多有两个不同字符的最长子串 

 
Given a string s , find the length of the longest substring t  that contains at most 2 distinct characters.
Example 1:
Input: "eceba"
Output: 3
Explanation: tis "ece" which its length is 3.
Example 2:
Input: "ccaabbb"
Output: 5
Explanation: tis "aabbb" which its length is 5.
 
这道题给我们一个字符串，让我们求最多有两个不同字符的最长子串。那么我们首先想到的是用 HashMap 来做，HashMap 记录每个字符的出现次数，然后如果 HashMap 中的映射数量超过两个的时候，我们需要删掉一个映射，比如此时 HashMap 中e有2个，c有1个，此时把b也存入了 HashMap，那么就有三对映射了，这时我们的 left 是0，先从e开始，映射值减1，此时e还有1个，不删除，left 自增1。这时 HashMap 里还有三对映射，此时 left 是1，那么到c了，映射值减1，此时e映射为0，将e从 HashMap 中删除，left 自增1，然后我们更新结果为 i - left + 1，以此类推直至遍历完整个字符串，参见代码如下：
 
解法一：

class Solution {
public:
    int lengthOfLongestSubstringTwoDistinct(string s) {
        int res = 0, left = 0;
        unordered_map<char, int> m;
        for (int i = 0; i < s.size(); ++i) {
            ++m[s[i]];
            while (m.size() > 2) {
                if (--m[s[left]] == 0) m.erase(s[left]);
                ++left;
            }
            res = max(res, i - left + 1);
        }
        return res;
    }
};

 
我们除了用 HashMap 来映射字符出现的个数，我们还可以映射每个字符最新的坐标，比如题目中的例子 "eceba"，遇到第一个e，映射其坐标0，遇到c，映射其坐标1，遇到第二个e时，映射其坐标2，当遇到b时，映射其坐标3，每次我们都判断当前 HashMap 中的映射数，如果大于2的时候，那么需要删掉一个映射，我们还是从 left=0 时开始向右找，看每个字符在 HashMap 中的映射值是否等于当前坐标 left，比如第一个e，HashMap 此时映射值为2，不等于 left 的0，那么 left 自增1，遇到c的时候，HashMap 中c的映射值是1，和此时的 left 相同，那么我们把c删掉，left 自增1，再更新结果，以此类推直至遍历完整个字符串，参见代码如下：
 
解法二：

class Solution {
public:
    int lengthOfLongestSubstringTwoDistinct(string s) {
        int res = 0, left = 0;
        unordered_map<char, int> m;
        for (int i = 0; i < s.size(); ++i) {
            m[s[i]] = i;
            while (m.size() > 2) {
                if (m[s[left]] == left) m.erase(s[left]);
                ++left;
            }
            res = max(res, i - left + 1);
        }
        return res;
    }
};

 
后来又在网上看到了一种解法，这种解法是维护一个 sliding window，指针 left 指向起始位置，right 指向 window 的最后一个位置，用于定位 left 的下一个跳转位置，思路如下：
1. 若当前字符和前一个字符相同，继续循环。
2. 若不同，看当前字符和 right 指的字符是否相同
    (1) 若相同，left 不变，右边跳到 i - 1
    (2) 若不同，更新结果，left 变为 right+1，right 变为 i - 1
最后需要注意在循环结束后，我们还要比较结果 res 和 s.size() - left 的大小，返回大的，这是由于如果字符串是 "ecebaaa"，那么当 left=3 时，i=5,6 的时候，都是继续循环，当i加到7时，跳出了循环，而此时正确答案应为 "baaa" 这4个字符，而我们的结果 res 只更新到了 "ece" 这3个字符，所以我们最后要判断 s.size() - left 和结果 res 的大小。
另外需要说明的是这种解法仅适用于于不同字符数为2个的情况，如果为k个的话，还是需要用上面两种解法。
 
解法三：

class Solution {
public:
    int lengthOfLongestSubstringTwoDistinct(string s) {
        int left = 0, right = -1, res = 0;
        for (int i = 1; i < s.size(); ++i) {
            if (s[i] == s[i - 1]) continue;
            if (right >= 0 && s[right] != s[i]) {
                res = max(res, i - left);
                left = right + 1;
            }
            right = i - 1;
        }
        return max(s.size() - left, res);
    }
};

 
还有一种不使用 HashMap 的解法，是在做 Fruit Into Baskets 这道题的时候在论坛上看到的，其实这两道题除了背景设定之外没有任何的区别，代码基本上都可以拷来直接用的。这里我们使用若干的变量，其中 cur 为当前最长子串的长度，first 和 second 为当前候选子串中的两个不同的字符，cntLast 为 second 字符的连续长度。我们遍历所有字符，假如遇到的字符是 first 和 second 中的任意一个，那么 cur 可以自增1，否则 cntLast 自增1，因为若是新字符的话，默认已经将 first 字符淘汰了，此时候选字符串由 second 字符和这个新字符构成，所以当前长度是 cntLast+1。然后再来更新 cntLast，假如当前字符等于 second 的话，cntLast 自增1，否则均重置为1，因为 cntLast 统计的就是 second 字符的连续长度。然后再来判断若当前字符不等于 second，则此时 first 赋值为 second， second 赋值为新字符。最后不要忘了用 cur 来更新结果 res，参见代码如下：
 
解法四：

class Solution {
public:
    int lengthOfLongestSubstringTwoDistinct(string s) {
        int res = 0, cur = 0, cntLast = 0;
        char first, second;
        for (char c : s) {
            cur = (c == first || c == second) ? cur + 1 : cntLast + 1;
            cntLast = (c == second) ? cntLast + 1 : 1;
            if (c != second) {
                first = second; second = c;
            }
            res = max(res, cur);
        }
        return res;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/159
 
类似题目：
Fruit Into Baskets
Longest Substring Without Repeating Characters
Sliding Window Maximum
Longest Substring with At Most K Distinct Characters
Subarrays with K Different Integers
 