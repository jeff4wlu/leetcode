[LeetCode] 153. Find Minimum in Rotated Sorted Array 寻找旋转有序数组的最小值 

 
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
Find the minimum element.
You may assume no duplicate exists in the array.
Example 1:
Input: [3,4,5,1,2] 
Output: 1
Example 2:
Input: [4,5,6,7,0,1,2]
Output: 0
 
这道寻找旋转有序数组的最小值肯定不能通过直接遍历整个数组来寻找，这个方法过于简单粗暴，这样的话，旋不旋转就没有意义。应该考虑将时间复杂度从简单粗暴的 O(n) 缩小到 O(lgn)，这时候二分查找法就浮现在脑海。这里的二分法属于博主之前的总结帖 LeetCode Binary Search Summary 二分搜索法小结 中的第五类，也是比较难的那一类，没有固定的 target 值比较，而是要跟数组中某个特定位置上的数字比较，决定接下来去哪一边继续搜索。这里用中间的值 nums[mid] 和右边界值 nums[right] 进行比较，若数组没有旋转或者旋转点在左半段的时候，中间值是一定小于右边界值的，所以要去左半边继续搜索，反之则去右半段查找，最终返回 nums[right] 即可，参见代码如下：
 
解法一：

class Solution {
public:
    int findMin(vector<int>& nums) {
        int left = 0, right = (int)nums.size() - 1;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > nums[right]) left = mid + 1;
            else right = mid;
        }
        return nums[right];
    }
};

 
下面这种分治法 Divide and Conquer 的解法，由热心网友 howard144 提供，这里每次将区间 [start, end] 从中间 mid 位置分为两段，分别调用递归函数，并比较返回值，每次取返回值较小的那个即可，参见代码如下：
 
解法二：

class Solution {
public:
    int findMin(vector<int>& nums) {
        return helper(nums, 0, (int)nums.size() - 1);
    }
    int helper(vector<int>& nums, int start, int end) {
        if (nums[start] <= nums[end]) return nums[start];
        int mid = (start + end) / 2;
        return min(helper(nums, start, mid), helper(nums, mid + 1, end));
    }
};

 
讨论：对于数组中有重复数字的情况，请参见博主的另一篇博文 Find Minimum in Rotated Sorted Array II。
 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/153
 
类似题目：
Search in Rotated Sorted Array
Find Minimum in Rotated Sorted Array II