[LeetCode] 28. Implement strStr() 实现strStr()函数 

 
Implement strStr().
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:
What should we return when needle is an empty string? This is a great question to ask during an interview.
For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
 
这道题让我们在一个字符串中找另一个字符串第一次出现的位置，那首先要做一些判断，如果子字符串为空，则返回0，如果子字符串长度大于母字符串长度，则返回 -1。然后开始遍历母字符串，这里并不需要遍历整个母字符串，而是遍历到剩下的长度和子字符串相等的位置即可，这样可以提高运算效率。然后对于每一个字符，都遍历一遍子字符串，一个一个字符的对应比较，如果对应位置有不等的，则跳出循环，如果一直都没有跳出循环，则说明子字符串出现了，则返回起始位置即可，代码如下：
 

class Solution {
public:
    int strStr(string haystack, string needle) {
        if (needle.empty()) return 0;
        int m = haystack.size(), n = needle.size();
        if (m < n) return -1;
        for (int i = 0; i <= m - n; ++i) {
            int j = 0;
            for (j = 0; j < n; ++j) {
                if (haystack[i + j] != needle[j]) break;
            }
            if (j == n) return i;
        }
        return -1;
    }
};

 
我们也可以写的更加简洁一些，开头直接套两个 for 循环，不写终止条件，然后判断假如j到达 needle 的末尾了，此时返回i；若此时 i+j 到达 haystack 的长度了，返回 -1；否则若当前对应的字符不匹配，直接跳出当前循环，参见代码如下：
 
解法二：

class Solution {
public:
    int strStr(string haystack, string needle) {
        for (int i = 0; ; ++i) {
            for (int j = 0; ; ++j) {
                if (j == needle.size()) return i;
                if (i + j == haystack.size()) return -1;
                if (needle[j] != haystack[i + j]) break;
            }
        }
        return -1;
    }
};