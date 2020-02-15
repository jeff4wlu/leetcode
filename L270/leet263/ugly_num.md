[LeetCode] 263. Ugly Number 丑陋数 

 
Write a program to check whether a given number is an ugly number.
Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
Example 1:
Input: 6
Output: true
Explanation: 6 = 2 × 3
Example 2:
Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2
Example 3:
Input: 14
Output: false 
Explanation: 14 is not ugly since it includes another prime factor 7.
Note:
1 is typically treated as an ugly number.
Input is within the 32-bit signed integer range: [−231,  231 − 1].
 
这道题让我们检测一个数是否为丑陋数，所谓丑陋数就是其质数因子只能是 2，3，5。那么最直接的办法就是不停的除以这些质数，如果剩余的数字是1的话就是丑陋数了，参见代码如下：
 
解法一：

class Solution {
public:
    bool isUgly(int num) {
        while (num >= 2) {
            if (num % 2 == 0) num /= 2;
            else if (num % 3 == 0) num /= 3;
            else if (num % 5 == 0) num /= 5;
            else return false;
        }
        return num == 1;
    }
};

 
我们也可以换一种写法，分别不停的除以 2，3，5，并且看最后剩下来的数字是否为1即可，参见代码如下：
 
解法二：

class Solution {
public:
    bool isUgly(int num) {
        if (num <= 0) return false;
        while (num % 2 == 0) num /= 2;
        while (num % 3 == 0) num /= 3;
        while (num % 5 == 0) num /= 5;
        return num == 1;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/263
 
类似题目：
Super Ugly Number
Ugly Number II
Happy Number
Count Primes