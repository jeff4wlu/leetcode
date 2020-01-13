[LeetCode] 95. Unique Binary Search Trees II 独一无二的二叉搜索树之二 

 
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
Example:
Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
 
这道题是之前的 Unique Binary Search Trees 的延伸，之前那个只要求算出所有不同的二叉搜索树的个数，这道题让把那些二叉树都建立出来。这种建树问题一般来说都是用递归来解，这道题也不例外，划分左右子树，递归构造。这个其实是用到了大名鼎鼎的分治法 Divide and Conquer，类似的题目还有之前的那道 Different Ways to Add Parentheses 用的方法一样，用递归来解，划分左右两个子数组，递归构造。刚开始时，将区间 [1, n] 当作一个整体，然后需要将其中的每个数字都当作根结点，其划分开了左右两个子区间，然后分别调用递归函数，会得到两个结点数组，接下来要做的就是从这两个数组中每次各取一个结点，当作当前根结点的左右子结点，然后将根结点加入结果 res 数组中即可，参见代码如下：
 
解法一：

class Solution {
public:
    vector<TreeNode*> generateTrees(int n) {
        if (n == 0) return {};
        return helper(1, n);
    }
    vector<TreeNode*> helper(int start, int end) {
        if (start > end) return {nullptr};
        vector<TreeNode*> res;
        for (int i = start; i <= end; ++i) {
            auto left = helper(start, i - 1), right = helper(i + 1, end);
            for (auto a : left) {
                for (auto b : right) {
                    TreeNode *node = new TreeNode(i);
                    node->left = a;
                    node->right = b;
                    res.push_back(node);
                }
            }
        }
        return res;
    }
};

 
我们可以使用记忆数组来优化，保存计算过的中间结果，从而避免重复计算。注意这道题的标签有一个是动态规划 Dynamic Programming，其实带记忆数组的递归形式就是 DP 的一种，memo[i][j] 表示在区间 [i, j] 范围内可以生成的所有 BST 的根结点，所以 memo 必须是一个三维数组，这样在递归函数中，就可以去 memo 中查找当前的区间是否已经计算过了，是的话，直接返回 memo 中的数组，否则就按之前的方法去计算，最后计算好了之后要更新 memo 数组，参见代码如下：
 
解法二：

class Solution {
public:
    vector<TreeNode*> generateTrees(int n) {
        if (n == 0) return {};
        vector<vector<vector<TreeNode*>>> memo(n, vector<vector<TreeNode*>>(n));
        return helper(1, n, memo);
    }
    vector<TreeNode*> helper(int start, int end, vector<vector<vector<TreeNode*>>>& memo) {
        if (start > end) return {nullptr};
        if (!memo[start - 1][end - 1].empty()) return memo[start - 1][end - 1];
        vector<TreeNode*> res;
        for (int i = start; i <= end; ++i) {
            auto left = helper(start, i - 1, memo), right = helper(i + 1, end, memo);
            for (auto a : left) {
                for (auto b : right) {
                    TreeNode *node = new TreeNode(i);
                    node->left = a;
                    node->right = b;
                    res.push_back(node);
                }
            }
        }
        return memo[start - 1][end - 1] = res;
    }
};
