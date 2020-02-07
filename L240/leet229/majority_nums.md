[LeetCode] 229. Majority Element II 求大多数之二 

 
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
Note: The algorithm should run in linear time and in O(1) space.
Example 1:
Input: [3,2,3]
Output: [3]
Example 2:
Input: [1,1,1,3,3,2,2,2]
Output: [1,2]
 
这道题让我们求出现次数大于 n/3 的数字，而且限定了时间和空间复杂度，那么就不能排序，也不能使用 HashMap，这么苛刻的限制条件只有一种方法能解了，那就是摩尔投票法 Moore Voting，这种方法在之前那道题 Majority Element 中也使用了。题目中给了一条很重要的提示，让先考虑可能会有多少个这样的数字，经过举了很多例子分析得出，任意一个数组出现次数大于 n/3 的数最多有两个，具体的证明博主就不会了，博主也不是数学专业的（热心网友用手走路提供了证明：如果有超过两个，也就是至少三个数字满足“出现的次数大于 n/3”，那么就意味着数组里总共有超过 3*(n/3) = n 个数字，这与已知的数组大小矛盾，所以，只可能有两个或者更少）。那么有了这个信息，使用投票法的核心是找出两个候选数进行投票，需要两遍遍历，第一遍历找出两个候选数，第二遍遍历重新投票验证这两个候选数是否为符合题意的数即可，选候选数方法和前面那篇 Majority Element 一样，由于之前那题题目中限定了一定会有大多数存在，故而省略了验证候选众数的步骤，这道题却没有这种限定，即满足要求的大多数可能不存在，所以要有验证，参加代码如下：
 

class Solution {
public:
    vector<int> majorityElement(vector<int>& nums) {
        vector<int> res;
        int a = 0, b = 0, cnt1 = 0, cnt2 = 0, n = nums.size();
        for (int num : nums) {
            if (num == a) ++cnt1;
            else if (num == b) ++cnt2;
            else if (cnt1 == 0) { a = num; cnt1 = 1; }
            else if (cnt2 == 0) { b = num; cnt2 = 1; }
            else { --cnt1; --cnt2; }
        }
        cnt1 = cnt2 = 0;
        for (int num : nums) {
            if (num == a) ++cnt1;
            else if (num == b) ++cnt2;
        }
        if (cnt1 > n / 3) res.push_back(a);
        if (cnt2 > n / 3) res.push_back(b);
        return res;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/229
 
类似题目：
Majority Element
Check If a Number Is Majority Element in a Sorted Array