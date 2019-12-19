[LeetCode] 24. Swap Nodes in Pairs 成对交换节点 

 
Given a linked list, swap every two adjacent nodes and return its head.
You may not modify the values in the list's nodes, only nodes itself may be changed.
 
Example:
Given 1->2->3->4, you should return the list as 2->1->4->3.
 
这道题不算难，是基本的链表操作题，我们可以分别用递归和迭代来实现。对于迭代实现，还是需要建立 dummy 节点，注意在连接节点的时候，最好画个图，以免把自己搞晕了，参见代码如下：
 
解法一：

class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        ListNode *dummy = new ListNode(-1), *pre = dummy;
        dummy->next = head;
        while (pre->next && pre->next->next) {
            ListNode *t = pre->next->next;
            pre->next->next = t->next;
            t->next = pre->next;
            pre->next = t;
            pre = t->next;
        }
        return dummy->next;
    }
};

 
递归的写法就更简洁了，实际上利用了回溯的思想，递归遍历到链表末尾，然后先交换末尾两个，然后依次往前交换：
 
解法二：

class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        if (!head || !head->next) return head;
        ListNode *t = head->next;
        head->next = swapPairs(head->next->next);
        t->next = head;
        return t;
    }
};
