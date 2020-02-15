[LeetCode] Single Number III 单独的数字之三 

 
Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.
Example:
Input:  [1,2,1,3,2,5]
Output: [3,5]
Note:
The order of the result is not important. So in the above example, [5, 3] is also correct.
Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?
 
这道题是之前那两道 Single Number 和 Single Number II 的再次延伸，说实话，这类位操作 Bit Manipulation 的题，如果之前没有遇到过类似的题目，楞想是很难相出来的，于是我只能上网搜大神们的解法，发现还真是巧妙啊。这道题其实是很巧妙的利用了 Single Number 的解法，因为那道解法是可以准确的找出只出现了一次的数字，但前提是其他数字必须出现两次才行。而这题有两个数字都只出现了一次，那么我们如果能想办法把原数组分为两个小数组，不相同的两个数字分别在两个小数组中，这样分别调用 Single Number 的解法就可以得到答案。那么如何实现呢，首先我们先把原数组全部异或起来，那么我们会得到一个数字，这个数字是两个不相同的数字异或的结果，我们取出其中任意一位为 ‘1’ 的位，为了方便起见，我们用 a &= -a 来取出最右端为 ‘1’ 的位，具体来说下这个是如何操作的吧。就拿题目中的例子来说，如果我们将其全部 '异或' 起来，我们知道相同的两个数 '异或' 的话为0，那么两个1，两个2，都抵消了，就剩3和5 '异或' 起来，那么就是二进制的 11 和 101 '异或' ，得到110。然后我们进行 a &= -a 操作。首先变负数吧，在二进制中负数采用补码的形式，而补码就是反码 +1，那么 110 的反码是 11...1001，那么加1后是 11...1010，然后和 110 相与，得到了 10，就是代码中的 diff 变量。得到了这个 diff，就可以将原数组分为两个数组了。为啥呢，我们想阿，如果两个相同的数字 '异或' ，每位都会是0，而不同的数字 '异或' ，一定会有对应位不同，一个0一个1，这样 '异或' 是1。比如3和5的二进制 11 和 101，如果从低往高看，最开始产生不同的就是第二位，那么我们用第二位来和数组中每个数字相与，根据结果的不同，一定可以把3和5区分开来，而其他的数字由于是成对出现，所以区分开来也是成对的，最终都会 '异或' 成0，不会3和5产生影响。分别将两个小组中的数字都异或起来，就可以得到最终结果了，参见代码如下：
 

class Solution {
public:
    vector<int> singleNumber(vector<int>& nums) {
        int diff = accumulate(nums.begin(), nums.end(), 0, bit_xor<int>());
        diff &= -diff;
        vector<int> res(2, 0);
        for (auto &a : nums) {
            if (a & diff) res[0] ^= a;
            else res[1] ^= a;
        }
        return res;
    }
};

 
类似题目：
Single Number II
Single Number