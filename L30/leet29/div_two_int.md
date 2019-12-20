[LeetCode] 29. Divide Two Integers 两数相除 

 
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
Return the quotient after dividing dividend by divisor.
The integer division should truncate toward zero.
Example 1:
Input: dividend = 10, divisor = 3
Output: 3
Example 2:
Input: dividend = 7, divisor = -3
Output: -2
Note:
Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.


思路：一个数除以被除数的本质就是包含多少个被除数．因此我们可以利用乘或位运算将被除数不断变大，直至找到看将被除数放大多少倍可以得到除数．举个栗子：100 / 5

１． 将５乘以２，得到10，而100 > 10，说明100包含更多的5

２．将10乘以２，得到20，100 > 20，因此还可以继续乘

３．将20乘以２，得到40，100 > 40

４．将40乘以２，得到80，100 > 40

５．将80乘以２，得到160，100 < 160

在此之前，我们对除数的放大倍数是以指数方式的，这样可以快速的找到一个其商的一个上界，然后将100－80剩下的作为被除数继续做如上操作．

我们也可以不用指数增长的方式来做，也就是一次增加一个被除数，最终也可以得到结果，但是这种方式的时间复杂度是O(n)，而指数增长被除数的时间复杂度是O(log n)，其效率相差极大．

还有一个会产生溢出情况的是，当除数是INT_MIN, 而被除数是-1的时候，因为负数的最小值比整数的最大值绝对值大１，因此需要注意这种情况．

代码如下：

class Solution {
public:
    int divide(int dividend, int divisor) {
        if(dividend==INT_MIN && divisor ==-1) return INT_MAX;
        if(dividend ==0) return 0;
        long ans =0, nums1=labs(dividend), nums2=labs(divisor);
        while(nums2 <= nums1)
        {
            int cnt = 1;
            while((nums2 << cnt) <= nums1) cnt++;
            ans += (1<<(cnt-1)), nums1 -= (nums2<<(cnt-1));
        }
        return (dividend>0 ^ divisor>0)?-ans:ans;
    }
};

