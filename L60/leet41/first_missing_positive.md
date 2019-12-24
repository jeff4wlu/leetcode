[LeetCode] 41. First Missing Positive 首个缺失的正数 

 
Given an unsorted integer array, find the smallest missing positive integer.
Example 1:
Input: [1,2,0]
Output: 3
Example 2:
Input: [3,4,-1,1]
Output: 2
Example 3:
Input: [7,8,9,11,12]
Output: 1
Note:
Your algorithm should run in O(n) time and uses constant extra space.
 
这道题让我们找缺失的首个正数，由于限定了 O(n) 的时间，所以一般的排序方法都不能用，最开始博主没有看到还限制了空间复杂度，所以想到了用 HashSet 来解，这个思路很简单，把所有的数都存入 HashSet 中，然后循环从1开始递增找数字，哪个数字找不到就返回哪个数字，如果一直找到了最大的数字（这里是 nums 数组的长度），则加1后返回结果 res，参见代码如下：
 
解法一：

// NOT constant space
class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        unordered_set<int> st(nums.begin(), nums.end());
        int res = 1, n = nums.size();
        while (res <= n) {
            if (!st.count(res)) return res;
            ++res;
        }
        return res;
    }
};

 
但是上面的解法不是 O(1) 的空间复杂度，所以需要另想一种解法，既然不能建立新的数组，那么只能覆盖原有数组，思路是把1放在数组第一个位置 nums[0]，2放在第二个位置 nums[1]，即需要把 nums[i] 放在 nums[nums[i] - 1]上，遍历整个数组，如果 nums[i] != i + 1, 而 nums[i] 为整数且不大于n，另外 nums[i] 不等于 nums[nums[i] - 1] 的话，将两者位置调换，如果不满足上述条件直接跳过，最后再遍历一遍数组，如果对应位置上的数不正确则返回正确的数，参见代码如下：
 
解法二：

class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        int n = nums.size();
        for (int i = 0; i < n; ++i) {
            while (nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                swap(nums[i], nums[nums[i] - 1]);
            }
        }
        for (int i = 0; i < n; ++i) {
            if (nums[i] != i + 1) return i + 1;
        }
        return n + 1;
    }
};


举个例子，假设有序列[4,2,6,1,-3]，首先看第一个数4，它正确的位置应该是在序列的第4个位置（位置数从1开始，正确的位置是第一个位置放1，第二个位置放2，第三个位置放3……最后我们只要看哪个位置放的不是理想的数，那么它就是第一个缺失的正数）。我们将4与第4个位置上的“1”进行交换，序列变成[1,2,6,4,-3]；接着我们还是看第一个数，现在变成了“1”，它的确在它正确的位置，好了，我们再看第二个数2，也在正确的位置。第三个数6，本来应该放在第6个位置，可是该序列总共就5个位置，所以不移动；第四个数4在它的正确位置，不动；第五个数是负数，不动。最后，从前往后看，发现在第三个位置本该出现的3没有出现，所有该序列缺失的第一个正数是3。
————————————————
版权声明：本文为CSDN博主「Jaster_wisdom」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/Jaster_wisdom/article/details/80660467