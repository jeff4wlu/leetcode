[LeetCode] 230. Kth Smallest Element in a BST 二叉搜索树中的第K小的元素 

 
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
Note: 
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
Example 1:
Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

Credits:
Special thanks to @ts for adding this problem and creating all test cases.
 
这又是一道关于二叉搜索树 Binary Search Tree 的题， LeetCode 中关于 BST 的题有 Validate Binary Search Tree， Recover Binary Search Tree， Binary Search Tree Iterator， Unique Binary Search Trees， Unique Binary Search Trees II，Convert Sorted Array to Binary Search Tree 和 Convert Sorted List to Binary Search Tree。那么这道题给的提示是让我们用 BST 的性质来解题，最重要的性质是就是左<根<右，如果用中序遍历所有的节点就会得到一个有序数组。所以解题的关键还是中序遍历啊。关于二叉树的中序遍历可以参见我之前的博客 Binary Tree Inorder Traversal，里面有很多种方法可以用，先来看一种非递归的方法，中序遍历最先遍历到的是最小的结点，只要用一个计数器，每遍历一个结点，计数器自增1，当计数器到达k时，返回当前结点值即可，参见代码如下：
 
解法一：

class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        int cnt = 0;
        stack<TreeNode*> s;
        TreeNode *p = root;
        while (p || !s.empty()) {
            while (p) {
                s.push(p);
                p = p->left;
            }
            p = s.top(); s.pop();
            ++cnt;
            if (cnt == k) return p->val;
            p = p->right;
        }
        return 0;
    }
};

 
当然，此题我们也可以用递归来解，还是利用中序遍历来解，代码如下：
 
解法二：

class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        return kthSmallestDFS(root, k);
    }
    int kthSmallestDFS(TreeNode* root, int &k) {
        if (!root) return -1;
        int val = kthSmallestDFS(root->left, k);
        if (k == 0) return val;
        if (--k == 0) return root->val;
        return kthSmallestDFS(root->right, k);
    }
};

 
再来看一种分治法的思路，由于 BST 的性质，可以快速定位出第k小的元素是在左子树还是右子树，首先计算出左子树的结点个数总和 cnt，如果k小于等于左子树结点总和 cnt，说明第k小的元素在左子树中，直接对左子结点调用递归即可。如果k大于 cnt+1，说明目标值在右子树中，对右子结点调用递归函数，注意此时的k应为 k-cnt-1，应为已经减少了 cnt+1 个结点。如果k正好等于 cnt+1，说明当前结点即为所求，返回当前结点值即可，参见代码如下：
 
解法三：

class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        int cnt = count(root->left);
        if (k <= cnt) {
            return kthSmallest(root->left, k);
        } else if (k > cnt + 1) {
            return kthSmallest(root->right, k - cnt - 1);
        }
        return root->val;
    }
    int count(TreeNode* node) {
        if (!node) return 0;
        return 1 + count(node->left) + count(node->right);
    }
};

 
这道题的 Follow up 中说假设该 BST 被修改的很频繁，而且查找第k小元素的操作也很频繁，问我们如何优化。其实最好的方法还是像上面的解法那样利用分治法来快速定位目标所在的位置，但是每个递归都遍历左子树所有结点来计算个数的操作并不高效，所以应该修改原树结点的结构，使其保存包括当前结点和其左右子树所有结点的个数，这样就可以快速得到任何左子树结点总数来快速定位目标值了。定义了新结点结构体，然后就要生成新树，还是用递归的方法生成新树，注意生成的结点的 count 值要累加其左右子结点的 count 值。然后在求第k小元素的函数中，先生成新的树，然后调用递归函数。在递归函数中，不能直接访问左子结点的 count 值，因为左子节结点不一定存在，所以要先判断，如果左子结点存在的话，那么跟上面解法的操作相同。如果不存在的话，当此时k为1的时候，直接返回当前结点值，否则就对右子结点调用递归函数，k自减1，参见代码如下：
 
解法四：

// Follow up
class Solution {
public:
    struct MyTreeNode {
        int val;
        int count;
        MyTreeNode *left;
        MyTreeNode *right;
        MyTreeNode(int x) : val(x), count(1), left(NULL), right(NULL) {}
    };
    
    MyTreeNode* build(TreeNode* root) {
        if (!root) return NULL;
        MyTreeNode *node = new MyTreeNode(root->val);
        node->left = build(root->left);
        node->right = build(root->right);
        if (node->left) node->count += node->left->count;
        if (node->right) node->count += node->right->count;
        return node;
    }
    
    int kthSmallest(TreeNode* root, int k) {
        MyTreeNode *node = build(root);
        return helper(node, k);
    }
    
    int helper(MyTreeNode* node, int k) {
        if (node->left) {
            int cnt = node->left->count;
            if (k <= cnt) {
                return helper(node->left, k);
            } else if (k > cnt + 1) {
                return helper(node->right, k - 1 - cnt);
            }
            return node->val;
        } else {
            if (k == 1) return node->val;
            return helper(node->right, k - 1);
        }
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/230
 
类似题目：
Binary Tree Inorder Traversal
Second Minimum Node In a Binary Tree