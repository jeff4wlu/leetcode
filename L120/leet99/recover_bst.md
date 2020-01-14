[LeetCode] 99. Recover Binary Search Tree 复原二叉搜索树 

 
Two elements of a binary search tree (BST) are swapped by mistake.
Recover the tree without changing its structure.
Example 1:
Input: [1,3,null,null,2]

   1
  /
 3
  \
   2

Output: [3,1,null,null,2]

   3
  /
 1
  \
   2
Example 2:
Input: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

Output: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
Follow up:
A solution using O(n) space is pretty straight forward.
Could you devise a constant space solution?
 
这道题要求我们复原一个二叉搜索树，说是其中有两个的顺序被调换了，题目要求上说 O(n) 的解法很直观，这种解法需要用到递归，用中序遍历树，并将所有节点存到一个一维向量中，把所有节点值存到另一个一维向量中，然后对存节点值的一维向量排序，在将排好的数组按顺序赋给节点。这种最一般的解法可针对任意个数目的节点错乱的情况，这里先贴上此种解法：
 
解法一：

// O(n) space complexity
class Solution {
public:
    void recoverTree(TreeNode* root) {
        vector<TreeNode*> list;
        vector<int> vals;
        inorder(root, list, vals);
        sort(vals.begin(), vals.end());
        for (int i = 0; i < list.size(); ++i) {
            list[i]->val = vals[i];
        }
    }
    void inorder(TreeNode* root, vector<TreeNode*>& list, vector<int>& vals) {
        if (!root) return;
        inorder(root->left, list, vals);
        list.push_back(root);
        vals.push_back(root->val);
        inorder(root->right, list, vals);
    }
};

 
然后博主上网搜了许多其他解法，看到另一种是用双指针来代替一维向量的，但是这种方法用到了递归，也不是 O(1) 空间复杂度的解法，这里需要三个指针，first，second 分别表示第一个和第二个错乱位置的节点，pre 指向当前节点的中序遍历的前一个节点。这里用传统的中序遍历递归来做，不过再应该输出节点值的地方，换成了判断 pre 和当前节点值的大小，如果 pre 的大，若 first 为空，则将 first 指向 pre 指的节点，把 second 指向当前节点。这样中序遍历完整个树，若 first 和 second 都存在，则交换它们的节点值即可。这个算法的空间复杂度仍为 O(n)，n为树的高度，参见代码如下：
 
解法二：

// Still O(n) space complexity
class Solution {
public:
    TreeNode *pre = NULL, *first = NULL, *second = NULL;
    void recoverTree(TreeNode* root) {
        inorder(root);
        swap(first->val, second->val);
    }
    void inorder(TreeNode* root) {
        if (!root) return;
        inorder(root->left);
        if (!pre) pre = root;
        else {
            if (pre->val > root->val) {
                if (!first) first = pre;
                second = root;
            }
            pre = root;
        }
        inorder(root->right);
    }
};

 
我们其实也可以使用迭代的写法，因为中序遍历 Binary Tree Inorder Traversal 也可以借助栈来实现，原理还是跟前面的相同，记录前一个结点，并和当前结点相比，如果前一个结点值大，那么更新 first 和 second，最后交换 first 和 second 的结点值即可，参见代码如下：
 
解法三：

// Always O(n) space complexity
class Solution {
public:
    void recoverTree(TreeNode* root) {
        TreeNode *pre = NULL, *first = NULL, *second = NULL, *p = root;
        stack<TreeNode*> st;
        while (p || !st.empty()) {
            while (p) {
                st.push(p);
                p = p->left;
            }
            p = st.top(); st.pop();
            if (pre) {
                if (pre->val > p->val) {
                    if (!first) first = pre;
                    second = p;
                }
            }
            pre = p;
            p = p->right;
        }
        swap(first->val, second->val);
    }
};

 
这道题的真正符合要求的解法应该用的 Morris 遍历，这是一种非递归且不使用栈，空间复杂度为 O(1) 的遍历方法，可参见博主之前的博客 Binary Tree Inorder Traversal，在其基础上做些修改，加入 first, second 和 parent 指针，来比较当前节点值和中序遍历的前一节点值的大小，跟上面递归算法的思路相似，代码如下：
 
解法四：

// Now O(1) space complexity
class Solution {
public:
    void recoverTree(TreeNode* root) {
        TreeNode *first = nullptr, *second = nullptr, *cur = root, *pre = nullptr ;
        while (cur) {
            if (cur->left){
                TreeNode *p = cur->left;
                while (p->right && p->right != cur) p = p->right;
                if (!p->right) {
                    p->right = cur;
                    cur = cur->left;
                    continue;
                } else {
                    p->right = NULL;
                }  
            }
            if (pre && cur->val < pre->val){
              if (!first) first = pre;
              second = cur;
            }
            pre = cur;
            cur = cur->right;
        }
        swap(first->val, second->val);
    }
};