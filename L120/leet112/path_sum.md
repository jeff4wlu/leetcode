[LeetCode] 112. Path Sum 二叉树的路径和 

 
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
Note: A leaf is a node with no children.
Example:
Given the below binary tree and sum = 22,
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
 
这道题给了一棵二叉树，问是否存在一条从跟结点到叶结点到路径，使得经过到结点值之和为一个给定的 sum 值，这里需要用深度优先算法 DFS 的思想来遍历每一条完整的路径，也就是利用递归不停找子结点的左右子结点，而调用递归函数的参数只有当前结点和 sum 值。首先，如果输入的是一个空结点，则直接返回 false，如果如果输入的只有一个根结点，则比较当前根结点的值和参数 sum 值是否相同，若相同，返回 true，否则 false。 这个条件也是递归的终止条件。下面就要开始递归了，由于函数的返回值是 Ture/False，可以同时两个方向一起递归，中间用或 || 连接，只要有一个是 True，整个结果就是 True。递归左右结点时，这时候的 sum 值应该是原 sum 值减去当前结点的值，参见代码如下：
 
解法一：

class Solution {
public:
    bool hasPathSum(TreeNode* root, int sum) {
        if (!root) return false;
        if (!root->left && !root->right && root->val == sum ) return true;
        return hasPathSum(root->left, sum - root->val) || hasPathSum(root->right, sum - root->val);
    }
};

 
我们也可以使用迭代的写法，这里用的也是先序遍历的迭代写法，先序遍历二叉树，左右子结点都需要加上其父结点值，这样当遍历到叶结点时，如果和 sum 相等了，那么就说明一定有一条从 root 过来的路径。注意这里不必一定要先处理右子结点，调换下顺序也是可以的，因为不论是先序遍历的根-左-右，还是根-右-左，并不会影响到找路径，参见代码如下：
 
解法二：

class Solution {
public:
    bool hasPathSum(TreeNode* root, int sum) {
        if (!root) return false;
        stack<TreeNode*> st{{root}};
        while (!st.empty()) {
            TreeNode *t = st.top(); st.pop();
            if (!t->left && !t->right) {
                if (t->val == sum) return true;
            }
            if (t->right) {
                t->right->val += t->val;
                st.push(t->right);
            }
            if (t->left) {
                t->left->val += t->val;
                st.push(t->left);
            }
        }
        return false;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/112
 
类似题目：
Path Sum IV 
Path Sum III  
Binary Tree Maximum Path Sum
Path Sum II
Sum Root to Leaf Numbers
Binary Tree Preorder Traversal