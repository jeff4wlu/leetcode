[LeetCode] 94. Binary Tree Inorder Traversal 二叉树的中序遍历 

 
Given a binary tree, return the inorder traversal of its nodes' values.
Example:
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
 
二叉树的中序遍历顺序为左-根-右，可以有递归和非递归来解，其中非递归解法又分为两种，一种是使用栈来接，另一种不需要使用栈。我们先来看递归方法，十分直接，对左子结点调用递归函数，根节点访问值，右子节点再调用递归函数，代码如下：
 
解法一：

class Solution {
public:
    vector<int> inorderTraversal(TreeNode *root) {
        vector<int> res;
        inorder(root, res);
        return res;
    }
    void inorder(TreeNode *root, vector<int> &res) {
        if (!root) return;
        if (root->left) inorder(root->left, res);
        res.push_back(root->val);
        if (root->right) inorder(root->right, res);
    }
};

 
下面再来看非递归使用栈的解法，也是符合本题要求使用的解法之一，需要用栈来做，思路是从根节点开始，先将根节点压入栈，然后再将其所有左子结点压入栈，然后取出栈顶节点，保存节点值，再将当前指针移到其右子节点上，若存在右子节点，则在下次循环时又可将其所有左子结点压入栈中。这样就保证了访问顺序为左-根-右，代码如下：
 
解法二： 

// Non-recursion
class Solution {
public:
    vector<int> inorderTraversal(TreeNode *root) {
        vector<int> res;
        stack<TreeNode*> s;
        TreeNode *p = root;
        while (p || !s.empty()) {
            while (p) {
                s.push(p);
                p = p->left;
            }
            p = s.top(); s.pop();
            res.push_back(p->val);
            p = p->right;
        }
        return res;
    }
};

 
下面这种解法跟 Binary Tree Preorder Traversal 中的解法二几乎一样，就是把结点值加入结果 res 的步骤从 if 中移动到了 else 中，因为中序遍历的顺序是左-根-右，参见代码如下：
 
解法三：

class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        stack<TreeNode*> s;
        TreeNode *p = root;
        while (!s.empty() || p) {
            if (p) {
                s.push(p);
                p = p->left;
            } else {
                p = s.top(); s.pop();
                res.push_back(p->val);
                p = p->right;
            }
        }
        return res;
    }
};

 
下面我们来看另一种很巧妙的解法，这种方法不需要使用栈，所以空间复杂度为常量，这种非递归不用栈的遍历方法有个专门的名字，叫 Morris Traversal，在介绍这种方法之前，我们先来引入一种新型树，叫 Threaded binary tree，这个还不太好翻译，我第一眼看上去以为是叫线程二叉树，但是感觉好像又跟线程没啥关系，后来看到网上有人翻译为螺纹二叉树，但本人认为这翻译也不太敢直视，很容易让人联想到为计划生育做出突出贡献的某世界著名品牌，但是苦于找不到更合理的翻译方法，就暂且叫螺纹二叉树吧。我们先来看看维基百科上关于它的英文定义：
A binary tree is threaded by making all right child pointers that would normally be null point to the inorder successor of the node (if it exists), and all left child pointers that would normally be null point to the inorder predecessor of the node.
就是说螺纹二叉树实际上是把所有原本为空的右子节点指向了中序遍历顺序之后的那个节点，把所有原本为空的左子节点都指向了中序遍历之前的那个节点，具体例子可以点击这里。那么这道题跟这个螺纹二叉树又有啥关系呢？由于我们既不能用递归，又不能用栈，那我们如何保证访问顺序是中序遍历的左-根-右呢。原来我们需要构建一个螺纹二叉树，需要将所有为空的右子节点指向中序遍历的下一个节点，这样中序遍历完左子结点后，就能顺利的回到其根节点继续遍历了。具体算法如下：
1. 初始化指针 cur 指向 root
2. 当 cur 不为空时
　 - 如果 cur 没有左子结点
　     a) 打印出 cur 的值
　　  b) 将 cur 指针指向其右子节点
　 - 反之
　    将 pre 指针指向 cur 的左子树中的最右子节点　
　　　  * 若 pre 不存在右子节点
　　　       a) 将其右子节点指回 cur
　　　　    b) cur 指向其左子节点
　　　  * 反之
　　　　　 a) 将 pre 的右子节点置空
　　　　　 b) 打印 cur 的值
　　　　　 c) 将 cur 指针指向其右子节点
 
解法四：

class Solution {
public:
    vector<int> inorderTraversal(TreeNode *root) {
        vector<int> res;
        if (!root) return res;
        TreeNode *cur, *pre;
        cur = root;
        while (cur) {
            if (!cur->left) {
                res.push_back(cur->val);
                cur = cur->right;
            } else {
                pre = cur->left;
                while (pre->right && pre->right != cur) pre = pre->right;
                if (!pre->right) {
                    pre->right = cur;
                    cur = cur->left;
                } else {
                    pre->right = NULL;
                    res.push_back(cur->val);
                    cur = cur->right;
                }
            }
        }
        return res;
    }
};

 
其实 Morris 遍历不仅仅对中序遍历有用，对先序和后序同样有用，具体可参见网友 NOALGO 博客，和 Annie Kim's Blog 的博客。所以对二叉树的三种常见遍历顺序(先序，中序，后序)就有三种解法(递归，非递归，Morris 遍历)，总共有九段代码呀，熟练掌握这九种写法才算初步掌握了树的遍历挖~~ 至于二叉树的层序遍历也有递归和非递归解法，至于有没有 Morris 遍历的解法还有待大神们的解答，若真有也请劳烦告知博主一声~~