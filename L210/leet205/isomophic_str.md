[LeetCode] 205. Isomorphic Strings 同构字符串 

 
Given two strings s and t, determine if they are isomorphic.
Two strings are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.
Example 1:
Input: s = "egg", t = "add"
Output: true
Example 2:
Input: s = "foo", t = "bar"
Output: false
Example 3:
Input: s = "paper", t = "title"
Output: true
Note:
You may assume both s and t have the same length.
  
这道题让我们求同构字符串，就是说原字符串中的每个字符可由另外一个字符替代，可以被其本身替代，相同的字符一定要被同一个字符替代，且一个字符不能被多个字符替代，即不能出现一对多的映射。根据一对一映射的特点，需要用两个 HashMap 分别来记录原字符串和目标字符串中字符出现情况，由于 ASCII 码只有 256 个字符，所以可以用一个 256 大小的数组来代替 HashMap，并初始化为0，遍历原字符串，分别从源字符串和目标字符串取出一个字符，然后分别在两个数组中查找其值，若不相等，则返回 false，若相等，将其值更新为 i + 1，因为默认的值是0，所以更新值为 i + 1，这样当 i=0 时，则映射为1，如果不加1的话，那么就无法区分是否更新了，代码如下：
 

class Solution {
public:
    bool isIsomorphic(string s, string t) {
        int m1[256] = {0}, m2[256] = {0}, n = s.size();
        for (int i = 0; i < n; ++i) {
            if (m1[s[i]] != m2[t[i]]) return false;
            m1[s[i]] = i + 1;
            m2[t[i]] = i + 1;
        }
        return true;
    }
};
