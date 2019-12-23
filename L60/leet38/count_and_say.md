[LeetCode] 38. Count and Say 计数和读法 

 
The count-and-say sequence is the sequence of integers with the first five terms as following:
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.
Note: Each term of the sequence of integers will be represented as a string.
 
Example 1:
Input: 1
Output: "1"
Example 2:
Input: 4
Output: "1211" 
 
这道计数和读法问题还是第一次遇到，看似挺复杂，其实仔细一看，算法很简单，就是对于前一个数，找出相同元素的个数，把个数和该元素存到新的 string 里。代码如下：
 

class Solution {
public:
    string countAndSay(int n) {
        if (n <= 0) return "";
        string res = "1";
        while (--n) {
            string cur = "";
            for (int i = 0; i < res.size(); ++i) {
                int cnt = 1;
                while (i + 1 < res.size() && res[i] == res[i + 1]) {
                    ++cnt;
                    ++i;
                }
                cur += to_string(cnt) + res[i];
            }
            res = cur;
        }
        return res;
    }
};

 
博主出于好奇打印出了前 12 个数字，发现一个很有意思的现象，不管打印到后面多少位，出现的数字只是由 1, 2 和3 组成，网上也有人发现了并分析了原因，参见这个帖子，前十二个数字如下：
 

1
1 1
2 1
1 2 1 1
1 1 1 2 2 1
3 1 2 2 1 1
1 3 1 1 2 2 2 1
1 1 1 3 2 1 3 2 1 1
3 1 1 3 1 2 1 1 1 3 1 2 2 1
1 3 2 1 1 3 1 1 1 2 3 1 1 3 1 1 2 2 1 1
1 1 1 3 1 2 2 1 1 3 3 1 1 2 1 3 2 1 1 3 2 1 2 2 2 1
3 1 1 3 1 1 2 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 3 2 1 1