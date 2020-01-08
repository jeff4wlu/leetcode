[LeetCode] 80. Remove Duplicates from Sorted Array II 有序数组中去除重复项之二 

 
Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
Example 1:
Given nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:
Given nums = [0,0,1,1,1,1,2,3,3],

Your function should return length = 7, with the first seven elements of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively.

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
 
这道题是之前那道 Remove Duplicates from Sorted Array 的拓展，这里允许最多重复的次数是两次，那么可以用一个变量 cnt 来记录还允许有几次重复，cnt 初始化为1，如果出现过一次重复，则 cnt 递减1，那么下次再出现重复，快指针直接前进一步，如果这时候不是重复的，则 cnt 恢复1，由于整个数组是有序的，所以一旦出现不重复的数，则一定比这个数大，此数之后不会再有重复项。理清了上面的思路，则代码很好写了：
 
解法一：

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int pre = 0, cur = 1, cnt = 1, n = nums.size();
        while (cur < n) {
            if (nums[pre] == nums[cur] && cnt == 0) ++cur;
            else {
                if (nums[pre] == nums[cur]) --cnt;
                else cnt = 1;
                nums[++pre] = nums[cur++];
            }
        }
        return nums.empty() ? 0 : pre + 1;
    }
};

 
这里其实也可以用类似于 Remove Duplicates from Sorted Array 中的解法三的模版，由于这里最多允许两次重复，那么当前的数字 num 只要跟上上个覆盖位置的数字 nusm[i-2] 比较，若 num 较大，则绝不会出现第三个重复数字（前提是数组是有序的），这样的话根本不需要管 nums[i-1] 是否重复，只要将重复个数控制在2个以内就可以了，参见代码如下：
 
解法二：

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int i = 0;
        for (int num : nums) {
            if (i < 2 || num > nums[i - 2]) {
                nums[i++] = num;
            }
        }
        return i;
    }
};