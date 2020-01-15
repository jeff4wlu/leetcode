LeetCode（116）：填充同一层的兄弟节点 

Medium！
题目描述：
给定一个二叉树
struct TreeLinkNode {
  TreeLinkNode *left;
  TreeLinkNode *right;
  TreeLinkNode *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
说明:
你只能使用额外常数空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
你可以假设它是一个完美二叉树（即所有叶子节点都在同一层，每个父节点都有两个子节点）。
示例:
给定完美二叉树，
     1
   /  \
  2    3
 / \  / \
4  5  6  7
调用你的函数后，该完美二叉树变为：
     1 -> NULL
   /  \
  2 -> 3 -> NULL
 / \  / \
4->5->6->7 -> NULL
解题思路：
这道题实际上是树的层序遍历的应用，可以参考之前的博客Binary Tree Level Order Traversal 二叉树层序遍历，既然是遍历，就有递归和非递归两种方法，最好两种方法都要掌握，都要会写。
面先来看递归的解法，由于是完全二叉树，所以若节点的左子结点存在的话，其右子节点必定存在，所以左子结点的next指针可以直接指向其右子节点，对于其右子节点的处理方法是，判断其父节点的next是否为空，若不为空，则指向其next指针指向的节点的左子结点，若为空则指向NULL。
C++解法一：

 1 // Recursion, more than constant space
 2 class Solution {
 3 public:
 4     void connect(TreeLinkNode *root) {
 5         if (!root) return;
 6         if (root->left) root->left->next = root->right;
 7         if (root->right) root->right->next = root->next? root->next->left : NULL;
 8         connect(root->left);
 9         connect(root->right);
10     }
11 };

对于非递归的解法要稍微复杂一点，但也不算特别复杂，需要用到queue来辅助，由于是层序遍历，每层的节点都按顺序加入queue中，而每当从queue中取出一个元素时，将其next指针指向queue中下一个节点即可。
C++ 解法二：

 1 // Non-recursion, more than constant space
 2 class Solution {
 3 public:
 4     void connect(TreeLinkNode *root) {
 5         if (!root) return;
 6         queue<TreeLinkNode*> q;
 7         q.push(root);
 8         q.push(NULL);
 9         while (true) {
10             TreeLinkNode *cur = q.front();
11             q.pop();
12             if (cur) {
13                 cur->next = q.front();
14                 if (cur->left) q.push(cur->left);
15                 if (cur->right) q.push(cur->right);
16             } else {
17                 if (q.size() == 0 || q.front() == NULL) return;
18                 q.push(NULL);
19             }
20         }
21     }
22 };

上面的方法巧妙的通过给queue中添加空指针NULL来达到分层的目的，使每层的最后一个节点的next可以指向NULL，那么我们可以换一种方法来实现分层，我们对于每层的开头元素开始遍历之前，先统计一下该层的总个数，用个for循环，这样for循环结束的时候，我们就知道该层已经被遍历完了。
C++ 解法三：

 1 class Solution {
 2 public:
 3     void connect(TreeLinkNode *root) {
 4         if (!root) return;
 5         queue<TreeLinkNode*> q;
 6         q.push(root);
 7         while (!q.empty()) {
 8             int size = q.size();
 9             for (int i = 0; i < size; ++i) {
10                 TreeLinkNode *t = q.front(); q.pop();
11                 if (i < size - 1) {
12                     t->next = q.front();
13                 }
14                 if (t->left) q.push(t->left);
15                 if (t->right) q.push(t->right);
16             }
17         }
18     }
19 };

上面三种方法虽然厉害，但是都不符合题意，题目中要求用O(1)的空间复杂度，所以我们来看下面这种碉堡了的方法。用两个指针start和cur，其中start标记每一层的起始节点，cur用来遍历该层的节点，设计思路之巧妙，不得不服。
C++ 解法四：

 1 class Solution {
 2 public:
 3     void connect(TreeLinkNode *root) {
 4         if (!root) return;
 5         TreeLinkNode *start = root, *cur = NULL;
 6         while (start->left) {
 7             cur = start;
 8             while (cur) {
 9                 cur->left->next = cur->right;
10                 if (cur->next) cur->right->next = cur->next->left;
11                 cur = cur->next;
12             }
13             start = start->left;
14         }
15     }
16 };
