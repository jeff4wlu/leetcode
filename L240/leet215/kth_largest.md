[LeetCode] 215. Kth Largest Element in an Array 数组中第k大的数字 

 
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
Example 1:
Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note: 
You may assume k is always valid, 1 ≤ k ≤ array's length.
 
这道题让我们求数组中第k大的数字，怎么求呢，当然首先想到的是给数组排序，然后求可以得到第k大的数字。先看一种利用 C++ 的 STL 中的集成的排序方法，不用我们自己实现，这样的话这道题只要两行就完事了，代码如下：
 
解法一：

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        sort(nums.begin(), nums.end());
        return nums[nums.size() - k];
    }
};

 
下面这种解法是利用了 priority_queue 的自动排序的特性，跟上面的解法思路上没有什么区别，当然我们也可以换成 multiset 来做，一个道理，参见代码如下：
 
解法二：

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int> q(nums.begin(), nums.end());
        for (int i = 0; i < k - 1; ++i) {
            q.pop();
        }
        return q.top();
    }
};

 
上面两种方法虽然简洁，但是确不是本题真正想考察的东西，可以说有一定的偷懒嫌疑。这道题最好的解法应该是下面这种做法，用到了快速排序 Quick Sort 的思想，这里排序的方向是从大往小排。对快排不熟悉的童鞋们随意上网搜些帖子看下吧，多如牛毛啊，总有一款适合你。核心思想是每次都要先找一个中枢点 Pivot，然后遍历其他所有的数字，像这道题从大往小排的话，就把大于中枢点的数字放到左半边，把小于中枢点的放在右半边，这样中枢点是整个数组中第几大的数字就确定了，虽然左右两部分各自不一定是完全有序的，但是并不影响本题要求的结果，因为左半部分的所有值都大于右半部分的任意值，所以我们求出中枢点的位置，如果正好是 k-1，那么直接返回该位置上的数字；如果大于 k-1，说明要求的数字在左半部分，更新右边界，再求新的中枢点位置；反之则更新右半部分，求中枢点的位置；不得不说，这个思路真的是巧妙啊～
 
解法三：

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        int left = 0, right = nums.size() - 1;
        while (true) {
            int pos = partition(nums, left, right);
            if (pos == k - 1) return nums[pos];
            if (pos > k - 1) right = pos - 1;
            else left = pos + 1;
        }
    }
    int partition(vector<int>& nums, int left, int right) {
        int pivot = nums[left], l = left + 1, r = right;
        while (l <= r) {
            if (nums[l] < pivot && nums[r] > pivot) {
                swap(nums[l++], nums[r--]);
            }
            if (nums[l] >= pivot) ++l;
            if (nums[r] <= pivot) --r;
        }
        swap(nums[left], nums[r]);
        return r;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/215
 
类似题目：
Wiggle Sort II
Top K Frequent Elements
Third Maximum Number
Kth Largest Element in a Stream
K Closest Points to Origin