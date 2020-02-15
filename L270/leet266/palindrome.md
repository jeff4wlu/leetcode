[LeetCode] 266. Palindrome Permutation 回文全排列 

 
Given a string, determine if a permutation of the string could form a palindrome.
Example 1:
Input: "code"
Output: false
Example 2:
Input: "aab"
Output: true
Example 3:
Input: "carerac"
Output: true
Hint:
Consider the palindromes of odd vs even length. What difference do you notice?
Count the frequency of each character.
If each character occurs even number of times, then it must be a palindrome. How about character which occurs odd number of times?
 
这道题让我们判断一个字符串的全排列有没有是回文字符串的，那么根据题目中的提示，我们分字符串的个数是奇偶的情况来讨论，如果是偶数的话，由于回文字符串的特性，每个字母出现的次数一定是偶数次，当字符串是奇数长度时，只有一个字母出现的次数是奇数，其余均为偶数，那么利用这个特性我们就可以解题，我们建立每个字母和其出现次数的映射，然后我们遍历 HashMap，统计出现次数为奇数的字母的个数，那么只有两种情况是回文数，第一种是没有出现次数为奇数的字母，再一个就是字符串长度为奇数，且只有一个出现次数为奇数的字母，参见代码如下：
 
解法一：

class Solution {
public:
    bool canPermutePalindrome(string s) {
        unordered_map<char, int> m;
        int cnt = 0;
        for (auto a : s) ++m[a];
        for (auto a : m) {
            if (a.second % 2 == 1) ++cnt;
        }
        return cnt == 0 || (s.size() % 2 == 1 && cnt == 1);
    }
};

 
那么我们再来看一种解法，这种方法用到了一个 HashSet，我们遍历字符串，如果某个字母不在 HashSet 中，我们加入这个字母，如果字母已经存在，我们删除该字母，那么最终如果 HashSet 中没有字母或是只有一个字母时，说明是回文串，参见代码如下：
 
解法二：

class Solution {
public:
    bool canPermutePalindrome(string s) {
        unordered_set<char> st;
        for (auto a : s) {
            if (!st.count(a)) st.insert(a);
            else st.erase(a);
        }
        return st.empty() || st.size() == 1;
    }
};

 
再来看一种 bitset 的解法，这种方法也很巧妙，我们建立一个 256 大小的 bitset，每个字母根据其 ASCII 码值的不同都有其对应的位置，然后我们遍历整个字符串，遇到一个字符，就将其对应的位置的二进制数 flip 一下，就是0变1，1变0，那么遍历完成后，所有出现次数为偶数的对应位置还应该为0，而出现次数为奇数的时候，对应位置就为1了，那么我们最后只要统计1的个数，就知道出现次数为奇数的字母的个数了，只要个数小于2就是回文数，参见代码如下：
 
解法三：

class Solution {
public:
    bool canPermutePalindrome(string s) {
        bitset<256> b;
        for (auto a : s) {
            b.flip(a);
        }
        return b.count() < 2;
    }
};

 
类似题目：
Longest Palindromic Substring
Valid Anagram
Palindrome Permutation II
Longest Palindromic Substring