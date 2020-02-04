[LeetCode] 198. House Robber 打家劫舍 

 
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
Example 1:
Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:
Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
Credits:
Special thanks to @ifanchu for adding this problem and creating all test cases. Also thanks to @ts for adding additional test cases.
 
这道题的本质相当于在一列数组中取出一个或多个不相邻数，使其和最大。那么对于这类求极值的问题首先考虑动态规划 Dynamic Programming 来解，维护一个一位数组 dp，其中 dp[i] 表示 [0, i] 区间可以抢夺的最大值，对当前i来说，有抢和不抢两种互斥的选择，不抢即为 dp[i-1]（等价于去掉 nums[i] 只抢 [0, i-1] 区间最大值），抢即为 dp[i-2] + nums[i]（等价于去掉 nums[i-1]）。再举一个简单的例子来说明一下吧，比如说 nums为{3, 2, 1, 5}，那么来看 dp 数组应该是什么样的，首先 dp[0]=3 没啥疑问，再看 dp[1] 是多少呢，由于3比2大，所以抢第一个房子的3，当前房子的2不抢，则dp[1]=3，那么再来看 dp[2]，由于不能抢相邻的，所以可以用再前面的一个的 dp 值加上当前的房间值，和当前房间的前面一个 dp 值比较，取较大值当做当前 dp 值，这样就可以得到状态转移方程 dp[i] = max(num[i] + dp[i - 2], dp[i - 1]), 且需要初始化 dp[0] 和 dp[1]，其中 dp[0] 即为 num[0]，dp[1] 此时应该为 max(num[0], num[1])，代码如下：
 
解法一：

class Solution {
public:
    int rob(vector<int>& nums) {
        if (nums.size() <= 1) return nums.empty() ? 0 : nums[0];
        vector<int> dp = {nums[0], max(nums[0], nums[1])};
        for (int i = 2; i < nums.size(); ++i) {
            dp.push_back(max(nums[i] + dp[i - 2], dp[i - 1]));
        }
        return dp.back();
    }
};

 
还有一种解法，核心思想还是用 DP，分别维护两个变量 robEven 和 robOdd，顾名思义，robEven 就是要抢偶数位置的房子，robOdd 就是要抢奇数位置的房子。所以在遍历房子数组时，如果是偶数位置，那么 robEven 就要加上当前数字，然后和 robOdd 比较，取较大的来更新 robEven。这里就看出来了，robEven 组成的值并不是只由偶数位置的数字，只是当前要抢偶数位置而已。同理，当奇数位置时，robOdd 加上当前数字和 robEven 比较，取较大值来更新 robOdd，这种按奇偶分别来更新的方法，可以保证组成最大和的数字不相邻，最后别忘了在 robEven 和 robOdd 种取较大值返回，代码如下：
 
解法二：

class Solution {
public:
    int rob(vector<int>& nums) {
        int robEven = 0, robOdd = 0, n = nums.size();
        for (int i = 0; i < n; ++i) {
            if (i % 2 == 0) {
                robEven = max(robEven + nums[i], robOdd);
            } else {
                robOdd = max(robEven, robOdd + nums[i]);
            }
        }
        return max(robEven, robOdd);
    }
};

 
上述方法还可以进一步简洁，我们使用两个变量 rob 和 notRob，其中 rob 表示抢当前的房子，notRob 表示不抢当前的房子，那么在遍历的过程中，先用两个变量 preRob 和 preNotRob 来分别记录更新之前的值，由于 rob 是要抢当前的房子，那么前一个房子一定不能抢，所以使用 preNotRob 加上当前的数字赋给 rob，然后 notRob 表示不能抢当前的房子，那么之前的房子就可以抢也可以不抢，所以将 preRob 和 preNotRob 中的较大值赋给 notRob，参见代码如下：
 
解法三：

class Solution {
public:
    int rob(vector<int>& nums) {
        int rob = 0, notRob = 0, n = nums.size();
        for (int i = 0; i < n; ++i) {
            int preRob = rob, preNotRob = notRob;
            rob = preNotRob + nums[i];
            notRob = max(preRob, preNotRob);
        }
        return max(rob, notRob);
    }
};