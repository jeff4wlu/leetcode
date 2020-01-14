[LeetCode] 104. Maximum Depth of Binary Tree 二叉树的最大深度 

 
Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
Note: A leaf is a node with no children.
Example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
 
求二叉树的最大深度问题用到深度优先搜索 Depth First Search，递归的完美应用，跟求二叉树的最小深度问题原理相同，参见代码如下：
 
C++ 解法一：

class Solution {
public:
    int maxDepth(TreeNode* root) {
        if (!root) return 0;
        return 1 + max(maxDepth(root->left), maxDepth(root->right));
    }
};

 
Java 解法一：
public class Solution {
    public int maxDepth(TreeNode root) {
        return root == null ? 0 : (1 + Math.max(maxDepth(root.left), maxDepth(root.right)));
    }
}
 
我们也可以使用层序遍历二叉树，然后计数总层数，即为二叉树的最大深度，注意 while 循环中的 for 循环的写法有个 trick，一定要将 q.size() 放在初始化里，而不能放在判断停止的条件中，因为q的大小是随时变化的，所以放停止条件中会出错，参见代码如下：
 
C++ 解法二：

class Solution {
public:
    int maxDepth(TreeNode* root) {
        if (!root) return 0;
        int res = 0;
        queue<TreeNode*> q{{root}};
        while (!q.empty()) {
            ++res;
            for (int i = q.size(); i > 0; --i) {
                TreeNode *t = q.front(); q.pop();
                if (t->left) q.push(t->left);
                if (t->right) q.push(t->right);
            }
        }
        return res;
    }
};

 
Java 解法二：

public class Solution {
    public int maxDepth(TreeNode root) {
        if (root == null) return 0;
        int res = 0;
        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
        while (!q.isEmpty()) {
            ++res;
            for (int i = q.size(); i > 0; --i) {
                TreeNode t = q.poll();
                if (t.left != null) q.offer(t.left);
                if (t.right != null) q.offer(t.right);
            }
        }
        return res;
    }
}