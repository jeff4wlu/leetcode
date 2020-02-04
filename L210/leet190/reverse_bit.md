[LeetCode] 190. Reverse Bits 颠倒二进制位 

 
Reverse bits of a given 32 bits unsigned integer.
 
Example 1:
Input: 00000010100101000001111010011100
Output: 00111001011110000010100101000000
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.
Example 2:
Input: 11111111111111111111111111111101
Output: 10111111111111111111111111111111
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10101111110010110010011101101001.
 
Note:
Note that in some languages such as Java, there is no unsigned integer type. In this case, both input and output will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above the input represents the signed integer -3 and the output represents the signed integer -1073741825.
 
Follow up:
If this function is called many times, how would you optimize it?
 
Credits:
Special thanks to @ts for adding this problem and creating all test cases.
 
这道题又是在考察位操作 Bit Operation，LeetCode 中有关位操作的题也有不少，比如 Repeated DNA Sequences，Single Number,   Single Number II ，和 Grey Code 等等。跟上面那些题比起来，这道题简直不能再简单了，我们只需要把要翻转的数从右向左一位位的取出来，如果取出来的是1，将结果 res 左移一位并且加上1；如果取出来的是0，将结果 res 左移一位，然后将n右移一位即可，参见代码如下：
 
解法一：

class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t res = 0;
        for (int i = 0; i < 32; ++i) {
            if (n & 1 == 1) {
                res = (res << 1) + 1;
            } else {
                res = res << 1;
            }
            n = n >> 1;
        }
        return res;
    }
};

 
我们可以简化上面的代码，去掉 if...else... 结构，可以结果 res 左移一位，然后再判断n的最低位是否为1，是的话那么结果 res 加上1，然后将n右移一位即可，代码如下：
 
解法二：

class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t res = 0;
        for (int i = 0; i < 32; ++i) {
            res <<= 1;
            if ((n & 1) == 1) ++res;
            n >>= 1;
        }
        return res;
    }
};

 
我们继续简化上面的解法，将 if 判断句直接揉进去，通过 ‘或’ 上一个n的最低位即可，用n ‘与’ 1提取最低位，然后将n右移一位即可，代码如下：
 
解法三：

class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t res = 0;
        for (int i = 0; i < 32; ++i) {
            res = (res << 1) | (n & 1);
            n >>= 1;
        }
        return res;
    }
};

 
博主还能进一步简化，这里不更新n的值，而是直接将n右移i位，然后通过 ‘与’ 1来提取出该位，加到左移一位后的结果 res 中即可，参加代码如下：
 
解法四：

class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t res = 0;
        for (int i = 0; i < 32; ++i) {
            res = (res << 1) + (n >> i & 1);
        }
        return res;
    }
};

 
我们也可以换一种角度来做，首先将n右移i位，然后通过 ‘与’ 1来提取出该位，然后将其左移 (32 - i) 位，然后 ‘或’ 上结果 res，就是其颠倒后应该在的位置，参见代码如下： 
 
解法五：

class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t res = 0;
        for (int i = 0; i < 32; ++i) {
            res |= ((n >> i) & 1) << (31 - i);
        }
        return res;
    }
};

 
讨论：这道题的最高票解法实在是很叼啊，参见这个帖子，但是博主没有太理解啊，希望哪位大神能讲解一下哈～
 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/190
 
类似题目：
Number of 1 Bits
Reverse Integer