[LeetCode] 255. Verify Preorder Sequence in Binary Search Tree 验证二叉搜索树的先序序列 

 
Given an array of numbers, verify whether it is the correct preorder traversal sequence of a binary search tree.
You may assume each number in the sequence is unique.
Consider the following binary search tree: 
     5
    / \
   2   6
  / \
 1   3
Example 1:
Input: [5,2,6,1,3]
Output: false
Example 2:
Input: [5,2,1,3,6]
Output: true
Follow up:
Could you do it using only constant space complexity?
 
这道题让给了一个一维数组，让我们验证其是否为一个二叉搜索树的先序遍历出的顺序，二叉搜索树的性质是左<根<右，如果用中序遍历得到的结果就是有序数组，而先序遍历的结果就不是有序数组了，但是难道一点规律都没有了吗，其实规律还是有的，根据二叉搜索树的性质，当前节点的值一定大于其左子树中任何一个节点值，而且其右子树中的任何一个节点值都不能小于当前节点值，可以用这个性质来验证，举个例子，比如下面这棵二叉搜索树：
 
     5
    / \
   2   6
  / \
 1   3
 
其先序遍历的结果是 {5, 2, 1, 3, 6}，先设一个最小值 low，然后遍历数组，如果当前值小于这个最小值 low，返回 false，对于根节点，将其压入栈中，然后往后遍历，如果遇到的数字比栈顶元素小，说明是其左子树的点，继续压入栈中，直到遇到的数字比栈顶元素大，那么就是右边的值了，需要找到是哪个节点的右子树，所以更新 low 值并删掉栈顶元素，然后继续和下一个栈顶元素比较，如果还是大于，则继续更新 low 值和删掉栈顶，直到栈为空或者当前栈顶元素大于当前值停止，压入当前值，这样如果遍历完整个数组之前都没有返回 false 的话，最后返回 true 即可，参见代码如下：
 
解法一：

class Solution {
public:
    bool verifyPreorder(vector<int>& preorder) {
        int low = INT_MIN;
        stack<int> s;
        for (auto a : preorder) {
            if (a < low) return false;
            while (!s.empty() && a > s.top()) {
                low = s.top(); s.pop();
            }
            s.push(a);
        }
        return true;
    }
};

 
下面这种方法和上面的思路相同，为了使空间复杂度为常量，我们不能使用 stack，所以直接修改 preorder，将 low 值存在 preorder 的特定位置即可，前提是不能影响当前的遍历，参见代码如下：
 
解法二：

class Solution {
public:
    bool verifyPreorder(vector<int>& preorder) {
        int low = INT_MIN, i = -1;
        for (auto a : preorder) {
            if (a < low) return false;
            while (i >= 0 && a > preorder[i]) {
                low = preorder[i--];
            }
            preorder[++i] = a;
        }
        return true;
    }
};

 
下面这种方法使用了分治法，跟之前那道验证二叉搜索树的题 Validate Binary Search Tree 的思路很类似，在递归函数中维护一个下界 lower 和上界 upper，那么当前遍历到的节点值必须在 (lower, upper) 区间之内，然后在给定的区间内搜第一个大于当前节点值的点，以此为分界，左右两部分分别调用递归函数，注意左半部分的 upper 更新为当前节点值 val，表明左子树的节点值都必须小于当前节点值，而右半部分的递归的 lower 更新为当前节点值 val，表明右子树的节点值都必须大于当前节点值，如果左右两部分的返回结果均为真，则整体返回真，参见代码如下：
 
解法三：

class Solution {
public:
    bool verifyPreorder(vector<int>& preorder) {
        return helper(preorder, 0, preorder.size() - 1, INT_MIN, INT_MAX);
    }
    bool helper(vector<int>& preorder, int start, int end, int lower, int upper) {
        if (start > end) return true;
        int val = preorder[start], i = 0;
        if (val <= lower || val >= upper) return false;
        for (i = start + 1; i <= end; ++i) {
            if (preorder[i] >= val) break;
        }
        return helper(preorder, start + 1, i - 1, lower, val) && helper(preorder, i, end, val, upper);
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/255
 
类似题目：
Binary Tree Preorder Traversal
Validate Binary Search Tree