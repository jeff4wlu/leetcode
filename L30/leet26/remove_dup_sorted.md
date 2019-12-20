[LeetCode] 26. Remove Duplicates from Sorted Array 有序数组中去除重复项 

 
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
Example 1:
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
Clarification:
Confused why the returned value is an integer but your answer is an array?
Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.
Internally you can think of this:
// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
 
这道题要我们从有序数组中去除重复项，和之前那道 Remove Duplicates from Sorted List 的题很类似，但是要简单一些，因为毕竟数组的值可以通过下标直接访问，而链表不行。那么这道题的解题思路是使用快慢指针来记录遍历的坐标，最开始时两个指针都指向第一个数字，如果两个指针指的数字相同，则快指针向前走一步，如果不同，则两个指针都向前走一步，这样当快指针走完整个数组后，慢指针当前的坐标加1就是数组中不同数字的个数，代码如下：
 
解法一：

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int pre = 0, cur = 0, n = nums.size();
        while (cur < n) {
            if (nums[pre] == nums[cur]) ++cur;
            else nums[++pre] = nums[cur++];
        }
        return nums.empty() ? 0 : (pre + 1);
    }
};

 
我们也可以用 for 循环来写，这里的j就是上面解法中的 pre，i就是 cur，所以本质上都是一样的，参见代码如下：
 
解法二：

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int j = 0, n = nums.size();
        for (int i = 0; i < n; ++i) {
            if (nums[i] != nums[j]) nums[++j] = nums[i];
        }
        return nums.empty() ? 0 : (j + 1);
    }
};

 
这里也可以换一种写法，用变量i表示当前覆盖到到位置，由于不能有重复数字，则只需要用当前数字 num 跟上一个覆盖到到数字 nums[i-1] 做个比较，只要 num 大，则一定不会有重复（前提是数组必须有序），参见代码如下：
 
解法三：

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int i = 0;
        for (int num : nums) {
            if (i < 1 || num > nums[i - 1]) {
                nums[i++] = num;
            }
        }
        return i;
    }
};
