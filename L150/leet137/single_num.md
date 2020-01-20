[LeetCode] 137. Single Number II 单独的数字之二 

 
Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.
Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
Example 1:
Input: [2,2,3,2]
Output: 3
Example 2:
Input: [0,1,0,1,0,1,99]
Output: 99
 
这道题是之前那道 Single Number 的延伸，那道题的解法就比较独特，是利用计算机按位储存数字的特性来做的，这道题就是除了一个单独的数字之外，数组中其他的数字都出现了三次，还是要利用位操作 Bit Manipulation 来解。可以建立一个 32 位的数字，来统计每一位上1出现的个数，如果某一位上为1的话，那么如果该整数出现了三次，对3取余为0，这样把每个数的对应位都加起来对3取余，最终剩下来的那个数就是单独的数字。代码如下：
 
解法一：

class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int res = 0;
        for (int i = 0; i < 32; ++i) {
            int sum = 0;
            for (int j = 0; j < nums.size(); ++j) {
                sum += (nums[j] >> i) & 1;
            }
            res |= (sum % 3) << i;
        }
        return res;
    }
};

 
还有一种解法，思路很相似，用3个整数来表示 INT 的各位的出现次数情况，one 表示出现了1次，two 表示出现了2次。当出现3次的时候该位清零。最后答案就是one的值。
ones   代表第 ith 位只出现一次的掩码变量
twos  代表第 ith 位只出现两次次的掩码变量
threes  代表第 ith 位只出现三次的掩码变量
假设现在有一个数字1，更新 one 的方法就是 ‘亦或’ 这个1，则 one 就变成了1，而 two 的更新方法是用上一个状态下的 one 去 ‘与’ 上数字1，然后 ‘或’ 上这个结果，这样假如之前 one 是1，那么此时 two 也会变成1，这 make sense，因为说明是当前位遇到两个1了；反之如果之前 one 是0，那么现在 two 也就是0。注意更新的顺序是先更新 two，再更新 one，不理解的话只要带个只有一个数字1的输入数组看一下就不难理解了。然后更新 three，如果此时 one 和 two 都是1了，由于先更新的 two，再更新的 one，two 为1，说明此时至少有两个数字1了，而此时 one 为1，说明了此时已经有了三个数字1，这块要仔细想清楚，因为 one 是要 ‘亦或’ 一个1的，值能为1，说明之前 one 为0，实际情况是，当第二个1来的时候，two 先更新为1，此时 one 再更新为0，下面 three 就是0了，那么 ‘与’ 上t hree 的相反数1不会改变 one 和 two 的值；那么当第三个1来的时候，two 还是1，此时 one 就更新为1了，那么 three 就更新为1了，此时就要清空 one 和 two 了，让它们 ‘与’ 上 three 的相反数0即可，最终结果将会保存在 one 中，参见代码如下：
 
解法二：

class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int one = 0, two = 0, three = 0;
        for (int i = 0; i < nums.size(); ++i) {
            two |= one & nums[i];
            one ^= nums[i];
            three = one & two;
            one &= ~three;
            two &= ~three;
        }
        return one;
    }
};

 
下面这种解法思路也十分巧妙，根据上面解法的思路，我们把数组中数字的每一位累加起来对3取余，剩下的结果就是那个单独数组该位上的数字，由于累加的过程都要对3取余，那么每一位上累加的过程就是 0->1->2->0，换成二进制的表示为 00->01->10->00，可以写出对应关系：
00 (+) 1 = 01
01 (+) 1 = 10
10 (+) 1 = 00 ( mod 3)
用 ab 来表示开始的状态，对于加1操作后，得到的新状态的 ab 的算法如下：
b = b xor r & ~a;
a = a xor r & ~b;
这里的 ab 就是上面的三种状态 00，01，10 的十位和各位，刚开始的时候，a和b都是0，当此时遇到数字1的时候，b更新为1，a更新为0，就是 01 的状态；再次遇到1的时候，b更新为0，a更新为1，就是 10 的状态；再次遇到1的时候，b更新为0，a更新为0，就是 00 的状态，相当于重置了；最后的结果保存在b中，明白了上面的分析过程，就能写出代码如下；
 
解法三：

class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int a = 0, b = 0;
        for (int i = 0; i < nums.size(); ++i) {
            b = (b ^ nums[i]) & ~a;
            a = (a ^ nums[i]) & ~b;
        }
        return b;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/137
 
类似题目：
Single Number
Single Number III