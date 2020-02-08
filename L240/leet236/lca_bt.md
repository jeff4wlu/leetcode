[LeetCode] 236. Lowest Common Ancestor of a Binary Tree 二叉树的最小共同父节点 

 
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]

 
Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
Example 2:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
 
Note:
All of the nodes' values will be unique.
p and q are different and both values will exist in the binary tree.
 
这道求二叉树的最小共同父节点的题是之前那道 Lowest Common Ancestor of a Binary Search Tree 的 Follow Up。跟之前那题不同的地方是，这道题是普通是二叉树，不是二叉搜索树，所以就不能利用其特有的性质，我们只能在二叉树中来搜索p和q，然后从路径中找到最后一个相同的节点即为父节点，可以用递归来实现，在递归函数中，首先看当前结点是否为空，若为空则直接返回空，若为p或q中的任意一个，也直接返回当前结点。否则的话就对其左右子结点分别调用递归函数，由于这道题限制了p和q一定都在二叉树中存在，那么如果当前结点不等于p或q，p和q要么分别位于左右子树中，要么同时位于左子树，或者同时位于右子树，那么我们分别来讨论：
- 若p和q分别位于左右子树中，那么对左右子结点调用递归函数，会分别返回p和q结点的位置，而当前结点正好就是p和q的最小共同父结点，直接返回当前结点即可，这就是题目中的例子1的情况。
- 若p和q同时位于左子树，这里有两种情况，一种情况是 left 会返回p和q中较高的那个位置，而 right 会返回空，所以最终返回非空的 left 即可，这就是题目中的例子2的情况。还有一种情况是会返回p和q的最小父结点，就是说当前结点的左子树中的某个结点才是p和q的最小父结点，会被返回。
- 若p和q同时位于右子树，同样这里有两种情况，一种情况是 right 会返回p和q中较高的那个位置，而 left 会返回空，所以最终返回非空的 right 即可，还有一种情况是会返回p和q的最小父结点，就是说当前结点的右子树中的某个结点才是p和q的最小父结点，会被返回，写法很简洁，代码如下：
 
解法一：

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
       if (!root || p == root || q == root) return root;
       TreeNode *left = lowestCommonAncestor(root->left, p, q);
       TreeNode *right = lowestCommonAncestor(root->right, p , q);
       if (left && right) return root;
       return left ? left : right;
    }
};

 
上述代码可以进行优化一下，如果当前结点不为空，且既不是p也不是q，那么根据上面的分析，p和q的位置就有三种情况，p和q要么分别位于左右子树中，要么同时位于左子树，或者同时位于右子树。我们需要优化的情况就是当p和q同时为于左子树或右子树中，而且返回的结点并不是p或q，那么就是p和q的最小父结点了，已经求出来了，就不用再对右结点调用递归函数了，这是为啥呢？因为根本不会存在 left 既不是p也不是q，同时还有p或者q在 right 中。首先递归的第一句就限定了只要遇到了p或者q，就直接返回，之后又限定了只有当 left 和 right 同时存在的时候，才会返回当前结点，当前结点若不是p或q，则一定是最小父节点，否则 left 一定是p或者q。这里的逻辑比较绕，不太好想，多想想应该可以理清头绪吧，参见代码如下：
 
解法二：

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
       if (!root || p == root || q == root) return root;
       TreeNode *left = lowestCommonAncestor(root->left, p, q);
       if (left && left != p && left != q) return left;
       TreeNode *right = lowestCommonAncestor(root->right, p , q);
　　　　if (left && right) return root;
       return left ? left : right;
    }
};

 
讨论：此题还有一种情况，题目中没有明确说明p和q是否是树中的节点，如果不是，应该返回 nullptr，而上面的方法就不正确了，对于这种情况请参见 Cracking the Coding Interview 5th Edition 的第 233-234 页。
 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/236
 
类似题目：
Lowest Common Ancestor of a Binary Search Tree