[LeetCode] 209. Minimum Size Subarray Sum 最短子数组之和 

 
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.
Example: 
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
Follow up:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).  
Credits:
Special thanks to @Freezen for adding this problem and creating all test cases.
 
这道题给定了我们一个数字，让求子数组之和大于等于给定值的最小长度，注意这里是大于等于，不是等于。跟之前那道 Maximum Subarray 有些类似，并且题目中要求实现 O(n) 和 O(nlgn) 两种解法，那么先来看 O(n) 的解法，需要定义两个指针 left 和 right，分别记录子数组的左右的边界位置，然后让 right 向右移，直到子数组和大于等于给定值或者 right 达到数组末尾，此时更新最短距离，并且将 left 像右移一位，然后再 sum 中减去移去的值，然后重复上面的步骤，直到 right 到达末尾，且 left 到达临界位置，即要么到达边界，要么再往右移动，和就会小于给定值。代码如下：
 
解法一

// O(n)
class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        if (nums.empty()) return 0;
        int left = 0, right = 0, sum = 0, len = nums.size(), res = len + 1;
        while (right < len) {
            while (sum < s && right < len) {
                sum += nums[right++];
            }
            while (sum >= s) {
                res = min(res, right - left);
                sum -= nums[left++];
            }
        }
        return res == len + 1 ? 0 : res;
    }
};

 
同样的思路，我们也可以换一种写法，参考代码如下：
 
解法二：

class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int res = INT_MAX, left = 0, sum = 0;
        for (int i = 0; i < nums.size(); ++i) {
            sum += nums[i];
            while (left <= i && sum >= s) {
                res = min(res, i - left + 1);
                sum -= nums[left++];
            }
        }
        return res == INT_MAX ? 0 : res;
    }
};

 
下面再来看看 O(nlgn) 的解法，这个解法要用到二分查找法，思路是，建立一个比原数组长一位的 sums 数组，其中 sums[i] 表示 nums 数组中 [0, i - 1] 的和，然后对于 sums 中每一个值 sums[i]，用二分查找法找到子数组的右边界位置，使该子数组之和大于 sums[i] + s，然后更新最短长度的距离即可。代码如下：
 
解法三：

// O(nlgn)
class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int len = nums.size(), sums[len + 1] = {0}, res = len + 1;
        for (int i = 1; i < len + 1; ++i) sums[i] = sums[i - 1] + nums[i - 1];
        for (int i = 0; i < len + 1; ++i) {
            int right = searchRight(i + 1, len, sums[i] + s, sums);
            if (right == len + 1) break;
            if (res > right - i) res = right - i;
        }
        return res == len + 1 ? 0 : res;
    }
    int searchRight(int left, int right, int key, int sums[]) {
        while (left <= right) {
            int mid = (left + right) / 2;
            if (sums[mid] >= key) right = mid - 1;
            else left = mid + 1;
        }
        return left;
    }
};

 
我们也可以不用为二分查找法专门写一个函数，直接嵌套在 for 循环中即可，参加代码如下：
 
解法四：

class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int res = INT_MAX, n = nums.size();
        vector<int> sums(n + 1, 0);
        for (int i = 1; i < n + 1; ++i) sums[i] = sums[i - 1] + nums[i - 1];
        for (int i = 0; i < n; ++i) {
            int left = i + 1, right = n, t = sums[i] + s;
            while (left <= right) {
                int mid = left + (right - left) / 2;
                if (sums[mid] < t) left = mid + 1;
                else right = mid - 1;
            }
            if (left == n + 1) break;
            res = min(res, left - i);
        }
        return res == INT_MAX ? 0 : res;
    }
};

 
讨论：本题有一个很好的 Follow up，就是去掉所有数字是正数的限制条件，而去掉这个条件会使得累加数组不一定会是递增的了，那么就不能使用二分法，同时双指针的方法也会失效，只能另辟蹊径了。其实博主觉得同时应该去掉大于s的条件，只保留 sum=s 这个要求，因为这样就可以在建立累加数组后用 2sum 的思路，快速查找 s-sum 是否存在，如果有了大于的条件，还得继续遍历所有大于 s-sum 的值，效率提高不了多少。
 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/209
 
类似题目：
Minimum Window Substring
Subarray Sum Equals K
Maximum Length of Repeated Subarray