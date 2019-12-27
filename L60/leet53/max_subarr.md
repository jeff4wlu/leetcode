LeetCode [53. Maximum Subarray] 难度[easy]

腾古
2017.04.26 21:34:19
字数 694
阅读 701
题目
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

解题思路
题目大意是给定一个数组，让我们找出连续总和最大的子序列。有两种解决策略，第一个是利用动态规划，第二个是使用分治算法。（brute-force 就不提了）

1. Dynamic programming
一般来说，最优问题都可以利用动态规划的方法求解。而对于DP算法，最重要的事情是要找出子问题的递归形式。这题中，我们可以定义问题的形式为
maxSubArray(int A[], int i)，则假设已经求得子问题的解 maxSubArray(A, i-1)，在这基础上，我们要求原问题 maxSubArray(A, i)。这样一来，它们之间的关系就很容易弄清楚了：

maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i];

时间复杂度为 O(n), 参考代码如下。

2. Divide and conquer
第二个方法是分治算法，我们先来考虑下面这个问题，在数组 A[low..high] 中，最大子序列 A[i..j] 可能存在的位置（假设 mid = int((low+high)/2) ):

A. 完全在 A[low..mid] 里；（low≤i≤j≤mid)
B. 完全在 A[mid+1..high] 里；（mid＜i≤j≤high)
C. 跨过了数组的中点（low≤i≤mid＜j≤high)

如下图所示：


Three kinds of situations.png
则我们只需要求得上述三种情况的最大子序列，然后进行比较得出最大的那个即可满足题意。其中情况 A 和 B 可以用递归的方法求得。我们考虑下情况 C，如下图：

Crossing the midpoint.png
显然要使 A[i,j] 最大，则需要使 A[i..mid] 和 A[mid+1..j] 最大，所以只需分别遍历所有可能的 A[i..mid] 和 A[mid+1..j] 值（low≤i≤mid＜j≤high），并分别记录下最大值即可。
时间复杂度：
time complexity.png

参考代码如下：
参考代码
Dynamic programming：
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int n=nums.size();
        int dp[n]={0};
        dp[0]=nums[0];
        int ans=dp[0];
        for(int i=1;i<n;i++){
            if(dp[i-1]>0)
                dp[i]=dp[i-1]+nums[i];
            else
                dp[i]=nums[i];
            ans=max(ans,dp[i]);
        }
        return ans;
    }
};
Divide and conquer：
class Solution {
public:
    int find_max_crossing_subarray(vector<int>& A,int low,int mid, int high){
        
        // Find a maximum subarray of the form A[i..mid]
        int left_sum=INT_MIN, max_left;  
        int sum=0;
        for(int i=mid;i>=low;i--){
            sum+=A[i];
            if(sum>left_sum){
                left_sum = sum;
                max_left = i;
            }
        }
        
        // Find a maximum subarray of the form A[mid + 1 .. j ]
        int right_sum=INT_MIN, max_right;  
        sum=0;
        for(int i=mid+1;i<=high;i++){
            sum+=A[i];
            if(sum>right_sum){
                right_sum = sum;
                max_right = i;
            }
        }
        
        // Return the indices and the sum of the two subarrays
        return left_sum+right_sum;
    }
    
    int find_maximum_subarray(vector<int>& A,int low,int high){
        
        //base case: only one element
        if(high==low)        
            return A[low];
        else{
            int mid = (low+high)/2;
            int left_sum=find_maximum_subarray(A,low,mid);
            int right_sum=find_maximum_subarray(A,mid+1,high);
            int cross_sum=find_max_crossing_subarray(A,low,mid,high);
            if((left_sum>=right_sum)&&(left_sum>=cross_sum))
                return left_sum;
            else if((right_sum>=left_sum)&&(right_sum>=cross_sum))
                return right_sum;
            else
                return cross_sum;
        }
    }
    int maxSubArray(vector<int>& nums) {
        return find_maximum_subarray(nums,0,nums.size()-1);
    }
};
反思与总结
两种方法相比较，分治算法的思想简单直接，易编程，但时间复杂度较高。而 DP 算法更加高效，但是需要正确得出子问题的形式，较有技巧性。分治算法的最典型的代表就是归并排序了，所以比较熟悉，而在实际问题中，我往往会看不出其实可以利用 DP 算法来解决。但其实 DP 问题也有一定规律性，现总结如下：

