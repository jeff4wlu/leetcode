[LeetCode] 33. Search in Rotated Sorted Array 在旋转有序数组中搜索 

 
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
You are given a target value to search. If found in the array return its index, otherwise return -1.
You may assume no duplicate exists in the array.
Your algorithm's runtime complexity must be in the order of O(log n).
Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
 
这道题让在旋转数组中搜索一个给定值，若存在返回坐标，若不存在返回 -1。我们还是考虑二分搜索法，但是这道题的难点在于不知道原数组在哪旋转了，还是用题目中给的例子来分析，对于数组 [0 1 2 4 5 6 7] 共有下列七种旋转方法（红色表示中点之前或者之后一定为有序的）：
0　　1　　2　　4　　5　　6　　7
7　　0　　1　　2　　4　　5　　6
6　　7　　0　　1　　2　　4　　5
5　　6　　7　　0　　1　　2　　4
4　　5　　6　　7　　0　　1　　2
2　　4　　5　　6　　7　　0　　1
1　　2　　4　　5　　6　　7　　0
二分搜索法的关键在于获得了中间数后，判断下面要搜索左半段还是右半段，观察上面红色的数字都是升序的（Github 中可能无法显示颜色，参见博客园上的帖子），可以得出出规律，如果中间的数小于最右边的数，则右半段是有序的，若中间数大于最右边数，则左半段是有序的，我们只要在有序的半段里用首尾两个数组来判断目标值是否在这一区域内，这样就可以确定保留哪半边了，代码如下：
 
解法一：

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0, right = nums.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] == target) return mid;
            if (nums[mid] < nums[right]) {
                if (nums[mid] < target && nums[right] >= target) left = mid + 1;
                else right = mid - 1;
            } else {
                if (nums[left] <= target && nums[mid] > target) right = mid - 1;
                else left = mid + 1;
            }
        }
        return -1;
    }
};

 
看了上面的解法，你可能会产生个疑问，为啥非得用中间的数字跟最右边的比较呢？难道跟最左边的数字比较不行吗，当中间的数字大于最左边的数字时，左半段也是有序的啊，如下所示（蓝色表示中点之前一定为有序的）：
0　　1　　2　　 4　　5　　6　　7
7　　0　　1　　 2　　4　　5　　6
6　　7　　0　　 1　　2　　4　　5
5　　6　　7　　 0　　1　　2　　4
4　　5　　6　　7　　0　　1　　2
2　　4　　5　　6　　7　　0　　1
1　　2　　4　　5　　6　　7　　0
貌似也可以做，但是有一个问题，那就是在二分搜索中，nums[mid] 和 nums[left] 还有可能相等的，当数组中只有两个数字的时候，比如 [3, 1]，那该去取那一边呢？由于只有两个数字且 nums[mid] 不等于 target，target 只有可能在右半边出现。最好的方法就是让其无法进入左半段，就需要左半段是有序的，而且由于一定无法同时满足 nums[left] <= target && nums[mid] > target，因为 nums[left] 和 nums[mid] 相等，同一个数怎么可能同时大于等于 target，又小于 target。由于这个条件不满足，则直接进入右半段继续搜索即可，所以等于的情况要加到 nums[mid] > nums[left] 的情况中，变成大于等于，参见代码如下：
 
解法二：

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0, right = nums.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] == target) return mid;
            if (nums[mid] >= nums[left]) {
                if (nums[left] <= target && nums[mid] > target) right = mid - 1;
                else left = mid + 1;
            } else {
                if (nums[mid] < target && nums[right] >= target) left = mid + 1;
                else right = mid - 1;
            }
        }
        return -1;
    }
};