[LeetCode] 103. Binary Tree Zigzag Level Order Traversal 二叉树的之字形层序遍历 

 
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
 
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
 
这道二叉树的之字形层序遍历是之前那道 Binary Tree Level Order Traversal 的变形，不同之处在于一行是从左到右遍历，下一行是从右往左遍历，交叉往返的之字形的层序遍历。最简单直接的方法就是利用层序遍历，并使用一个变量 cnt 来统计当前的层数（从0开始），将所有的奇数层的结点值进行翻转一下即可，参见代码如下：
 
解法一：

class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        if (!root) return {};
        vector<vector<int>> res;
        queue<TreeNode*> q{{root}};
        int cnt = 0;
        while (!q.empty()) {
            vector<int> oneLevel;
            for (int i = q.size(); i > 0; --i) {
                TreeNode *t = q.front(); q.pop();
                oneLevel.push_back(t->val);
                if (t->left) q.push(t->left);
                if (t->right) q.push(t->right);
            }
            if (cnt % 2 == 1) reverse(oneLevel.begin(), oneLevel.end());
            res.push_back(oneLevel);
            ++cnt;
        }
        return res;
    }
};

 
我们可以将上面的解法进行优化一下，翻转数组虽然可行，但是比较耗时，假如能够直接计算出每个结点值在数组中的坐标，就可以直接进行更新了。由于每层的结点数是知道的，就是队列的元素个数，所以可以直接初始化数组的大小。此时使用一个变量 leftToRight 来标记顺序，初始时是 true，当此变量为 true 的时候，每次加入数组的位置就是i本身，若变量为 false 了，则加入到 size-1-i 位置上，这样就直接相当于翻转了数组。每层遍历完了之后，需要翻转 leftToRight 变量，同时不要忘了将 oneLevel 加入结果 res，参见代码如下：
 
解法二：

class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        if (!root) return {};
        vector<vector<int>> res;
        queue<TreeNode*> q{{root}};
        bool leftToRight = true;
        while (!q.empty()) {
            int size = q.size();
            vector<int> oneLevel(size);
            for (int i = 0; i < size; ++i) {
                TreeNode *t = q.front(); q.pop();
                int idx = leftToRight ? i : (size - 1 - i);
                oneLevel[idx] = t->val;
                if (t->left) q.push(t->left);
                if (t->right) q.push(t->right);
            }
            leftToRight = !leftToRight;
            res.push_back(oneLevel);
        }
        return res;
    }
};

 
我们也可以使用递归的方法来解，这里实际上用的是先序遍历，递归函数需要一个变量 level 来记录当前的深度，由于 level 是从0开始的，假如结果 res 的大小等于 level，就需要在结果 res 中新加一个空集，这样可以保证 res[level] 不会越界。取出 res[level] 之后，判断 levle 的奇偶，若其为偶数，则将 node->val 加入 oneLevel 的末尾，若为奇数，则加在 oneLevel 的开头。然后分别对 node 的左右子结点调用递归函数，此时要传入 level+1 即可，参见代码如下：
 
解法三：

class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> res;
        helper(root, 0, res);
        return res;
    }
    void helper(TreeNode* node, int level, vector<vector<int>>& res) {
        if (!node) return;
        if (res.size() <= level) {
            res.push_back({});
        }
        vector<int> &oneLevel = res[level];
        if (level % 2 == 0) oneLevel.push_back(node->val);
        else oneLevel.insert(oneLevel.begin(), node->val);
        helper(node->left, level + 1, res);
        helper(node->right, level + 1, res);
    }
};