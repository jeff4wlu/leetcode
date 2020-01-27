[LeetCode] 152. Maximum Product Subarray 求最大子数组乘积 

 
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.
Example 1:
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 
这个求最大子数组乘积问题是由最大子数组之和 Maximum Subarray 演变而来，但是却比求最大子数组之和要复杂，因为在求和的时候，遇到0，不会改变最大值，遇到负数，也只是会减小最大值而已。而在求最大子数组乘积的问题中，遇到0会使整个乘积为0，而遇到负数，则会使最大乘积变成最小乘积，正因为有负数和0的存在，使问题变得复杂了不少。比如，现在有一个数组 [2, 3, -2, 4]，可以很容易的找出所有的连续子数组，[2]，[3]，[-2]，[4]，[2, 3]，[3, -2]，[-2, 4]，[2, 3, -2]，[3, -2, 4]，[2, 3, -2, 4]，然后可以很轻松的算出最大的子数组乘积为6，来自子数组 [2, 3]。但如何写代码来实现自动找出最大子数组乘积呢，博主最先想到的方比较简单粗暴，就是找出所有的子数组，然后算出每一个子数组的乘积，然后比较找出最大的一个，需要两个 for 循环，第一个 for 遍历整个数组，第二个 for 遍历含有当前数字的子数组，就是按以下顺序找出子数组: [2]，[2, 3]，[2, 3, -2]，[2, 3, -2, 4]，[3]，[3, -2]，[3, -2, 4]，[-2]，[-2, 4]，[4]，在本地测试的一些数组全部通过，于是兴高采烈的拿到 OJ 上测试，结果丧心病狂的 OJ 用一个有 15000 个数字的数组来测试，然后说程序的运行时间超过了要求值，一看代码，果然如此，时间复杂度 O(n2), 得想办法只用一次循环搞定。想来想去想不出好方法，于是到网上搜各位大神的解决方法。其实这道题最直接的方法就是用 DP 来做，而且要用两个 dp 数组，其中 f[i] 表示子数组 [0, i] 范围内并且一定包含 nums[i] 数字的最大子数组乘积，g[i] 表示子数组 [0, i] 范围内并且一定包含 nums[i] 数字的最小子数组乘积，初始化时 f[0] 和 g[0] 都初始化为 nums[0]，其余都初始化为0。那么从数组的第二个数字开始遍历，那么此时的最大值和最小值只会在这三个数字之间产生，即 f[i-1]*nums[i]，g[i-1]*nums[i]，和 nums[i]。所以用三者中的最大值来更新 f[i]，用最小值来更新 g[i]，然后用 f[i] 来更新结果 res 即可，由于最终的结果不一定会包括 nums[n-1] 这个数字，所以 f[n-1] 不一定是最终解，不断更新的结果 res 才是，参见代码如下：
 
解法一：

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int res = nums[0], n = nums.size();
        vector<int> f(n, 0), g(n, 0);
        f[0] = nums[0];
        g[0] = nums[0];
        for (int i = 1; i < n; ++i) {
            f[i] = max(max(f[i - 1] * nums[i], g[i - 1] * nums[i]), nums[i]);
            g[i] = min(min(f[i - 1] * nums[i], g[i - 1] * nums[i]), nums[i]);
            res = max(res, f[i]);
        }
        return res;
    }
};

 
我们可以对上面的解法进行空间上的优化，以下摘自 OJ 官方解答，大体思路相同，写法更加简洁：
Besides keeping track of the largest product, we also need to keep track of the smallest product. Why? The smallest product, which is the largest in the negative sense could become the maximum when being multiplied by a negative number.
Let us denote that:
f(k) = Largest product subarray, from index 0 up to k.
 
Similarly,
g(k) = Smallest product subarray, from index 0 up to k.
 
Then,
f(k) = max( f(k-1) * A[k], A[k], g(k-1) * A[k] )
g(k) = min( g(k-1) * A[k], A[k], f(k-1) * A[k] )
 
There we have a dynamic programming formula. Using two arrays of size n, we could deduce the final answer as f(n-1). Since we only need to access its previous elements at each step, two variables are sufficient.

public int maxProduct(int[] A) {
   assert A.length > 0;
   int max = A[0], min = A[0], maxAns = A[0];
   for (int i = 1; i < A.length; i++) {
      int mx = max, mn = min;
      max = Math.max(Math.max(A[i], mx * A[i]), mn * A[i]);
      min = Math.min(Math.min(A[i], mx * A[i]), mn * A[i]);
      maxAns = Math.max(max, maxAns);
   }
   return maxAns;
}

 
根据上述描述可以写出代码如下：
 
