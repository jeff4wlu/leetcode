[LeetCode] 172. Factorial Trailing Zeroes 求阶乘末尾零的个数 

 
Given an integer n, return the number of trailing zeroes in n!.
Example 1:
Input: 3
Output: 0
Explanation: 3! = 6, no trailing zero.
Example 2:
Input: 5
Output: 1
Explanation: 5! = 120, one trailing zero.
Note: Your solution should be in logarithmic time complexity.
Credits:
Special thanks to @ts for adding this problem and creating all test cases.
 
这道题并没有什么难度，是让求一个数的阶乘末尾0的个数，也就是要找乘数中 10 的个数，而 10 可分解为2和5，而2的数量又远大于5的数量（比如1到 10 中有2个5，5个2），那么此题即便为找出5的个数。仍需注意的一点就是，像 25，125，这样的不只含有一个5的数字需要考虑进去，参加代码如下：
 
C++ 解法一：

class Solution {
public:
    int trailingZeroes(int n) {
        int res = 0;
        while (n) {
            res += n / 5;
            n /= 5;
        }
        return res;
    }
};

 
Java 解法一：

public class Solution {
    public int trailingZeroes(int n) {
        int res = 0;
        while (n > 0) {
            res += n / 5;
            n /= 5;
        }
        return res;
    }
}

 
这题还有递归的解法，思路和上面完全一样，写法更简洁了，一行搞定碉堡了。
 
C++ 解法二：
class Solution {
public:
    int trailingZeroes(int n) {
        return n == 0 ? 0 : n / 5 + trailingZeroes(n / 5);
    }
};
 
Java 解法二：
public class Solution {
    public int trailingZeroes(int n) {
        return n == 0 ? 0 : n / 5 + trailingZeroes(n / 5);
    }
}
 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/172
 
类似题目：
Number of Digit One
Preimage Size of Factorial Zeroes Function    