解法二：

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        if (nums.empty()) return 0;
        int res = nums[0], mn = nums[0], mx = nums[0];
        for (int i = 1; i < nums.size(); ++i) {
            int tmax = mx, tmin = mn;
            mx = max(max(nums[i], tmax * nums[i]), tmin * nums[i]);
            mn = min(min(nums[i], tmax * nums[i]), tmin * nums[i]);
            res = max(res, mx);
        }
        return res;
    }
};

 
下面这种方法也是用两个变量来表示当前最大值和最小值的，但是没有无脑比较三个数，而是对于当前的 nums[i] 值进行了正负情况的讨论：
1. 当遍历到一个正数时，此时的最大值等于之前的最大值乘以这个正数和当前正数中的较大值，此时的最小值等于之前的最小值乘以这个正数和当前正数中的较小值。
2. 当遍历到一个负数时，先用一个变量t保存之前的最大值 mx，然后此时的最大值等于之前最小值乘以这个负数和当前负数中的较大值，此时的最小值等于之前保存的最大值t乘以这个负数和当前负数中的较小值。
3. 在每遍历完一个数时，都要更新最终的最大值。
P.S. 如果这里改成求最小值的话，就是求最小子数组乘积，并且时间复杂度是醉人的 O(n)，是不是很强大呢，参见代码如下：
 
解法三：

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int res = nums[0], mx = res, mn = res;
        for (int i = 1; i < nums.size(); ++i) {
            if (nums[i] > 0) {
                mx = max(mx * nums[i], nums[i]);
                mn = min(mn * nums[i], nums[i]);
            } else {
                int t = mx;
                mx = max(mn * nums[i], nums[i]);
                mn = min(t * nums[i], nums[i]);
            }
            res = max(res, mx);
        }
        return res;
    }
};

 
下面这道题使用了一个 trick 来将上面解法的分情况讨论合成了一种，在上面的解法中分析了当 nums[i] 为正数时，最大值和最小值的更新情况，为负数时，稍有不同的就是最小值更新时要用到之前的最大值，而不是更新后的最大值，所以才要用变量t来保存之前的结果。而下面这种方法的巧妙处在于先判断一个当前数字是否是负数，是的话就交换最大值和最小值。那么此时的 mx 就是之前的 mn，所以 mx 的更新还是跟上面的方法是统一的，而在在更新 mn 的时候，之前的 mx 已经保存到 mn 中了，而且并没有改变，所以可以直接拿来用，不得不说，确实叼啊，参见代码如下：
 
解法四：

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int res = nums[0], mx = res, mn = res;
        for (int i = 1; i < nums.size(); ++i) {
            if (nums[i] < 0) swap(mx, mn);
            mx = max(nums[i], mx * nums[i]);
            mn = min(nums[i], mn * nums[i]);
            res = max(res, mx);
        }
        return res;
    }
};

 
再来看一种画风不太一样的解法，这种解法遍历了两次，一次是正向遍历，一次是反向遍历，相当于正向建立一个累加积数组，每次用出现的最大值更新结果 res，然后再反响建立一个累加积数组，再用出现的最大值更新结果 res，注意当遇到0的时候，prod 要重置为1。至于为啥正反两次遍历就可以得到正确的结果了呢？主要还是由于负数个数的关系，因为负数可能会把最大值和最小值翻转，那么当有奇数个负数时，如果只是正向遍历的话，可能会出错，比如 [-1, -2, -3]，累加积会得到 -1，2，-6，看起来最大值只能为2，其实不对，而如果我们再反向来一遍，累加积为 -3，6，-6，就可以得到6了。所以当负数个数为奇数时，首次出现和末尾出现的负数就很重要，有可能会是最大积的组成数字，所以遍历两次就不会漏掉组成最大值的机会，参见代码如下：
 
解法五：

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int res = nums[0], prod = 1, n = nums.size();
        for (int i = 0; i < n; ++i) {
            res = max(res, prod *= nums[i]);
            if (nums[i] == 0) prod = 1;
        }
        prod = 1;
        for (int i = n - 1; i >= 0; --i) {
            res = max(res, prod *= nums[i]);
            if (nums[i] == 0) prod = 1;
        }
        return res;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/152
 
类似题目：
Maximum Subarray
House Robber
Product of Array Except Self
Maximum Product of Three Numbers
Subarray Product Less Than K